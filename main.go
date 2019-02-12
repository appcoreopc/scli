package main

import (
	"github.com/appcoreopc/scli/cmd"
	"github.com/appcoreopc/scli/fops"
)

func main() {

	cmd.Execute()
	//c := httpClient.Client{}
	//c.Download("http://www.golang-book.com/public/pdf/gobook.pdf")

	fz := fops.FileUnzipper{}
	fz.Unzip("command.cli.zip", "command.cli")
}
