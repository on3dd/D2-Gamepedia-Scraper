package main

import (
	"Dota2-Gamepedia-Scraper/scraper"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var (
	port string = "8080"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/heroes/{name}", HeroHandler)

	fmt.Printf("Server successfully started at port %v...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func HeroHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("Got a request for %v...\n", vars["name"])
	//fmt.Fprintf(w, "Hero: %v\n", vars["name"])

	tmpl := template.Must(template.ParseFiles("templates/template.html"))
	resp := scraper.Scrap(vars["name"])

	defer tmpl.Execute(w, resp)
}
