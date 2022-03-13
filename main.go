package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"ipfs-scraper/api"
	"ipfs-scraper/config"
	"ipfs-scraper/db"
	"ipfs-scraper/ipfs"
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
		log.Fatal().Err(err)
	}

	ipfsUri, err := loadEnvVar(IPFS_URI)
	if err != nil {
		log.Fatal().Err(err)
	}

	dbCfg := &config.Database{
		URI:      dbUri,
		Database: "ipfsthat",
	}

	database, err := db.New(dbCfg)
	if err != nil {
		log.Fatal().Err(err)
	}

	// Check we are able to connect to IPFS node
	err = ipfs.PingNode(ipfsUri)
	if err != nil {
		log.Fatal().Err(err)
	}

	api.Serve(database, ipfsUri)
}
