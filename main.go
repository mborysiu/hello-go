package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving...")
		w.Write([]byte("world!\n"))
	})
	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
