package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello Go")
	})

	http.HandleFunc("/go", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello GoLang, welcome to My life")
	})

	http.ListenAndServe(":9090", nil)
}
