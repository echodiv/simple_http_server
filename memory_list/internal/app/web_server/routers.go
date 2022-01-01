package webserver

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Main page
func (s *WebServer) MainResponse() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("---> inside main router")
		io.WriteString(rw, "Hello")
	}
}

// Hendler for get list element from storage by name
func (s *WebServer) GetElementByName() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("---> inside get element router")
		vars := mux.Vars(r)
		val, err := s.storage.GetElementByName(vars["name"])
		if err != nil {
			io.WriteString(rw, err.Error())
			return
		}
		if r, err := json.Marshal(val); err != nil {
			io.WriteString(rw, err.Error())
			return
		} else {
			rw.WriteHeader(http.StatusOK)

			io.WriteString(rw, string(r))
		}
	}
}

// Handler for create element in server storage
func (s *WebServer) CreateElement() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("---> inside create element router")
		vars := mux.Vars(r)
		if err := s.storage.AddNewElement(vars["name"]); err != nil {
			log.Println(err.Error())
			io.WriteString(rw, err.Error())
			return
		}
		rw.WriteHeader(http.StatusOK)
	}
}
