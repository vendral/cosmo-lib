package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	ServerPort = 8090
)

var tmpl = template.Must(template.ParseFiles("tmpl/index.html"))

func indexHandler(w http.ResponseWriter, req *http.Request) {
	tmpl.Execute(w, nil)
}

func findHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		req.ParseForm()

		fmt.Println()
		fmt.Println("Cosmo-lib is trying to find a book!")
		fmt.Printf("Author: %s, Title: %s", req.Form["author"], req.Form["title"])
	}
}

func main() {

	StringServerPort := fmt.Sprintf("%s%v", ":", ServerPort)

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/find", findHandler)

	fmt.Println("Cosmo-lib is listening on port:", ServerPort)
	fmt.Printf("Go to Cosmo-lib: http://localhost%s", StringServerPort)

	http.ListenAndServe(StringServerPort, mux)
}
