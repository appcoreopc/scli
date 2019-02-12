package services

import (
	"github.com/appcoreopc/scli/httpClient"
	"github.com/appcoreopc/scli/fops"
)

type CliService struct {
}

// Any default configuration get from a service
// which a just a simple key / value db

func (s *CliService) Execute(commandPath string) {

	// check directory exist
	fs := fops.FileService{}

	if !fs.Exist(commandPath) {
		

	}
}

func (s *CliService) DownloadAndSetup() {

	httpClient.Client

}
