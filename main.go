package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"ipfs-scraper/api"
	"ipfs-scraper/config"
	"ipfs-scraper/db"
	"os"
)

var URL = "https://github.com/rs/zerolog"
var IPFS_URL = "/ip4/127.0.0.1/tcp/5001"

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	dbCfg := &config.Database{
		URI:      "mongodb://root:rootpassword@127.0.0.1:27017",
		Database: "ipfsthat",
	}

	database, err := db.New(dbCfg)
	if err != nil {
		log.Fatal().Err(err)
	}

	api.Serve(database, IPFS_URL)

	//dir, err := scraper.FetchPage(URL)
	//if err != nil {
	//	log.Fatal().Err(err).Msg("Failed to fetch page")
	//}
	//_, err = ipfs.StoreDir(IPFS_URL, dir)
	//if err != nil {
	//	log.Fatal().Err(err).Msg("Failed to store dir")
	//}

	// TODO: Clean up local copy of site
}
