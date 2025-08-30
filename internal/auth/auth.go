package auth

import (
	"encoding/json"
	"gonas/global"
	"gonas/internal/utils"
	"net/http"
	"os"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

const usersFile = "internal/auth/users.json"

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, _ := template.ParseFiles("web/template.html", "web/login.html")
		tmpl.Execute(w, nil)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	var users []utils.User
	data, err := os.ReadFile(usersFile)
	if err != nil {

	}
	json.Unmarshal(data, &users)

	for _, u := range users {
		if u.Username == username && checkPasswordHash(password, u.Password) {
			session, _ := global.Store.Get(r, "gonas-session")
			session.Values["authenticated"] = true
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)

}

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	session, _ := global.Store.Get(r, "gonas-session")
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("web/template.html", "web/register.html")
		if err != nil {
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Missing fields", http.StatusBadRequest)
		return
	}
	hash, err := hashPassword(password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	var users []utils.User
	data, err := os.ReadFile(usersFile)
	if err != nil && !os.IsNotExist(err) {
		http.Error(w, "Error reading users file", http.StatusInternalServerError)
		return
	}

	if len(data) > 0 {
		if err := json.Unmarshal(data, &users); err != nil {
			http.Error(w, "Error parsing users file", http.StatusInternalServerError)
			return
		}
	}

	users = append(users, utils.User{Username: username, Password: hash})

	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		http.Error(w, "Error encoding users", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile(usersFile, out, 0644)
	if err != nil {
		http.Error(w, "Error writing users file", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := global.Store.Get(r, "gonas-session")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
