package handler

import (
	"log"
	"net/http"
	"os"
	"strings"
)

const STORAGE = "internal/storage"

func Delete(w http.ResponseWriter, r *http.Request) {

}
func List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html")

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

	fileStr := strings.Join(files, "\n")
	folderStr := strings.Join(folders, "\n")
	w.Write([]byte(fileStr))
	w.Write([]byte(folderStr))
}

func Download(w http.ResponseWriter, r *http.Request) {

}

func Upload(w http.ResponseWriter, r *http.Request) {

}
