package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

// type item struct {
// 	Testreport  string `json:"test_reported"`
// 	Activecases string `json:"active_cases"`
// }

type Item struct {
	Testreport  string 
	Activecases string 
}

func main() {
	c := colly.NewCollector()

	items := Item{}
	c.OnHTML("div.information_block", func(h *colly.HTMLElement) {
			items.Testreport =   h.ChildText("div.testing_result > strong.testing_count") 
			items.Activecases =  h.ChildText("div.block-active-cases > span.icount") 
	})

	c.OnScraped(func (r *colly.Response) {
		enc := json.NewEncoder(os.Stdout)
		out := enc.Encode(items)
		fmt.Println(out)
	})
	c.Visit("https://www.mygov.in/covid-19")


}
