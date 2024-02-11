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

	http.HandleFunc("/health", health)
	http.HandleFunc("/home", home)
	http.HandleFunc("/article", articleRequest)

	fmt.Println("-- Application Started --")
	http.ListenAndServe(":8080", nil)
}

func health(w http.ResponseWriter, r *http.Request) {
	// return a 200 status code with body "OK" to indicate the application is running
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// Parse the form values to invoke either articleSummaries or articleViews
// based on the form value "index" of either "summaries" or "views"
func articleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		if r.FormValue("index") == "summaries" {
			articleSummaries(w, r)
		} else if r.FormValue("index") == "views" {
			articleViews(w, r)
		}
	}
}

func articleViews(w http.ResponseWriter, r *http.Request) {
	// if the method is POST then print the form values
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(r.Form)
		fmt.Println("TODO: Implement articleViews")
	}
}

func articleSummaries(w http.ResponseWriter, r *http.Request) {
	// if the method is POST then print the form values
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(r.Form)
	}

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
