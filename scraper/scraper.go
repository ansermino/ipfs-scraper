package scraper

import (
	"github.com/rs/zerolog/log"
	neturl "net/url"
	"os"
	"os/exec"
	"strings"
)

func getSafeFilename(url string) (string, error) {
	u, err := neturl.Parse(url)
	if err != nil {
		return "", err
	}
	return u.Hostname() + "_" + strings.Replace(u.EscapedPath(), "/", "_", -1), nil
}

func FetchPage(url string) (string, error) {
	dir, err := getSafeFilename(url)
	if err != nil {
		return "", err
	}

	err = os.Mkdir(dir, 0755)
	if err != nil {
		return "", err
	}

	err = os.Chdir(dir)
	if err != nil {
		return "", err
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
	if err != nil {
		return "", err
	}

	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	log.Info().Str("path", path).Msg("Download successful")

	return path, nil
}
