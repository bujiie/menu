package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func fetchMenu(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return goquery.NewDocumentFromReader(res.Body)
}

func ArizmendiBakery() {
	doc, err := fetchMenu("http://www.arizmendi-bakery.org/arizmendi-emeryville-pizza")
	if err != nil {
		log.Fatal(err)
	}

	meta := doc.Find(".sqs-block.html-block.sqs-block-html")
	fmt.Println("Arizmendi Bakery")
	fmt.Println(meta.Find("blockquote").Text())
	fmt.Println()

	doc.Find(".sqs-block.button-block.sqs-block-button").Each(func(i int, s *goquery.Selection) {
		fmt.Println(strings.ToUpper(strings.TrimSpace(s.Find("a").Text())))
		fmt.Println("----------------------------------------------")

		s.Next().Find("p").Each(func(j int, t *goquery.Selection) {
			fmt.Println(t.Text())
		})
		fmt.Println()
	})
}

func StandardFare() {
	doc, err := fetchMenu("https://standardfareberkeley.com/lunch/")
	if err != nil {
		log.Fatal(err)
	}

	meta := doc.Find(".sqs-block.html-block.sqs-block-html")
	fmt.Println(meta.Find("h1:first-of-type strong").Text())
	fmt.Println(meta.Find("h2:first-of-type").Text() + " @ " + meta.Find("H2:nth-of-type(2)").Text())
	fmt.Println(strings.ToUpper(meta.Find("h1:nth-of-type(2)").Text()))
	fmt.Println()

	doc.Find(".menu-item").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".menu-item-title")
		description := s.Find(".menu-item-description")
		altDescription := s.Find(".menu-item-price-bottom")

		if (len(description.Text()) == 0 && len(altDescription.Text()) == 0) {
			fmt.Println(strings.ToUpper(title.Text()))
			fmt.Println("----------------------------------------------")
		} else {
			fmt.Println(title.Text())
			if (len(description.Text()) > 0) {
				fmt.Println("- " + description.Text())
			} else {
				fmt.Println("- " + strings.TrimSpace(altDescription.Text()))
			}
			fmt.Println()
		}
	})
}

func main() {
	StandardFare()
	fmt.Println()
	ArizmendiBakery()
}