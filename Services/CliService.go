package services

import (
	"fmt"
	"strconv"

	"github.com/appcoreopc/scli/Fops"
	"github.com/appcoreopc/scli/HttpClient"
	"github.com/appcoreopc/scli/Model"
)

type CliService struct {
}

// Any default configuration get from a service
// which a just a simple key / value db

func (s *CliService) Execute(commandPath string) {

	s.InitVersioningCheck()

	// Continue running //
	// check directory exist
	fs := Fops.FileService{}
	if fs.Exist(commandPath) {
		// show help or run argument
	}
}

func (s *CliService) InitVersioningCheck() int {

	hc := &HttpClient.Client{}
	// checks service
	settings := Model.CliSettingsModel{}
	_ = hc.GetJson("calls a default service", &settings)

	version, err := strconv.Atoi(settings.Version)
	if err != nil {
		fmt.Println("error converting to integer")
	}

	switch version {

	case 0:
		break
	case 1:

	}

	return 0
}

type ActionType int

const (
	NoAction ActionType = iota
	Upgrade
	ReInstallCli
)

func (a ActionType) ToString() string {

	return [...]string{"No Action", "Upgrade", "ReInstallCli"}[a]
}
