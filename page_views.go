package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// GetDailyPageviews returns the top 1000 most viewed articles on Wikipedia for the given date
// The date must be in the format YYYY/MM/DD
// The returned data is filtered to remove common titles such as "Main_Page", "Special:Search", and "Wikipedia:Featured_pictures"
// NOTE: The API endpoint doesn't seem to return data for current dates, so you'll have to use a date in the past
func GetDailyPageviews(date string) (ViewsResponse, error) {
	// Define the URL of the API endpoint
	url := "https://wikimedia.org/api/rest_v1/metrics/pageviews/top/en.wikipedia/all-access/" + date

	// Send a GET request to the API endpoint
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ViewsResponse{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ViewsResponse{}, err
	}

	// Parse the response body as JSON
	var data ViewsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return ViewsResponse{}, err
	}

	return RemoveCommonTitles(data), nil
}

func GetDailyPageviewsForPreviousNumberOfDays(days int) ([]ViewsResponse, error) {
	// Get the current date
	currentDate := time.Now()

	// Get the daily pageviews for the last 7 days
	var data []ViewsResponse
	for i := 1; i <= days; i++ {
		date := currentDate.AddDate(0, 0, -i).Format("2006/01/02")

		d, err := GetDailyPageviews(date)
		if err != nil {
			fmt.Println(err)
			return []ViewsResponse{}, err
		}

		data = append(data, d)
	}

	return data, nil
}

func RemoveCommonTitles(data ViewsResponse) ViewsResponse {
	for i, item := range data.Items {

		// Create a new slice to hold the filtered articles
		articles := Articles{}

		for _, article := range item.Articles {
			if article.Article == "Main_Page" {
				continue
			}
			if article.Article == "Special:Search" {
				continue
			}
			if article.Article == "Wikipedia:Featured_pictures" {
				continue
			}

			articles = append(articles, article)
		}

		// Replace the original articles with the filtered articles
		data.Items[i].Articles = articles
	}

	return data
}
