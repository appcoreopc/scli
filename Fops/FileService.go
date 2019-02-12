package Fops

import (
	"log"
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

func (f *FileService) RemoveDir(path string) bool {

	err := os.Remove(path)
	if err != nil {
		return true
	}
	return false
}

func (f *FileService) CreateDir(path string) bool {

	if err := os.MkdirAll(path, 07555); err != nil {
		log.Println("Unable create directory: ", path)
		return false
	}
	return true
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
