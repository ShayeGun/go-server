package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ShayeGun/go-server/internal/service"
	"github.com/ShayeGun/go-server/internal/util"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RepositoryInterface interface {
	GetUserTable() service.UserRepositoryInterface
}

type externalDependencies struct {
	RepositoryInterface
	//logger
	// cache
}

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount(exd *externalDependencies) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		// Non-functional
		r.Get("/health", app.healthCheckHandler)
		r.Get("/slow", app.slowHandler) // for testing GracefulShutdown()

		// Functional
		r.Get("/users/{uid}", app.getUser(exd))
		r.Post("/users", app.addUser(exd))
		r.Delete("/users/{uid}", app.deleteUser(exd))
		r.Patch("/users/{uid}", app.updateUser(exd))
	})

	return r
}

func (app *application) run(mux http.Handler) {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 30,
	}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	log.Printf("server is running on port %s\n", app.config.addr)

	util.GracefulShutdown(srv, serverCtx, serverStopCtx)

	err := srv.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-serverCtx.Done()

	log.Println("server shutdown")

}
