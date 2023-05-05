package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	serverAddress = "0.0.0.0:8080"
)

func main() {
	flag.StringVar(&serverAddress, "server-address", serverAddress, "")
	flag.Parse()

	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/-/healthz", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`UP`))
	})
	server := &http.Server{
		Addr:    serverAddress,
		Handler: serverMux,
	}

	log.Printf("server listening on %s", server.Addr)
	server.ListenAndServe()
}
