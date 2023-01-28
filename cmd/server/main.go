package main

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/songmatrix/sync-service/config"
	"github.com/songmatrix/sync-service/database"
	"github.com/songmatrix/sync-service/server"
)

func main() {
	log.Info().
		Bool("debug", config.Debug).
		Str("server_address", config.ServerAddress).
		Msg("starting server")

	db, err := database.Init()
	if err != nil {
		log.Fatal().Err(err)
	}
	defer db.Close()

	router := server.Init()
	if err := http.ListenAndServe(config.ServerAddress, router); err != nil {
		log.Fatal().Err(err)
	}
}
