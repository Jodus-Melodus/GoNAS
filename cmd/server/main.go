package main

import (
	"log"
	"net/http"
	"os"

	"gonas/global"
	"gonas/internal/auth"
	"gonas/internal/handler"

	"github.com/gorilla/sessions"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	global.Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	http.HandleFunc("/login", auth.Login)
	http.HandleFunc("/logout", auth.Logout)
	http.HandleFunc("/register", auth.Register)
	http.HandleFunc("/", auth.AuthMiddleware(handler.Home))

	http.HandleFunc("/upload", auth.AuthMiddleware(handler.Upload))
	http.HandleFunc("/download", auth.AuthMiddleware(handler.Download))
	http.HandleFunc("/list", auth.AuthMiddleware(handler.List))
	http.HandleFunc("/delete", auth.AuthMiddleware(handler.Delete))

	log.Printf("Starting server on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
