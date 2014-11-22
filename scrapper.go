package gowler

import "net/http"
import "log"
import "os"
import "io"
import "code.google.com/p/go.net/html"

func ScrapLinks(url string) (hyperlinks []string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Unable to perform get on the url %v", url)
		os.Exit(1)
	}
	defer resp.Body.Close()
	hyperLinks := aggregateAnchorLinks(resp.Body)
	return hyperLinks, nil
}

func aggregateAnchorLinks(httpBody io.Reader) []string {
	links := make([]string, 0)
	domPage := html.NewTokenizer(httpBody)
	for {
		tokenType := domPage.Next()
		if tokenType == html.ErrorToken {
			return links
		}
		token := domPage.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
	}
}
