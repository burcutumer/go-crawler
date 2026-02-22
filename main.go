package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://news.ycombinator.com/"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("status code error:", resp.StatusCode)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("HTML can not loaded ")
	}

	doc.Find(".athing").Each(func(i int, s *goquery.Selection) {
		rankRaw := s.Find("span.rank").Text()
		rankClean := strings.TrimSuffix(rankRaw, ".")
		atag := s.Find(".titleline > a")
		title := atag.Text()
		link, hrefExists := atag.Attr("href")
		if hrefExists {
			fmt.Printf("%s - %s , link: %s \n", rankClean, title, link)
		}

	})

	doc.Find(".subtext").Each(func(i int, s *goquery.Selection) {
		score := s.Find(".score").Text()
		user := s.Find(".subline>a").First().Text()
		date := s.Find("span.age").Text()
		fmt.Printf("score: %s,  user: %s, date: %s \n", score, user, date)

	})

}
