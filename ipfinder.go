package ipfinder

import (
	"fmt"
	"io"
	"net/http"
)

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

type IPFinder struct {
	URL    string
	Client HTTPClient
}

func NewIPFinder(url string) *IPFinder {
	return &IPFinder{
		URL:    url,
		Client: &http.Client{},
	}
}

func (f *IPFinder) GetIP() (string, error) {
	response, err := f.Client.Get(f.URL)
	if err != nil {
		return "", fmt.Errorf("error making the request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	return string(body), nil
}
