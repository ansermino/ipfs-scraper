package api

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"ipfs-scraper/db"
	"ipfs-scraper/ipfs"
	"ipfs-scraper/models"
	"ipfs-scraper/scraper"
	"net/http"
)

type api struct {
	db       db.Database
	ipfsAddr string
}

type allPagesResp struct {
	Pages []*models.PageInfo `json:"pages"`
}

type versionsResp struct {
	Versions []*models.PageVersion `json:"versions"`
}

func Serve(db db.Database, ipfsAddr string) {
	api := &api{
		db:       db,
		ipfsAddr: ipfsAddr,
	}

	serve(api)
}

func serve(api *api) {

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		err := addPage(w, r, api)
		if err != nil {
			log.Error().Err(err).Msg("Request failed")
		}
	})

	http.HandleFunc("/pages", func(w http.ResponseWriter, r *http.Request) {
		err := getAllPages(w, r, api)
		if err != nil {
			log.Error().Err(err).Msg("Request failed")
		}
	})

	http.HandleFunc("/versions", func(w http.ResponseWriter, r *http.Request) {
		err := getVersions(w, r, api)
		if err != nil {
			log.Error().Err(err).Msg("Request failed")
		}
	})

	log.Info().Msg("Server listening...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal().Err(err)
	}
}

func addPage(w http.ResponseWriter, r *http.Request, api *api) error {
	url := r.URL.Query().Get("url")

	log.Info().Str("url", url).Msg("Handling /add request")

	path, info, err := scraper.FetchPage(url)
	if err != nil {
		return err
	}

	version, err := ipfs.StoreDir(api.ipfsAddr, path, info)
	if err != nil {
		return err
	}

	err = api.db.CreatePageInfo(context.TODO(), info)
	if err != nil {
		return err
	}

	err = api.db.CreatePageVersion(context.TODO(), version)
	if err != nil {
		return err
	}

	r.Body.Close()

	return nil
}

func getAllPages(w http.ResponseWriter, r *http.Request, api *api) error {
	log.Info().Msg("Handling /list request")

	pages, err := api.db.ViewAllPages(context.TODO())
	if err != nil {
		return err
	}

	resp := allPagesResp{Pages: pages}

	bz, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	_, err = w.Write(bz)
	log.Info().Msgf("Served %d pages", len(pages))
	return err
}

func getVersions(w http.ResponseWriter, r *http.Request, api *api) error {
	title := r.URL.Query().Get("title")
	log.Info().Str("title", title).Msg("Handling /versions request")

	vers, err := api.db.ViewPageVersions(context.TODO(), title)
	if err != nil {
		return err
	}

	resp := versionsResp{vers}

	bz, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	_, err = w.Write(bz)
	if err != nil {
		return err
	}
	log.Info().Msgf("Served %d versions", len(vers))
	return nil
}
