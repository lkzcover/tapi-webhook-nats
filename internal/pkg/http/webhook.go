package http

import (
	"io"
	"log"
	"net/http"
)

func (s *Server) webhook(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	if s.debugMode {
		log.Println(string(body))
	}

	if err := s.nc.Publish(s.ncSubj, body); err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
