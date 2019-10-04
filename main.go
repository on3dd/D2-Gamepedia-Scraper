package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type Response struct {
	Categories *[]Category
}

type Category struct {
	Header string
	Subcategories *[]Subcategory
}

type Subcategory struct {
	Header string
	Links *[]Link
}

type Link struct {
	Title string
	Link string
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("dota2.gamepedia.com"), )

	c.OnHTML("h2 > span[class=mw-headline]", func(e *colly.HTMLElement) {
		getResponses(e)
	})

	c.Visit("https://dota2.gamepedia.com/Pudge/Responses")
}

func getResponses(e *colly.HTMLElement) {
	// Print category header
	category := "- " + e.Text
	fmt.Println(category)

	el := e.DOM.Parent().Next()

	// While next element is a "p" or "ul" tag
	for !(el.Nodes[0].Data == "h2" || el.Nodes[0].Data == "table") {
		// If find "p" then print subcategory header
		// Otherwise parse ul
		if el.Nodes[0].Data == "p" {
			// Getting data from "b" tag if it's necessary
			text := el.Find("b").Text()
			if text != "" {
				if strings.Split(text, "")[0] == " " {
					text = strings.Join(strings.Split(text, "")[1:], "")
				}
				fmt.Println("-- " + text)
			}
		} else {
			for _, item := range el.Find("li").Nodes {
				text := "---" + item.LastChild.Data
				// Fix possible problems
				if text == "---i" || text == "---b" {
					text = "--- " + item.LastChild.FirstChild.Data
				} else if text == "---!" {
					text = "--- Shitty wizard!"
				}
				// Select attr val from span > audio > source
				src := item.FirstChild.FirstChild.FirstChild.Attr[0].Val
				fmt.Println(text, src)
			}
		}
		el = el.Next()
	}
	fmt.Println()
}