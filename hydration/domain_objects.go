package main

import (
	"time"
)

// Data structure for TrendDB
type Article struct {
	Title   string
	Views   int
	Date    time.Time
	Extract string
}

type Articles []struct {
	Article string `json:"article"`
	Views   int    `json:"views"`
	Extract string `json:"extract"`
}

type ViewsResponse struct {
	Items []struct {
		Year     string   `json:"year"`
		Month    string   `json:"month"`
		Day      string   `json:"day"`
		Articles Articles `json:"articles"`
	} `json:"items"`
}

type WikiResponse struct {
	Query struct {
		Pages map[string]struct {
			Extract string `json:"extract"`
		} `json:"pages"`
	} `json:"query"`
}
