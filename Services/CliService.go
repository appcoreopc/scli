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

func (s *CliService) Execute(installVersion int, cliSettingsModel *Model.ToolModel) {

	log.Println("cli service executing")

	if s.InitVersioningCheck(installVersion, cliSettingsModel.Version) {

		log.Println("Download package from " + cliSettingsModel.Packageurl)

		c := HttpClient.Client{}
		c.Download(cliSettingsModel.Packageurl)

		// unzip //
		fz := Fops.FileUnzipper{}
		fz.Unzip("command.cli.zip", ".")
	}

	// Continue reading and displaying help //

	// jr := Fops.JsonReader{}
	// jr.GetCommandJson("command.cli/command.cli.json")

}

func (s *CliService) InitVersioningCheck(installedVersion, currentVersion int) bool {
	return currentVersion > installedVersion
}

type ICliService interface {
	InitVersioningCheck(installedVersion, currentVersion int) bool

	Execute(cliSettingsModel *Model.CliSettingsModel)
}
