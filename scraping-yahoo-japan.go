package main

import (
	"fmt"

	goquery "github.com/PuerkitoBio/goquery"
)

func RunYahooJapan() {
	url := "https://news.yahoo.co.jp/"
	doc1 := getPages(url)
	doc1.Find(".topicsList_main li a").Each(func(index int, s1 *goquery.Selection) {
		//title1 := s1.Text()
		attr1, exists1 := s1.Attr("href")
		//fmt.Printf("STEP1: %s (%s)\n", title1, attr1)

		if exists1 {
			doc2 := getPages(attr1)
			cnt := 0
			doc2.Find(".tpcNews div a").Each(func(index int, s2 *goquery.Selection) {
				//title2 := s2.Text()
				attr2, exists2 := s2.Attr("href")
				//fmt.Printf("STEP2: %s (%s)\n", title2, attr2)
				cnt += 1
				if exists2 && cnt == 1 {
					doc3 := getPages(attr2)
					doc3.Find(".hd h1").Each(func(index int, s3 *goquery.Selection) {
						title := s3.Text()
						fmt.Println("------------------------------------------------------------------------------------------------")
						fmt.Println(title)
						fmt.Println("------------------------------------------------------------------------------------------------")
					})
					doc3.Find(".ynDetailText").Each(func(index int, s4 *goquery.Selection) {
						article := s4.Text()
						//attr3, exists3 := s.Attr("href")
						fmt.Println(article)
						fmt.Println("")
						ReadMessage(article)
						fmt.Println("")
					})
				}
			})
		}
	})
}
