package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"ipfs-scraper/api"
	"ipfs-scraper/config"
	"ipfs-scraper/db"
	"os"
)

var MONGODB_URI = "MONGODB_URI"
var IPFS_URI = "IPFS_URI"

func loadEnvVar(ev string) (string, error) {
	value := os.Getenv(ev)
	if value == "" {
		return "", fmt.Errorf("%s is required", ev)
	}
	return value, nil
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	dbUri, err := loadEnvVar(MONGODB_URI)
	if err != nil {
		log.Fatal().Err(err).Msg("Loading env var failed")
	}

	ipfsUri, err := loadEnvVar(IPFS_URI)
	if err != nil {
		log.Fatal().Err(err).Msg("Loading env var failed")
	}

	dbCfg := &config.Database{
		URI:      dbUri,
		Database: "ipfsthat",
	}

	database, err := db.New(dbCfg)
	if err != nil {
		log.Fatal().Err(err).Msg("DB connection failed")
	}

	api.Serve(database, ipfsUri)
}
