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

	data := utils.PageData{
		Files: []string{},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Template execute error:", err)
	}
}
