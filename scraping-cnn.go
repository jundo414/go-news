package main

import (
	"fmt"

	goquery "github.com/PuerkitoBio/goquery"
)

func RunCNN() {
	url := "https://edition.cnn.com"
	doc1 := getPages(url)
	fmt.Println("url:", url)
	fmt.Println("doc1:", doc1)

	doc1.Find(".cd a").Each(func(index int, s1 *goquery.Selection) {
		fmt.Println("STEP1")
		//title1 := s1.Text()

		attr1, exists1 := s1.Attr("href")
		//fmt.Printf("STEP1: %s (%s)\n", title1, attr1)
		fmt.Println("attr1:", attr1, "exists1:", exists1)

		if exists1 {
			fmt.Println("STEP2")
			doc2 := getPages(attr1)
			fmt.Println("attr1:", attr1)
			doc2.Find(".pg-headline").Each(func(index int, s2 *goquery.Selection) {
				fmt.Println("STEP3-1")
				title := s2.Text()
				fmt.Println("------------------------------------------------------------------------------------------------")
				fmt.Println(title)
				fmt.Println("------------------------------------------------------------------------------------------------")
			})
			doc2.Find(".l-container div p").Each(func(index int, s3 *goquery.Selection) {
				fmt.Println("STEP3-2")
				article := s3.Text()
				//attr3, exists3 := s.Attr("href")
				fmt.Println(article)
				fmt.Println("")
				ReadMessage(article)
				fmt.Println("")
			})
		}
	})
}
