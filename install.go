package main

import (
	"context"
	"io"
	"os/exec"

	"github.com/google/go-github/github"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

type Installer struct {
	api    *github.Client
	apiCtx context.Context
	path   string
}

type Release struct {
	RepoOwner string
	RepoName  string
	AssetID   int64
}

func (i *Installer) Install() error {
	release, err := parsePath(i.path)
	if err != nil {
		return err
	}
	src, rUrl, err := i.api.Repositories.DownloadReleaseAsset(i.apiCtx, rel.RepoOwner, rel.RepoName, rel.AssetID)
	if err != nil {
		return err
	}
	// uncompress
	asset, err := selfupdate.UncompressCommand(src, url, cmd)
	if err != nil {
		return err
	}

	return saveToGoBin(asset)
}

func parsePath(path string) (*Release, error) {
	// TODO: path covert Release struct
	return release, nil
}

func saveToGoBin(asset io.Reader) {

}

func Uninstall() {
	exec.Command(path)
}
