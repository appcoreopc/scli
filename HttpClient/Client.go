package HttpClient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

func (c *Client) GetJson(targetUrl string, target interface{}) error {

	var myClient = &http.Client{Timeout: 10 * time.Second}

	r, err := myClient.Get(targetUrl)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

type IClientHttp interface {
	GetJson(targetUrl string, target interface{}) error

	Download(downloadUri string)
}
