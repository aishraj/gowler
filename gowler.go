package gowler

import (
	"log"
)

func Gowler(beginUrl string, maxDepth int) (err error) {

	scrappedLinks, err := ScrapLinks(beginUrl)
	log.Println(scrappedLinks)
	return err
}
