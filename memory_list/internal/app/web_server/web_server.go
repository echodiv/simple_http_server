package webserver

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/echodiv/simple_server/memory_list/internal/app/storage"
)

type Server interface {
	New(string) *Server
	Start() error
	configureRouter()
}

type WebServer struct {
	adress  string
	router  *mux.Router
	storage *storage.Storage
}

// Create instance of web application
func NewWebServer(address string) *WebServer {
	log.Printf(">>Create new WebServer")
	return &WebServer{
		adress:  address,
		router:  mux.NewRouter(),
		storage: storage.NewStorage(),
	}
}

// Start Web server
func (s *WebServer) Start() error {
	log.Printf(">>Start new WebServer")
	s.configureRouter()
	return http.ListenAndServe(s.adress, s.router)

}

// Handler registration
func (s *WebServer) configureRouter() {
	s.router.HandleFunc("/", s.MainResponse())
	s.router.HandleFunc("/name/{name}", s.CreateElement()).Methods("POST")
	s.router.HandleFunc("/name/{name}", s.GetElementByName()).Methods("GET")
}
