package handler

import (
	"fmt"
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
		fmt.Println("Error reading file: ", err)
		return
	}

	w.Write(data)

	// fmt.Println(r.Method)     // Type(verb)
	// fmt.Println(r.Host)       // server ip (localhost)
	// fmt.Println(r.RemoteAddr) // client addr
	// fmt.Println(r.URL)        // rest of path (everything after ip and port)
}
