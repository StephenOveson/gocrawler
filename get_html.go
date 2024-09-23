package main

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	if resp.StatusCode >= 400 {
		return "", errors.New("unable to access url")
	}
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", errors.New("url does not contain html")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
