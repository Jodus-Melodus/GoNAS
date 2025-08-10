package handler

import (
	"gonas/internal/utils"
	"log"
	"net/http"
	"text/template"
)

const PATH = "web/home.html"

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("web/template.html", "web/home.html")
	if err != nil {
		http.Error(w, "Template parsing error", 500)
		return
	}

	session, _ := store.Get(r, "gonas-session")
	data := utils.PageData{
		Authenticated: session.Values["authenticated"] == true,
		Files:         nil,
		Folders:       nil,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Template execute error:", err)
	}
}
