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

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/logout", handler.Logout)
	http.HandleFunc("/", handler.AuthMiddleware(handler.Home))

	http.HandleFunc("/upload", handler.AuthMiddleware(handler.Upload))
	http.HandleFunc("/download", handler.AuthMiddleware(handler.Download))
	http.HandleFunc("/list", handler.AuthMiddleware(handler.List))
	http.HandleFunc("/delete", handler.AuthMiddleware(handler.Delete))

	log.Printf("Starting server on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
