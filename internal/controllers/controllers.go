package controllers

import (
	"encoding/json"
	"net/http"
	"test/internal/models"
	"test/internal/store"
)

func GetSongs(w http.ResponseWriter, r *http.Request) {
	var songs []models.Song
	store.Db.Find(&songs)
	json.NewEncoder(w).Encode(songs)
}
