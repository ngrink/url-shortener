package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ngrink/url-shortener/internal/database"
	"github.com/ngrink/url-shortener/internal/modules/auth"
	"github.com/ngrink/url-shortener/internal/modules/urls"
	"github.com/ngrink/url-shortener/internal/modules/users"
	"gorm.io/gorm"
)

type Server struct {
	addr   string
	router *mux.Router
	db     *gorm.DB
}

func NewServer(addr string) *Server {
	server := &Server{
		addr:   addr,
		router: mux.NewRouter(),
		db:     database.DB,
	}

	server.SetupRoutes()
	server.SetupAPIRoutes()

	return server
}

func (s *Server) SetupRoutes() {
	urls.SetupRoutes(s.router)
}

func (s *Server) SetupAPIRoutes() {
	public := s.router.PathPrefix("/api/v1").Subrouter()
	protected := s.router.PathPrefix("/api/v1").Subrouter()

	protected.Use(auth.Authorized)

	users.SetupAPIRoutes(public, protected)
	auth.SetupAPIRoutes(public, protected)
	urls.SetupAPIRoutes(public, protected)
}

func (s *Server) Run() {
	log.Println("Starting server on " + s.addr)
	log.Fatal(http.ListenAndServe(s.addr, s.router))
}
