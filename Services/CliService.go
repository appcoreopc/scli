package Services

import (
	"encoding/json"
	"fmt"
	"log"
	"path"
	"strings"
	"sync"

	"github.com/appcoreopc/scli/Fops"
	"github.com/appcoreopc/scli/HttpClient"
	"github.com/appcoreopc/scli/Model"
)

type CliService struct {
	wgOps sync.WaitGroup
}

// Any default configuration get from a service
// which a just a simple key / value db

func (s *CliService) Execute(installVersion int, cliSettingsModel *Model.ToolModel) {

	s.wgOps.Add(1)

	go s.RunExecute(installVersion, cliSettingsModel)

	s.wgOps.Wait()

	// log.Println("cli service executing")
	// if s.InitVersioningCheck(installVersion, cliSettingsModel.Version) {

	// 	log.Println("Download package from " + cliSettingsModel.Packageurl)

	// 	c := HttpClient.Client{}
	// 	c.Download(cliSettingsModel.Packageurl)

	// 	// unzip //
	// 	fz := Fops.FileUnzipper{}
	// 	fz.Unzip("command.cli.zip", ".")

	// }

	// Continue reading and displaying help //

	// jr := Fops.JsonReader{}
	// jr.GetCommandJson("command.cli/command.cli.json")

}

func (s *CliService) RunExecute(installVersion int, cliSettingsModel *Model.ToolModel) {

	defer s.wgOps.Done()

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

func (s *CliService) RunUpdate(model *Model.CommandCliModel) {

	s.wgOps.Add(1)

	go s.ExecRunSelfUpdate(model)

	s.wgOps.Wait()
}

func (s *CliService) ExecRunSelfUpdate(model *Model.CommandCliModel) {

	defer s.wgOps.Done()

	jsonWriter := new(Fops.JsonReader)

	dataJsonBytes, err := json.Marshal(model)

	if err != nil {
		panic(err)
	}

	jsonWriter.WriteJson("tools/command.cli.json", dataJsonBytes)

	fs := new(Fops.FileService)

	for _, element := range model.Tools {

		fmt.Println("element", element.Packageurl)
		c := HttpClient.Client{}
		c.Download(element.Packageurl)

		//rmdir

		// unzip //
		fz := Fops.FileUnzipper{}
		targetPath := path.Base(element.Packageurl)
		var location = strings.TrimSuffix(targetPath, path.Ext(targetPath))

		targetDir := "tools/" + location
		fs.RemoveDir(targetDir)
		fz.Unzip(targetPath, targetDir)
	}
}

func (s *CliService) InitVersioningCheck(installedVersion, currentVersion int) bool {
	return currentVersion > installedVersion
}

type ICliService interface {
	InitVersioningCheck(installedVersion, currentVersion int) bool

	Execute(cliSettingsModel *Model.CliSettingsModel)
}
