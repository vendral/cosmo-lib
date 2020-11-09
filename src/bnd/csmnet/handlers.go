package csmnet

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var indexPath, _ = filepath.Abs("./src/fnd/tmpl/index.html")

var tmpl = template.Must(template.ParseFiles(indexPath))

// IndexHandler handles index page
func IndexHandler(w http.ResponseWriter, req *http.Request) {
	tmpl.Execute(w, nil)
}

// FindHandler handles "find" POST request
func FindHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		req.ParseForm()

		fmt.Println()
		fmt.Println("Cosmo-lib is trying to find a book!")
		fmt.Printf("Author: %s, Title: %s", req.Form["author"], req.Form["title"])
	}
}
