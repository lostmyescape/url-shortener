package main

import (
	"database/sql"
	"github.com/user/urlProject/config"
	"github.com/user/urlProject/internal/storage"
	"log"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	cfg := config.LoadConfig()

	storage, err := storage.NewStorage(cfg)
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
		return
	}

	//defer storage.DB.Close()

	defer func(DB *sql.DB) {
		err := storage.DB.Close()
		if err != nil {
			log.Fatalf("DB close error: %v", err)
			return
		}
	}(storage.DB)

	// TODO: init config: cleanenv
	// TODO: init logger: slog
	// TODO: init storage: postgres
	// TODO: init router: chi, "chi render"
	// TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	}

	return log

}
