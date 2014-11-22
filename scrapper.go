package gowler

import "io"
import "code.google.com/p/go.net/html"

func AggregateAnchorLinks(httpBody io.Reader) []string {
	//TODO: Do not scrap relative anchors
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
