package handler

import (
	"log"
	"net/http"
	"os"
)

const PATH = "web/home.html"

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	data, err := os.ReadFile(PATH)
	if err != nil {
		log.Println("Error reading file: ", err)
		return
	}

	w.Write(data)
}
