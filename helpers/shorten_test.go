package helpers

import (
	"testing"
)

func TestShorten(t *testing.T) {
	url := "https://www.example.com/long/url/that/needs/to/be/shortened"
	shortURL := Shorten(url)
	if len(shortURL) != shortURLLength {
		t.Errorf("Expected short URL length of %d, but got %d", shortURLLength, len(shortURL))
	}
}
