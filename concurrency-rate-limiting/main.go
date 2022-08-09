package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/pkg/errors"
)

const (
	defaultLimit = 100
)

func main() {
	// command-line flags.
	filename := flag.String("filename", "", "specify the filename that the urls should be read from [required]")
	threshold := flag.Int("threshold", defaultLimit, "set a threshold on the number of goroutines running [optional]")

	flag.Parse()

	if filename == nil || len(*filename) == 0 {
		log.Fatal("filename not provided")
	}

	if *threshold == defaultLimit {
		log.Printf("threshold on the max number of goroutines running not provided, using default limit %v", defaultLimit)
	} else {
		log.Printf("using %v as the threshold on the max number of goroutines running concurrently", *threshold)
	}

	urls, err := readFile(*filename)
	if err != nil {
		log.Fatal(err)
	}

	// ch is a buffered channel to ensure rate limiting.
	ch := make(chan struct{}, *threshold)
	// urlsHash is a slice that will maintain the order.
	urlsHash := make([]string, len(urls))

	var wg sync.WaitGroup

	for idx, url := range urls {
		ch <- struct{}{}

		wg.Add(1)
		// anonymous function which starts as a goroutine.
		// the anonymous function gets the content of the url,
		// calculates the MD5 of the content, and stores it
		// in a slice to maintain the order.
		go func(idx int, url string) {
			defer wg.Done()

			data, err := getURLContent(url)
			if err != nil {
				log.Fatal(err)
			}

			urlsHash[idx] = generateHash(data)

			<-ch
		}(idx, url)
	}

	wg.Wait()

	for _, hash := range urlsHash {
		log.Println(hash)
	}
}

// readFile reads the file provided and returns the contents.
func readFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open file")
	}
	defer file.Close()

	urls := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.Wrap(err, "unable to read file")
	}

	return urls, nil
}

// getURLContent sends a GET request to the provided URL and returns the response if there is no error.
func getURLContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "unable to send request")
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response body")
	}

	return data, nil
}

// generateHash returns the MD5 checksum of the data provided.
func generateHash(data []byte) string {
	hash := md5.Sum(data)

	return fmt.Sprintf("%x", hash)
}
