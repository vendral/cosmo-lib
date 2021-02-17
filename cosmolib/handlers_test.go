package main

import (
	"bytes"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {

	tmpl := template.Must(template.ParseFiles(TestPagePath))

	var stringTemplate bytes.Buffer
	tmpl.Execute(&stringTemplate, nil)

	t.Run("Testpage page", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/testpage", nil)
		response := httptest.NewRecorder()

		TestPageHandler(response, request)

		got := response.Body.String()
		want := stringTemplate.String()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
