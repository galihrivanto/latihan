package utils

import (
	"testing"
)

func TestGetChildURLs(t *testing.T) {
	urls, err := GetChildURLs("https://google.com")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(urls)
}

func TestGetChildURLsWithDomain(t *testing.T) {
	urls, err := GetChildURLs("https://google.com", "google.com")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(urls)
}
