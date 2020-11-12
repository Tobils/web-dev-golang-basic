package main

import (
	"log"
	"net/http"
)

func main()  {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloWorldHandler)

	log.Println("Starting web on port 8081")

	err := http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello world, saya sedang belajar Golang web"))
}