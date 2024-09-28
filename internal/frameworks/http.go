package frameworks

import (
	"log"
	"net/http"
	"test/internal/handlers"
)

type Server struct {
	mux http.Handler
}

func NewServer() *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/song", handlers.SongHandler)
	mux.HandleFunc("/songs/", func(w http.ResponseWriter, r *http.Request) {})
	return &Server{mux: mux}

}

func (s *Server) Start() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: s.mux,
	}

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
