package services

import (
	"net/url"
	"strings"
)

func ProcessURL(originalURL string, operation string) (string, error) {
	parsedURL, err := url.Parse(originalURL)
	if err != nil {
		return "", err
	}

	switch operation {
	case "canonical":
		parsedURL.RawQuery = ""
		parsedURL.Path = strings.TrimSuffix(parsedURL.Path, "/")
		return parsedURL.String(), nil

	case "redirection":
		parsedURL.Host = "www.byfood.com"
		parsedURL.Path = strings.ToLower(parsedURL.Path)
		return parsedURL.String(), nil

	case "all":
		parsedURL.Host = "www.byfood.com"
		parsedURL.Path = strings.ToLower(parsedURL.Path)
		parsedURL.RawQuery = ""
		parsedURL.Path = strings.TrimSuffix(parsedURL.Path, "/")
		return parsedURL.String(), nil

	default:
		return "", nil
	}
}
