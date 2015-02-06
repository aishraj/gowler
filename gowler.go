package gowler

import (
	"crypto/tls"
	"github.com/PuerkitoBio/purell"
	"log"
	"net/http"
	"os"
	"strings"
)

//TODO: accept a set of beginUrls instead
func Gowler(beginUrl string, crawlDelay int) (err error) {
	log.SetPrefix("gowler.go   ")
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
	//Purell looks like the only library which does this efficiently.
	resolvedLinks = purell.MustNormalizeURLString(url, purell.FlagsUsuallySafeGreedy)
	return resolvedLinks
}

func ScrapLinks(url string, urlChannel chan string) {
	if !isUrlValid(url) {
		return
	}
	log.Println("Scraping URL: ", url)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}
	//TODO: use one more channel and a go routine to filter out mime time text/html only.
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("Unable to perform get on the url %v", url)
		log.Fatal(err)
		os.Exit(1)
	}

	req.Header.Set("User-Agent", "Golang Spider - Gowler 0.0.1")

	resp, err := httpClient.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	hyperLinks := AggregateAnchorLinks(resp.Body)
	for _, link := range hyperLinks {
		go func() {
			urlChannel <- link
		}()
	}
}

func isUrlValid(url string) bool {
	if strings.Contains(url, "?") {
		return false
	}
	return true
}
