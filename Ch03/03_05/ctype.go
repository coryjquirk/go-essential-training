package main

import (
	"fmt"
	"net/http"
)

func main() {
	contentType, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Printf("contentType: %s\n", contentType)
	}
}

// contentType will reutrn the value of Content-Type header
// returned by making an HTTP GET request to URL
func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	contentType := resp.Header.Get("Content-type")
	if contentType == "" {
		return "", fmt.Errorf("can't find Content-Type header")
	}
	return contentType, nil
}
