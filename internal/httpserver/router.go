package httpserver

import (
	"context"
	"errors"
	"log"
	"net/http"
	"teamtask-api/internal/config"
	"teamtask-api/internal/shared"
	"time"
)

func New(cfg *config.Config) http.Handler {
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("GET /health", shared.HealthCheckHandler)

	return mux
}

// Run menjalakan http server dan graceful shutdown
// ketika ctx dibatalkan (mis. menerima SIGINT/SIGTERM)
func Run(ctx context.Context, cfg *config.Config, handler http.Handler) error {
	server := &http.Server{
		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
		Handler:      handler,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	// jalankan ListenAndServe dalam goroutine

	errCh := make(chan error, 1)
	go func() {
		log.Printf("Server started at %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
			return
		}
	}()

	// tunggu ctx dibatalkan atau ada error
	select {
	case <-ctx.Done():
		log.Println("Shutting down server...")
		// graceful shutdown
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer shutdownCancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("Server shutdown error: %v", err)
			return err
		}
		log.Println("Server gracefully stopped")
		return nil
	case err := <-errCh:
		return err
	}
}
