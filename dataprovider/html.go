package dataprovider

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type HTMLDataProvider struct {
	client *http.Client
}

func NewHTMLDataProvider() *HTMLDataProvider {
	return &HTMLDataProvider{
		client: &http.Client{
			Timeout: 2 * time.Second,
		},
	}
}

func (p *HTMLDataProvider) Fetch(url string) (io.Reader, error) {
	resp, err := p.client.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status != 200. Code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return strings.NewReader(string(data)), nil
}
