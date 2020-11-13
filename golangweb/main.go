package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About Page"))
	}

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloWorldHandler)
	mux.HandleFunc("/suhada", handler.SuhadaHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/product", handler.ProductHandler)

	// dengan anonimous function
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is Profile"))
	})

	log.Println("Starting web on port 8081")

	err := http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}
