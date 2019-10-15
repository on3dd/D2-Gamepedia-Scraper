package main

import (
	"Dota2-Gamepedia-Scraper/scraper"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"time"
)

var (
	port string = "8080"
)

// TODO: Fix the bug with Lina and Crystal Maiden audio playback
// TODO: Add 404 page
// TODO: Add an exception for pages with the broken audios

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/heroes/{name}", heroHandler)
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	fmt.Printf("%v: Server successfully started at port %v...\n", time.Now().Format(time.UnixDate), port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

}

func heroHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("%v: Got a request for %v...\n", time.Now().Format(time.UnixDate),  vars["name"])

	tmpl := template.Must(template.ParseFiles("templates/hero.html"))
	resp := scraper.Scrap(vars["name"])

	tmpl.Execute(w, resp)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v: Got invalid page request.\n", time.Now().Format(time.UnixDate))

	tmpl := template.Must(template.ParseFiles("templates/404.html"))
	tmpl.Execute(w, nil)
}