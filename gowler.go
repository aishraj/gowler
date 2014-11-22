package gowler

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
)

func Gowler(beginUrl string, maxDepth int) (err error) {
	log.SetPrefix("gowler.go")
	urlChannel := make(chan string)
	go func() {
		urlChannel <- beginUrl
	}()
	for url := range urlChannel {
		url = sanitizeUrl(url)
		go ScrapLinks(url, urlChannel)
	}
	return nil
}

func sanitizeUrl(url string) (resolvedLinks string) {
	//TODO: Write a real function that does this task.
	resolvedLinks = url
	return resolvedLinks
}

func ScrapLinks(url string, urlChannel chan string) {
	log.Println("Scraping URL: ", url)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}
	resp, err := httpClient.Get(url)
	if err != nil {
		log.Println("Unable to perform get on the url %v", url)
		os.Exit(1)
	}
	defer resp.Body.Close()
	hyperLinks := AggregateAnchorLinks(resp.Body)
	for _, link := range hyperLinks {
		go func() {
			urlChannel <- link
		}()
	}
}
