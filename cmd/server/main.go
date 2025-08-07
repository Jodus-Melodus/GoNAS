package main

import (
	"log"
	"net/http"
	"os"

	"gonas/internal/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	http.HandleFunc("/upload", handler.Upload)
	http.HandleFunc("/download", handler.Download)
	http.HandleFunc("/list", handler.List)
	http.HandleFunc("/delete", handler.Delete)

	log.Printf("Starting server on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
