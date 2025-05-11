package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Info struct {
	title string
}

func main() {
	url := "https://ria.ru/"
	response, error := http.Get(url)
	document, err := goquery.NewDocumentFromReader(response.Body)
	if (err != nil) || (error != nil) {
		log.Fatalln(err)
		log.Fatalln(error)
	}
	info := make([]Info, 0)
	document.Find("span.cell-list__item-title").Each(func(i int, s *goquery.Selection) {
		ttl := s.Text()
		info = append(info, Info{
			title: ttl,
		})
	})
	fmt.Println(info)
}
