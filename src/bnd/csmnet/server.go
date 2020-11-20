package csmnet

import (
	"fmt"
	"net/http"
)

type server struct {
	ServerPort       int
	StringServerPort string
	Mux              *http.ServeMux
}

var ()

func serverInit() server {

	var srv server

	srv.ServerPort = 8090
	srv.StringServerPort = fmt.Sprintf("%s%v", ":", srv.ServerPort)

	srv.Mux = http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	srv.Mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	srv.Mux.HandleFunc("/testpage", TestPageHandler)
	srv.Mux.HandleFunc("/find", FindHandler)

	return srv
}

// ServerStart starts main server
func ServerStart() {

	var srv = serverInit()

	fmt.Println("[CosmoBackendInfo]: Server started and listening on port:", srv.ServerPort)
	fmt.Printf("[CosmoBackendInfo]: http://localhost%s/testpage", srv.StringServerPort)

	http.ListenAndServe(srv.StringServerPort, srv.Mux)
}
