package gowler

import "testing"

func TestScraper(t *testing.T) {

	const in = "http://www.aishraj.com"

	out := make([]string, 0)
	out = append(out, "//github.com/aishraj")
	scrapResult, e := ScrapLinks("http://www.aishraj.com")
	if e == nil && !contains(scrapResult, out[0]) {
		t.Errorf("ScrapLinks(%v) = %v, want %v", in, scrapResult, out)
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
