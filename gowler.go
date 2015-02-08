package gowler

import (
	"crypto/tls"
	"errors"
	"github.com/PuerkitoBio/purell"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//TODO: accept a set of beginUrls instead
func Gowler(beginUrls []string, crawlDelay int) (err error) {
	log.SetPrefix("gowler.go ")
	urlChannel := make(chan string)
	go func() {
		for _, urlString := range beginUrls {
			urlChannel <- urlString
		}
	}()
	for UrlString := range urlChannel {
		UrlString = sanitizeUrl(UrlString)
		if len(UrlString) == 0 {
			log.Println("URL %s was not sanitized. Skipping.", UrlString)
			continue
		}
		go ScrapLinks(UrlString, urlChannel)
	}
	return nil
}

func sanitizeUrl(UrlString string) (resolvedLinks string) {
	//Purell looks like the only library which does this efficiently.

	resolvedLinks = purell.MustNormalizeURLString(UrlString, purell.FlagsUsuallySafeGreedy)
	/*
		if err != nil {
			log.Println(err)
		}*/
	return resolvedLinks
}

func ScrapLinks(UrlString string, urlChannel chan string) {
	UrlString, err := validateUrl(UrlString)
	if err != nil {
		log.Println("The given urn %s is invalid.", UrlString)
		return
	}
	log.Println("Scraping URL: ", UrlString)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", UrlString, nil)

	if err != nil {
		log.Println("Unable to perform get on the url %v", UrlString)
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

	//TODO: use one more channel and a go routine to filter out mime time text/html only.
	//The other channel will contain the actual URl list that has to be used.
	// The go routine called will send a head request, and examine the content type header
	// Only if the contenty type is text/html we add it to the actual frontier.

	// Now regarding the frontier :
	// The frontier should be : - asynchrounus, should have a logic for ageing.
	//lets use a domain based re-visit policy, with aging for now.
	for _, link := range hyperLinks {
		go func() {
			urlChannel <- link
		}()
	}
}

//TODO: This is very rudimentary. Will need better TLS/HTTPS support.
func validateUrl(urlStr string) (string, error) {
	if strings.Contains(urlStr, "?") {
		return "", errors.New("Url contains a query param")
	} else {
		u, err := url.Parse(urlStr)
		if err != nil {
			log.Panic(err)
		}
		if u.Scheme == "" {
			u.Scheme = "http"
		}
		return u.String(), nil
	}
}
