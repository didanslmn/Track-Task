package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"teamtask-api/internal/config"
	"teamtask-api/internal/database"
	"teamtask-api/internal/httpserver"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("file .env tidak ditemukan")
	}
	// 1. Load confiuguration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	log.Printf("Statring API in %s mode", cfg.Env)

	// 2. Initialize Database Connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	pool, err := database.InitDB(ctx, cfg.Database)
	cancel()

	if err != nil {
		log.Fatalf("failed to initialize database pool: %v", err)
	}
	defer pool.Close()

	// 3. Build HttpServer + Health Route
	routes := httpserver.New(cfg)

	// 4. Setup context yang dibatalkan saat menerima SIGINT/SIGTERM
	shutdownCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 5. Jalankan server hingga shudown
	if err := httpserver.Run(shutdownCtx, cfg, routes); err != nil {
		log.Fatalf("server failed: %v", err)
	}

}
