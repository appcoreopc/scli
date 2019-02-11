package httpClient

import (
	"fmt"

	"github.com/cavaliercoder/grab"
)

type Client struct {
	filename    string
	downloadUri string
}

func (c *Client) Download(downloadUri string) {

	resp, err := grab.Get(".", downloadUri)

	if err != nil {
		fmt.Println("error dowloading:", downloadUri)
	}

	fmt.Print("downloaded file", resp.Filename)
}
