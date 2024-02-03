package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetArticleExtract(title string) (string, error) {
	// Define the URL of the API endpoint
	// exintro: Return only content before the first section
	baseURL := "https://en.wikipedia.org/w/api.php"
	params := "?format=json&action=query&prop=extracts&exlimit=max&explaintext&exintro&titles=" + url.QueryEscape(title)

	// Send a GET request to the API endpoint
	resp, err := http.Get(baseURL + params)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse the response body as JSON
	var data WikiResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	// Extract the article content from the parsed JSON data
	for _, page := range data.Query.Pages {
		if page.Extract != "" {
			return page.Extract, nil
		} else {
			return "", fmt.Errorf("no extract found for the given title: %s", title)
		}
	}

	return "", fmt.Errorf("no extract found for the given title")
}

func GetArticlesExtract(articles Articles) (Articles, error) {
	for i, article := range articles {
		if i >= breakpoint {
			break
		}
		extract, err := GetArticleExtract(article.Article)
		if err != nil {
			return Articles{}, err
		}
		articles[i].Extract = extract
	}

	return articles, nil
}
