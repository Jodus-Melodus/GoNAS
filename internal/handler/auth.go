package handler

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

var (
	store = sessions.NewCookieStore([]byte("super-secret-key-1234567890abcdef"))
)

func Login(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/template.html", "web/login.html")
	if err != nil {
		http.Error(w, "Template parsing error", 500)
		return
	}

	if r.Method == http.MethodPost {
		user := r.FormValue("username")
		pass := r.FormValue("password")

		if user == "admin" && pass == "password123" {
			session, _ := store.Get(r, "gonas-session")
			session.Values["authenticated"] = true
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Template execute error:", err)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	session, _ := store.Get(r, "gonas-session")
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "gonas-session")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}
