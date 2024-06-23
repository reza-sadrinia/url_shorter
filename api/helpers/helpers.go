package helpers

import (
	"os"
	"strings"
)

// EnforceHTTP ensures the URL starts with "http://"
func EnforceHTTP(url string) string {
	if len(url) < 4 || url[:4] != "http" {
		return "http://" + url
	}
	return url
}

// RemoveDomainError checks if the URL is the same as APP_DOMAIN
func RemoveDomainError(url string) bool {
	if url == os.Getenv("APP_DOMAIN") {
		return false
	}

	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	if newURL == os.Getenv("APP_DOMAIN") {
		return false
	}
	return true
}
