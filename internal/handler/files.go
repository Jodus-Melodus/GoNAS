package handler

import (
	"fmt"
	"gonas/global"
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
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := filepath.Clean(r.URL.Query().Get("path"))

	baseDir := STORAGE
	if path != STORAGE {
		baseDir = filepath.Join(baseDir, path)
	}

	fmt.Println(baseDir)

	info, err := os.Stat(baseDir)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	if info.IsDir() {
		err = os.RemoveAll(baseDir)
	} else {
		err = os.Remove(baseDir)
	}

	if err != nil {
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/list", http.StatusSeeOther)
	w.WriteHeader(http.StatusOK)
}

func List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	directoryName := filepath.Clean(r.URL.Query().Get("directory"))

	folders := []utils.DirectoryInfo{}
	files := []utils.FileInfo{}
	baseDir := STORAGE
	if directoryName != STORAGE {
		baseDir = filepath.Join(baseDir, directoryName)
	}

	directories, err := os.ReadDir(baseDir)
	if err != nil {
		http.Error(w, "Error reading directories", http.StatusInternalServerError)
		return
	}

	for _, entry := range directories {
		fullPath := filepath.Join(directoryName, entry.Name())

		if entry.IsDir() {
			folders = append(folders, utils.DirectoryInfo{
				Name: entry.Name(),
				Path: fullPath,
			})
		} else {
			files = append(files, utils.FileInfo{
				Name: entry.Name(),
				Path: fullPath,
			})
		}
	}

	tmpl, err := template.ParseFiles("web/template.html", "web/list.html")
	if err != nil {
		http.Error(w, "Template parsing error", 500)
		return
	}

	session, _ := global.Store.Get(r, "gonas-session")

	if directoryName == "" {
		folders = nil
	}

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

	safePath := filepath.Join(STORAGE, filepath.Clean(filename))

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
	if r.Method == http.MethodGet {
		tmpl, _ := template.ParseFiles("web/template.html", "web/upload.html")
		tmpl.Execute(w, nil)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file from form", http.StatusBadRequest)
		return
	}

	destinationPath := filepath.Join(STORAGE, handler.Filename)
	destination, err := os.Create(destinationPath)
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}
	defer destination.Close()

	_, err = io.Copy(destination, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
