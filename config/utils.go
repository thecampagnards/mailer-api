package config

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// IsValidURL
func IsValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return false
	}
	return true
}

// FetchBodyFromURL
func FetchBodyFromURL(u string) (string, error) {
	if !IsValidURL(u) {
		return "", errors.New("The URL is not valid")
	}

	// Fetching the url
	resp, err := http.Get(u)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	// Get the body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}
