package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/xoxoharsh/go-student-api/internal/config"
	"github.com/xoxoharsh/go-student-api/internal/http/handlers/student"
	"github.com/xoxoharsh/go-student-api/internal/storage/sqlite"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// database setup

	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatalf("Failed to setup storage: %s", err.Error())
	}

	slog.Info("Storage is initialized", slog.String("storage", "sqlite"))

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))

	// setup server
	server := http.Server {
		Addr: cfg.Addr, 
		Handler: router,
	}

	slog.Info("Server is running", slog.String("port", cfg.HTTPServer.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatalf("Failed to start server: %s", err.Error())
		}
	}()

	<-done
	
	slog.Info("Server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")
}