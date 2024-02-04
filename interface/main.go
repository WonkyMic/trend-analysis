package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var trendDB *TrendDB

func main() {
	trendDB = NewTrendDB()

	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/dist/", http.StripPrefix("/dist/", fs))

	http.HandleFunc("/", root)
	http.HandleFunc("/article/summaries", summaries)

	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "index.html")
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func summaries(w http.ResponseWriter, r *http.Request) {
	trendDB.open()
	defer trendDB.close()

	articles, err := trendDB.selectArticleSummaries()

	if err != nil {
		fmt.Println("Error selecting article summaries")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("./templates/table.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err = tmpl.Execute(w, articles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
