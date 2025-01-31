package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetUpInitialRoutes(r *chi.Mux) *chi.Mux {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Non-functional
	r.Get("/v1/health", healthCheckHandler)
	r.Get("/v1/slow", slowHandler) // for testing GracefulShutdown()

	return r
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 5)
	w.Write([]byte("slow done"))
}
