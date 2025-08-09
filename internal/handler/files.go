package handler

import (
	"gonas/internal/utils"
	"log"
	"net/http"
	"os"
	"text/template"
)

const STORAGE = "internal/storage"

func Delete(w http.ResponseWriter, r *http.Request) {

}
func List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	entries, err := os.ReadDir(STORAGE)
	if err != nil {
		log.Fatal(err)
	}

	folders := []string{}
	files := []string{}

	for _, entry := range entries {
		name := entry.Name()

		if entry.IsDir() {
			folders = append(folders, name)
		} else {
			files = append(files, name)
		}
	}

	tmpl, err := template.ParseFiles("web/template.html", "web/list.html")
	if err != nil {
		http.Error(w, "Template parsing error", 500)
		return
	}

	data := utils.PageData{
		Files:   files,
		Folders: folders,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Template execute error:", err)
	}
}

func Download(w http.ResponseWriter, r *http.Request) {

}

func Upload(w http.ResponseWriter, r *http.Request) {

}
