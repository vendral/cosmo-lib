package tests

import (
	"testing"
	"net/http"
    "net/http/httptest"
    "html/template"
    "bytes"
	"bnd/csmnet"
)

func TestIndex(t *testing.T) {

    tmpl := template.Must(template.ParseFiles(csmnet.IndexPath))

    var stringTemplate bytes.Buffer
    tmpl.Execute(&stringTemplate, nil)

    t.Run("Test index page", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/index", nil)
        response := httptest.NewRecorder()

		csmnet.IndexHandler(response, request)

        got := response.Body.String()
        want := stringTemplate.String()

        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }
    })
}