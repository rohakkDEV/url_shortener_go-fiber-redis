package helpers

import (
	"os"
	"strings"
)

func EnforceHTTP(rawURL string) string {
	rawURL = strings.TrimSpace(rawURL)

	if !strings.HasPrefix(rawURL, "http://") &&
		!strings.HasPrefix(rawURL, "https://") {
		return "http://" + rawURL
	}

	return rawURL
}

func RemoveDomainError(rawURL string) bool {
	domain := os.Getenv("DOMAIN")

	if rawURL == domain {
		return false
	}

	newURL := strings.TrimPrefix(rawURL, "http://")
	newURL = strings.TrimPrefix(newURL, "https://")
	newURL = strings.TrimPrefix(newURL, "www.")
	newURL = strings.Split(newURL, "/")[0]

	if strings.EqualFold(newURL, domain) {
		return false
	}

	return true
}
