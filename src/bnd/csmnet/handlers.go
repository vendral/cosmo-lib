package csmnet

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"os"
)

var IndexPath, _ = filepath.Abs(os.Getenv("GOPATH")+"/src/fnd/tmpl/index.html")

var tmpl = template.Must(template.ParseFiles(IndexPath))

// IndexHandler handles index page
func IndexHandler(w http.ResponseWriter, req *http.Request) {
	tmpl.Execute(w, nil)
}

// FindHandler handles "find" POST request
func FindHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		req.ParseForm()

		fmt.Println()
		fmt.Println("Cosmo-lib is trying to find a book!")
		fmt.Printf("Author: %s, Title: %s", req.Form["author"], req.Form["title"])
	}
}
