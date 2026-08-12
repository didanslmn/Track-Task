package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"teamtask-api/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Migration error: %v", err)
	}
}

func run() error {
	// 1. Load .env (abaikan jika tidak ada, fallback ke system env)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// 2. Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	// 3. Setup DSN (URL-encoded untuk keamanan karakter khusus pada password)
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.Database.User, cfg.Database.Password),
		Host:   fmt.Sprintf("%s:%d", cfg.Database.Host, cfg.Database.Port),
		Path:   cfg.Database.DBName,
	}
	q := u.Query()
	q.Set("sslmode", cfg.Database.SSLMode)
	if cfg.Database.ChannelBinding != "" {
		q.Set("sslchannelbinding", cfg.Database.ChannelBinding)
	}
	u.RawQuery = q.Encode()
	dbURL := u.String()

	// 4. Buat path migration bisa dikonfigurasi via environment
	migrationDir := os.Getenv("MIGRATION_PATH")
	if migrationDir == "" {
		migrationDir = "internal/database/migrations"
	}
	absMigrationDir, err := filepath.Abs(migrationDir)
	if err != nil {
		return fmt.Errorf("failed to resolve absolute path for migration dir: %w", err)
	}
	sourceURL := fmt.Sprintf("file://%s", filepath.ToSlash(absMigrationDir))

	// 5. Inisialisasi migrate
	m, err := migrate.New(sourceURL, dbURL)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}

	// Pastikan koneksi ditutup dengan benar di akhir (selalu dijalankan walau terjadi error)
	defer func() {
		sourceErr, dbErr := m.Close()
		if sourceErr != nil {
			log.Printf("Warning: error closing source: %v", sourceErr)
		}
		if dbErr != nil {
			log.Printf("Warning: error closing DB: %v", dbErr)
		}
	}()

	// 6. Setup CLI Flags & Arguments
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [command] [args]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Commands:\n")
		fmt.Fprintf(os.Stderr, "  up          Apply all pending migrations\n")
		fmt.Fprintf(os.Stderr, "  down        Rollback 1 migration\n")
		fmt.Fprintf(os.Stderr, "  force <v>   Force set version to <v> (useful for fixing dirty state)\n")
		fmt.Fprintf(os.Stderr, "  version     Print current migration version\n")
		fmt.Fprintf(os.Stderr, "  drop        Drop all tables (DANGER: only use in dev!)\n\n")
		fmt.Fprintf(os.Stderr, "Environment:\n")
		fmt.Fprintf(os.Stderr, "  MIGRATION_PATH  Path to migration folder (default: ./migrations)\n")
	}

	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		return fmt.Errorf("missing command")
	}

	command := flag.Arg(0)

	// 7. Execute Command
	switch command {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			return fmt.Errorf("failed to run up migrations: %w", err)
		}
		if err == migrate.ErrNoChange {
			log.Println("No new migrations to apply.")
		} else {
			log.Println("Migrations applied successfully!")
		}

	case "down":
		if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
			return fmt.Errorf("failed to run down migration: %w", err)
		}
		if err == migrate.ErrNoChange {
			log.Println("Already at base version (0).")
		} else {
			log.Println("Migration rolled back successfully!")
		}

	case "force":
		if flag.NArg() < 2 {
			flag.Usage()
			return fmt.Errorf("usage: force <version>")
		}
		version, err := strconv.ParseInt(flag.Arg(1), 10, 64)
		if err != nil {
			return fmt.Errorf("invalid version format: %w", err)
		}
		if version > int64(math.MaxInt) || version < int64(math.MinInt) {
			return fmt.Errorf("version %d exceeds integer bounds", version)
		}

		if err := m.Force(int(version)); err != nil {
			return fmt.Errorf("failed to force version: %w", err)
		}
		log.Printf("Successfully forced version to %d\n", version)

	case "version":
		version, dirty, err := m.Version()
		if err != nil && err != migrate.ErrNilVersion {
			return fmt.Errorf("failed to get version: %w", err)
		}
		if err == migrate.ErrNilVersion {
			log.Println("No migrations applied yet (version 0).")
		} else {
			log.Printf("Current version: %d (Dirty: %v)\n", version, dirty)
		}

	case "drop":
		log.Println("WARNING: Dropping all database objects...")
		if err := m.Drop(); err != nil {
			return fmt.Errorf("failed to drop database: %w", err)
		}
		log.Println("Database dropped successfully.")

	default:
		flag.Usage()
		return fmt.Errorf("unknown command: %s", command)
	}

	return nil
}
