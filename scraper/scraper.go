package scraper

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"ipfs-scraper/models"
	neturl "net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func getSafeFilename(url string) (string, error) {
	u, err := neturl.Parse(url)
	if err != nil {
		return "", err
	}
	return u.Hostname() + "_" + strings.Replace(u.EscapedPath(), "/", "_", -1), nil
}

func FetchPage(url string) (string, *models.PageInfo, error) {
	log.Info().Str("url", url).Msg("Fetching page...")
	dir, err := getSafeFilename(url)
	if err != nil {
		return "", nil, err
	}

	err = os.Mkdir(dir, 0755)
	if err != nil {
		return "", nil, err
	}

	err = os.Chdir(dir)
	if err != nil {
		return "", nil, err
	}

	cmd := exec.Command("wget",
		"-q",
		"--show-progress",
		"--page-requisites",
		"--html-extension",
		"--convert-links",
		"--random-wait",
		"-e",
		"robots=off",
		"-nd",
		"--span-hosts",
		url)

	log.Info().Str("cmd", cmd.String()).Msg("Running wget")

	err = cmd.Run()
	//if err != nil {
	//	return "", nil, err
	//}

	path, err := os.Getwd()
	if err != nil {
		return "", nil, err
	}

	err = findAndRenameIndex(path)
	if err != nil {
		return "", nil, err
	}

	err = os.Chdir("../")
	if err != nil {
		return "", nil, err
	}

	log.Info().Str("path", path).Msg("Download successful")

	// TODO: Use actual page title
	info := &models.PageInfo{
		Url:   url,
		Title: dir,
	}
	return path, info, nil
}

func findAndRenameIndex(path string) error {
	htmlFiles := []string{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		if err != nil {

			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".html" {
			htmlFiles = append(htmlFiles, path)
		}

		return nil
	})

	if err != nil {
		return err
	}

	if len(htmlFiles) > 1 {
		for _, file := range htmlFiles {
			if filepath.Base(file) == "index.html" {
				return nil
			}
		}
		return fmt.Errorf("too many html files, can't locate index: %v", htmlFiles)
	} else if len(htmlFiles) < 1 {
		return errors.New("no html files found")
	}

	err = os.Rename(htmlFiles[0], filepath.Join(filepath.Dir(htmlFiles[0]), "index.html"))
	if err != nil {
		return err
	}
	return nil
}
