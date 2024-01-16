package main

import (
	"fmt"
)

const (
	breakpoint = 3
)

func main() {
	data, err := GetDailyPageviewsForPreviousNumberOfDays(7)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, d := range data {
		for i, item := range d.Items {
			d.Items[i].Articles, err = GetArticlesExtract(item.Articles)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	printData(data)
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
