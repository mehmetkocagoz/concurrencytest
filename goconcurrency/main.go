package main

import (
	"fmt"
	"net/http"
)

func fetchData(url string) (string, error) {
	// Make an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("Error fetching %s: %v", url, err), err
	}
	defer resp.Body.Close()

	// Read the response body
	body := make([]byte, 2000)
	_, err = resp.Body.Read(body)
	if err != nil {
		return fmt.Sprintf("Error reading response for %s: %v", url, err), err
	}

	// Return the data as a string
	return fmt.Sprintf("URL: %s\nResponse: %s", url, body), nil
}

func main() {
	// List of URLs to fetch
	urls := []string{"https://www.facebook.com", "https://www.google.com", "https://www.github.com", "https://online.yasar.edu.tr/"}

	// Fetch data sequentially
	for _, url := range urls {
		result, err := fetchData(url)
		if err != nil {
			fmt.Println(result)
		} else {
			fmt.Println(result)
		}
	}
}
