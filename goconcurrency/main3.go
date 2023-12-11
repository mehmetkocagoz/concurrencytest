package main

import (
	"fmt"
	"net/http"
	"sync"
)

func fetchData(url string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Make an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body := make([]byte, 2000)
	_, err = resp.Body.Read(body)
	if err != nil {
		ch <- fmt.Sprintf("Error reading response for %s: %v", url, err)
		return
	}

	// Send the data through the channel
	ch <- fmt.Sprintf("URL: %s\nResponse: %s", url, body)
}

func main() {
	// List of URLs to fetch
	urls := []string{"https://www.facebook.com", "https://www.google.com", "https://www.github.com", "https://"}

	// Create a channel to receive data
	dataChannel := make(chan string, len(urls))

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Launch a goroutine for each URL
	for _, url := range urls {
		wg.Add(1)
		go fetchData(url, dataChannel, &wg)
	}

	// Use a goroutine to close the channel when all goroutines are done
	go func() {
		wg.Wait()
		close(dataChannel)
	}()

	// Read from the channel until it's closed
	for data := range dataChannel {
		fmt.Println(data)
	}
}
