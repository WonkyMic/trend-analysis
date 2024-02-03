package main

import (
	"fmt"
	"log"
	"time"
)

const (
	breakpoint = 99
)

func main() {
	data, err := GetDailyPageviewsForPreviousNumberOfDays(7)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Retrieve extract for each article
	for _, d := range data {
		for i, item := range d.Items {
			d.Items[i].Articles, err = GetArticlesExtract(item.Articles)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}

	saveData(data)
	// printData(data)
	fmt.Println("--Hydration Complete--")
}

func printData(data []ViewsResponse) {
	for _, d := range data {
		for _, item := range d.Items {
			fmt.Printf("Date: %s/%s/%s\n", item.Year, item.Month, item.Day)
			for i, article := range item.Articles {
				if i >= breakpoint {
					break
				}
				fmt.Println(article.Article, article.Views)
				// fmt.Println(article.Extract)
			}
		}
	}
}

func saveData(data []ViewsResponse) {
	tdb := NewTrendDB()
	defer tdb.close()
	for _, d := range data {
		for _, item := range d.Items {
			for i, article := range item.Articles {
				if i >= breakpoint {
					break
				}
				// parse the date string
				date, err := time.Parse("2006-01-02", fmt.Sprintf("%s-%s-%s", item.Year, item.Month, item.Day))
				if err != nil {
					log.Fatal(err)
				}

				tdb.savePageViews(Article{
					Title: article.Article,
					Views: article.Views,
					Date:  date,
				})

				tdb.saveArticleSummary(Article{
					Title:   article.Article,
					Extract: article.Extract,
				})

			}
		}
	}
}
