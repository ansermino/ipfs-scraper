package ipfs

import (
	"context"
	"fmt"
	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	ipfs "github.com/ipfs/go-ipfs-http-client"
	"github.com/rs/zerolog/log"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)
import ma "github.com/multiformats/go-multiaddr"

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

func StoreDir(addr, path string) (cid.Cid, error) {
	multi, err := ma.NewMultiaddr(addr)
	if err != nil {
		return cid.Undef, err
	}
	api, err := ipfs.NewApi(multi)
	if err != nil {
		return cid.Undef, err
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

	return rootCid, nil
}
