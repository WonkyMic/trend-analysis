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
	insertPageViewQuery       = "INSERT INTO public.page_views (title, views, date) VALUES ($1, $2, $3)"
	insertArticleSummaryQuery = "INSERT INTO public.article_summary (title, extract) VALUES ($1, $2)"
)

type TrendDB struct {
	db *sql.DB
}

func NewTrendDB() *TrendDB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := os.Getenv("DB_CONN_STR")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to trenddb!")

	return &TrendDB{
		db: db,
	}
}

func (t *TrendDB) close() {
	t.db.Close()
}

func (t *TrendDB) savePageViews(article Article) {
	_, err := t.db.Exec(insertPageViewQuery, article.Title, article.Views, article.Date)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"page_views_unique\"" {
			return
		} else {
			log.Fatal(err)
		}
	}
}

func (t *TrendDB) saveArticleSummary(article Article) {
	_, err := t.db.Exec(insertArticleSummaryQuery, article.Title, article.Extract)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"article_summary_unique\"" {
			return
		} else {
			log.Fatal(err)
		}
	}
}
