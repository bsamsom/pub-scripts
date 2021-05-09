package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"strings"
)

func main() {
	outputDir := "D:\\Insync\\DnD.5e\\Fantasy_Grounds\\modules\\"
	c := colly.NewCollector(colly.MaxBodySize(100 * 1024 * 1024))
	extention := regexp.MustCompile(`.mod`)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		if extention.MatchString(link) {
			//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	c.OnResponse(func(r *colly.Response) {
		//fmt.Println("response received", r.StatusCode)
		filename := strings.Replace(r.FileName(), "Resources_Software_Fantasy_Grounds_Modules_Rulesets_DnD_5E_", "", -1)
		filename = strings.Replace(filename, "_", " ", -1)
		if extention.MatchString(filename) {
			fmt.Println("Saving:", outputDir+filename)
			r.Save(outputDir + filename)
		}
		return
	})

	/*
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})
	*/
	c.Visit("https://thetrove.is/Resources/Software/Fantasy%20Grounds/Modules%20Rulesets/DnD/5E/")

}
