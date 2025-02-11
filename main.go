package main

import (
	"fmt"
	"net/http"
)

func helloMir(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Привет")
}

func main() {
	router := http.NewServeMux()

	server := http.Server{
		Addr:    ":8888",
		Handler: router,
	}

	router.HandleFunc("/", helloMir)

	server.ListenAndServe()
}
