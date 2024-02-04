package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	selectArticleSummariesQuery = "SELECT title, summary FROM public.article_summary"
)

type TrendDB struct {
	connStr string
	db      *sql.DB
}

func NewTrendDB() *TrendDB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := os.Getenv("DB_CONN_STR")

	return &TrendDB{
		connStr: connStr,
		db:      nil,
	}
}

func (t *TrendDB) open() {
	db, err := sql.Open("postgres", t.connStr)
	if err != nil {
		log.Fatal(err)
	}
	t.db = db
}

func (t *TrendDB) close() {
	t.db.Close()
}

func (t *TrendDB) selectArticleSummaries() ([]ArticleSummary, error) {
	rows, err := t.db.Query(selectArticleSummariesQuery)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var articles []ArticleSummary
	for rows.Next() {

		var articleSummary ArticleSummary
		err = rows.Scan(&articleSummary.Title, &articleSummary.Summary)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		articles = append(articles, articleSummary)
	}

	return articles, nil
}
