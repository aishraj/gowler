package gowler

import (
	"log"
)

func Gower(beginUrl string) (err error) {

	scrappedLinks, err := ScrapLinks(beginUrl)
	log.Println(scrappedLinks)
	return err
}
