package main

import (
	"net/http"

	goquery "github.com/PuerkitoBio/goquery"
)

func getPages(url string) *goquery.Document {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.80 Safari/537.36")

	// リスエスト投げてレスポンスを得る
	client := &http.Client{}
	resp, err := client.Do(req)

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		panic(err)
	}
	return doc
}
