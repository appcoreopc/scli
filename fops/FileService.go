package fops

import (
	"os"
)

type FileService struct {
	defaultFolder string
}

func (f *FileService) Exist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func (f *FileService) Unpack(zipFile string) {

	if !f.Exist(f.defaultFolder) {
		os.MkdirAll(f.defaultFolder, 0755)
	}

	f.Unzip(zipFile)
}

func (f *FileService) Unzip(zipfile string) {

}

func (f *FileService) New() FileService {
	fs := FileService{}
	fs.defaultFolder = "command.cli"
	return fs
}
