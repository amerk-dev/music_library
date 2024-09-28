package handlers

import (
	"net/http"
)

func SongHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("Method - " + r.Method))
	case "POST":
		w.Write([]byte("Method - " + r.Method))
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
