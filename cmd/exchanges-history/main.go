package main

import (
	"context"
	"database/sql"
	"os"

	"github.com/rs/zerolog"

	"github.com/mkokoulin/exchanges-history-app/internal/config"
	"github.com/mkokoulin/exchanges-history-app/internal/database/postgres"
	"github.com/mkokoulin/exchanges-history-app/internal/handlers"
	"github.com/mkokoulin/exchanges-history-app/internal/router"
	"github.com/mkokoulin/exchanges-history-app/internal/server"
)

func main() {
	logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.New()

	_, err := postgres.RunMigration(cfg.DataBaseURI)
	if err != nil {
		logger.Error().Msg(err.Error())
	}

	logger.Log().Msg("ServerAddress: " + cfg.ServerAddress)
	logger.Log().Msg("BaseURL: " + cfg.BaseURL)
	logger.Log().Msg("DataBase: " + cfg.DataBaseURI)

	db, err := sql.Open("postgres", cfg.DataBaseURI)
	if err != nil {
		logger.Error().Msg(err.Error())
	}

	repo := postgres.New(db, &logger)

	h := handlers.New(repo, cfg.BaseURL, &logger)
	r := router.New(h)

	s := server.New(cfg.ServerAddress, r)
	defer s.Shutdown(ctx)

	s.Start()
}
