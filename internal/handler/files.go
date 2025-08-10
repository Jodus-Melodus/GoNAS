package handler

import (
	"encoding/json"
	"fmt"
	"gonas/internal/utils"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

const STORAGE = "internal/storage"

func Delete(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request utils.DeleteRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil || request.Name == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	path := "internal/storage/" + request.Name
	info, err := os.Stat(path)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	if info.IsDir() {
		err = os.RemoveAll(path)
	} else {
		err = os.Remove(path)
	}

	if err != nil {
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
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

	session, _ := store.Get(r, "gonas-session")
	data := utils.PageData{
		Authenticated: session.Values["authenticated"] == true,
		Files:         files,
		Folders:       folders,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Template execute error:", err)
	}
}

func Download(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	filename := r.URL.Query().Get("filename")
	if filename == "" {
		http.Error(w, "Missing file name", http.StatusBadRequest)
		return
	}

	safePath := filepath.Join("internal/storage", filepath.Clean(filename))

	file, err := os.Open(safePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	f, err := file.Stat()
	if err != nil {
		http.Error(w, "Could not get file info", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=\""+f.Name()+"\"")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", f.Size()))

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Failed to send file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Upload(w http.ResponseWriter, r *http.Request) {

}
