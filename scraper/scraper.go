package scraper

import (
	_ "fmt"
	"github.com/gocolly/colly"
	"strings"
)

type Response struct {
	Header     string
	Categories []*Category
}

type Category struct {
	Header        string
	Subcategories []*Subcategory
}

type Subcategory struct {
	Header string
	Links  []*Link
}

type Link struct {
	Title string
	Link  string
}

// Scrap takes hero name and scraps gamepedia for hero's data
func Scrap(heroName string) *Response {
	c := colly.NewCollector(
		colly.AllowedDomains("dota2.gamepedia.com"))

	heroName = strings.Replace(heroName, "_", " ", -1)
	resp := &Response{
		Header:     heroName,
		Categories: make([]*Category, 0)}

	c.OnHTML("h2 > span[class=mw-headline]", func(e *colly.HTMLElement) {
		// Get the category instance and append it to response
		ct := getResponses(e)
		resp.Categories = append(resp.Categories, ct)

		// M
	})

	url := "https://dota2.gamepedia.com/" + heroName + "/Responses"
	err := c.Visit(url)
	if err != nil {
		panic("Error: " + err.Error())
	}

	return resp
}

// Get the heroes replicas from selection
func getResponses(e *colly.HTMLElement) *Category {
	if e.Text == "Rylai's Battle Blessing" {
		return &Category{}
	}

	ct := &Category{
		Header:        e.Text,
		Subcategories: make([]*Subcategory, 0)}

	// Print category header
	//fmt.Println("- " + ct.Header)

	el := e.DOM.Parent().Next()

	// While next element is a "p" or "ul" tag
	for !(el.Nodes[0].Data == "h2" || el.Nodes[0].Data == "table") {
		// If find "p" then print subcategory header
		// Otherwise parse ul
		if el.Nodes[0].Data == "p" {
			// Getting data from "b" tag if it's necessary
			text := el.Find("b").Text()

			// If text isn't empty then create a new subcategory
			if text != "" {
				// If first letter is whitespace delete it
				if strings.Split(text, "")[0] == " " {
					text = strings.Join(strings.Split(text, "")[1:], "")
				}

				if text == "Aghanim's Scepter" {
					el = el.Next().Next()
				}

				// Get the num of the subcategories
				//subCategoriesNum := len(ct.Subcategories)

				// Put new subcategory in a category
				ct.Subcategories = append(ct.Subcategories, &Subcategory{
					Header: text,
					Links:  make([]*Link, 0)})

				//fmt.Println("-- " + ct.Subcategories[subCategoriesNum].Header)
			}
		} else {
			// Iteratively go through "li" nodes
			for _, item := range el.Find("li").Nodes {
				// Get the text value of element
				text := item.LastChild.Data

				// Fix possible problems
				if text == "i" || text == "b" {
					text = item.LastChild.FirstChild.Data
				} else if text == "!" {
					text = "Shitty wizard!"
				} else {
					text = strings.Join(strings.Split(text, "")[1:], "")
				}

				// Get the link from attr value from span > audio > source
				link := item.FirstChild.FirstChild.FirstChild.Attr[0].Val

				// Get the num of subcategories in the category
				subCategoriesNum := len(ct.Subcategories)

				// If there is no subcategories in category
				// then put a default subcategory
				if subCategoriesNum == 0 {
					ct.Subcategories = append(ct.Subcategories, &Subcategory{
						Header: "default header",
						Links:  make([]*Link, 0),
					})
					subCategoriesNum = 1
				}

				sb := ct.Subcategories[subCategoriesNum-1]

				// Get the num of the links in the subcategory
				//linksNum := len(sb.Links)

				// Put the new link in a subcategory
				sb.Links = append(sb.Links, &Link{
					Title: text,
					Link:  link,
				})

				//fmt.Println("--- "+sb.Links[linksNum].Title, sb.Links[linksNum].Link)
			}
		}
		el = el.Next()
	}
	//fmt.Println()
	return ct
}
