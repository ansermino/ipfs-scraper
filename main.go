package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"ipfs-scraper/ipfs"
	"ipfs-scraper/scraper"
	"os"
)

var URL = "https://github.com/rs/zerolog"
var API_URL = "/ip4/127.0.0.1/tcp/5001"

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	
	dir, err := scraper.FetchPage(URL)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to fetch page")
	}
	_, err = ipfs.StoreDir(API_URL, dir)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to store dir")
	}

	// TODO: Clean up local copy of site
}
