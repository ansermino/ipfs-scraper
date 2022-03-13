package ipfs

import (
	"context"
	"fmt"
	"github.com/ipfs/go-cid"
	ipfsApi "github.com/ipfs/go-ipfs-api"
	files "github.com/ipfs/go-ipfs-files"
	ipfs "github.com/ipfs/go-ipfs-http-client"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/rs/zerolog/log"
	"io/fs"
	"ipfs-scraper/models"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func getUnixfsNode(path string) (files.Node, error) {
	st, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	f, err := files.NewSerialFile(path, false, st)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// PingNode checks if an IPFS http API is connectable
func PingNode(addr string) error {
	api := ipfsApi.NewShell(addr)

	if !api.IsUp() {
		return fmt.Errorf("client %s is down", addr)
	}

	return nil
}

func StoreDir(addr, path string, info *models.PageInfo) (*models.PageVersion, error) {
	log.Info().Str("path", path).Msg("Pushing directory to IPFS...")
	multi, err := ma.NewMultiaddr(addr)
	if err != nil {
		return nil, err
	}
	api, err := ipfs.NewApi(multi)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var rootCid cid.Cid
	err = filepath.Walk(path, func(subPath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		someDirectory, err := getUnixfsNode(subPath)
		if err != nil {
			return fmt.Errorf("Could not get path: %s, err: %s", subPath, err)
		}

		currentCid, err := api.Unixfs().Add(ctx, someDirectory)
		if err != nil {
			return fmt.Errorf("Could not add Directory: %s", err)
		}

		log.Info().Str("cid", currentCid.String()).Str("path", path).Msg("Stored file")

		if strings.EqualFold(path, subPath) {
			rootCid = currentCid.Cid()
		}

		return nil
	})

	log.Info().Msgf("Added directory to IPFS with root CID %s\n", rootCid.String())

	version := &models.PageVersion{
		Url:       info.Url,
		Title:     info.Title,
		Timestamp: time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST"),
		Cid:       rootCid.String(),
	}

	log.Info().Str("path", path).Msg("Removing local copy...")
	err = os.RemoveAll(path)
	if err != nil {
		return nil, err
	}

	return version, nil
}
