package csmnet

import (
	"html/template"
	"net/http"
	"path/filepath"
	"os"
)

var TestPagePath, _ = filepath.Abs(os.Getenv("GOPATH")+"/src/fnd/tmpl/testpage.html")


// TestPageHandler handles main test page
func TestPageHandler(w http.ResponseWriter, req *http.Request) {
	var tmpl = template.Must(template.ParseFiles(TestPagePath))

	tmpl.Execute(w, nil)
}

// FindHandler handles "find" POST request
func FindHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		req.ParseForm()

		Info.Println("/find request is executed")
		Info.Printf("/find request params: author: %s, title: %s", req.Form["author"], req.Form["title"])

		// Go back to TestPage
		TestPageHandler(w, req)
	}
}
