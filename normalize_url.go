package main

import (
	"fmt"
	"net/url"
	"path"
)

func normalizeURL(rawURL string) (string, error) {
	parsedUrl, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("unable to parse url: %w", err)
	}

	normalizedUrl, err := url.JoinPath(parsedUrl.Hostname() + path.Clean(parsedUrl.EscapedPath()))
	if err != nil {
		return "", fmt.Errorf("unable to parse url: %w", err)
	}

	return normalizedUrl, nil
}
