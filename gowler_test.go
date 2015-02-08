package gowler

import "testing"
import "log"

func TestGowler(t *testing.T) {
	const testUrl1 = "http://aishraj.com"
	const testUrl2 = "https://en.wikipedia.org"
	urlList := []string{testUrl1, testUrl2}
	err := Gowler(urlList, 3)
	if err != nil {
		log.Println("Unable to start crawling on the site %v", urlList)
	}
}
