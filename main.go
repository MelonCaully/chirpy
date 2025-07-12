package main

import "net/http"

func main() {
	serve := http.NewServeMux()
	server := http.Server{
		Handler: serve,
		Addr:    ":8080",
	}
	serve.Handle("/", http.FileServer(http.Dir(".")))
	server.ListenAndServe()

}
