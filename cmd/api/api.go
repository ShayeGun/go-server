package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ShayeGun/go-server/internal/routes"
	"github.com/ShayeGun/go-server/internal/util"

	"github.com/go-chi/chi/v5"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount(exd *routes.ExternalDependencies) http.Handler {
	r := chi.NewRouter()

	// Non-functional
	routes.SetUpInitialRoutes(r)

	// Functional
	u := routes.NewUserRoutes(exd)
	r = u.SetupUserRoutes(r)

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
