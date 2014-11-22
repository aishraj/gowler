package gowler

import "testing"
import "log"

func TestGowler(t *testing.T) {
	const testUrl = "http://aishraj.com"
	err := Gowler(testUrl, 3)
	if err != nil {
		log.Println("Unable to start crawling on the site %v", testUrl)
	}
}
