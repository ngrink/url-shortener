package app

import (
	"log"
	"mime"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ngrink/url-shortener/internal/database"
	"github.com/ngrink/url-shortener/internal/modules/auth"
	"github.com/ngrink/url-shortener/internal/modules/urls"
	"github.com/ngrink/url-shortener/internal/modules/users"
	"github.com/ngrink/url-shortener/web"
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
	server.ServeStatic()

	return server
}

func (s *Server) SetupRoutes() {
	web.SetupRoutes(s.router)
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

func (s *Server) ServeStatic() {
	mime.AddExtensionType(".css", "text/css")
	mime.AddExtensionType(".js", "application/javascript")

	fs := http.FileServer(http.Dir("./web/assets/"))
	s.router.PathPrefix("/assets").Handler(
		http.StripPrefix("/assets", fs),
	)
}

func (s *Server) Run() {
	log.Println("Starting server on " + s.addr)
	log.Fatal(http.ListenAndServe(s.addr, s.router))
}
