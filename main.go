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

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/heroes/{name}", HeroHandler)

	fmt.Printf("%v: Server successfully started at port %v...\n", time.Now().Format(time.UnixDate), port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func HeroHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("%v: Got a request for %v...\n", time.Now().Format(time.UnixDate),  vars["name"])

	tmpl := template.Must(template.ParseFiles("templates/template.html"))
	resp := scraper.Scrap(vars["name"])

	tmpl.Execute(w, resp)
}