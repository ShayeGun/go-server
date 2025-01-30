package util

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func GracefulShutdown(s *http.Server, ctx context.Context, cf context.CancelFunc) {
	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		fmt.Println("received SIGNAL")

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, shutdownCtxCancel := context.WithTimeout(ctx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := s.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}

		shutdownCtxCancel()
		cf()
	}()
}

func WriteJSONError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, res any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		WriteJSONError(w, http.StatusInternalServerError, err.Error())
	}
}
