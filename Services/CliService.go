package Services

import (
	"log"

	"github.com/appcoreopc/scli/Fops"
	"github.com/appcoreopc/scli/HttpClient"
	"github.com/appcoreopc/scli/Model"
)

type CliService struct {
}

// Any default configuration get from a service
// which a just a simple key / value db

func (s *CliService) Execute(cliSettingsModel *Model.CliSettingsModel) {

	log.Println("cli service executing")

	if s.InitVersioningCheck(cliSettingsModel.InstallVersion, cliSettingsModel.CurrentVersion) {

		log.Println("Download package from " + cliSettingsModel.ServiceUrl)
		c := HttpClient.Client{}
		c.Download(cliSettingsModel.ServiceUrl)
	}

	// Continue running //
	// check directory exist
	fs := Fops.FileService{}
	if fs.Exist(cliSettingsModel.CommandPath) {
		// show help or run argument
	}
}

func (s *CliService) InitVersioningCheck(installedVersion, currentVersion float32) bool {
	return currentVersion > installedVersion
}
