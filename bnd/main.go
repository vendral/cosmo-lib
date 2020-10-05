package main

import (
	"fmt"
	"net/http"
)

var (
	ServerPort = 8090
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello\n")
}

func main() {

	StringServerPort := fmt.Sprintf("%s%v", ":", ServerPort)

	http.HandleFunc("/cosmo", hello)

	fmt.Println("Cosmo-lib is listening on port:", ServerPort)
	fmt.Printf("Go to Cosmo-lib: http://localhost%s", StringServerPort)

	http.ListenAndServe(StringServerPort, nil)
}
