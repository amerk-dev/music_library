package handlers

import (
	"net/http"
	"test/internal/controllers"
)

func SongHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		controllers.GetSongs(w, r)
	case "POST":
		w.Write([]byte("Method - " + r.Method))
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
