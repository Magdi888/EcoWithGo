package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Magdi888/EcoWithGo/service/user"
	"github.com/gorilla/mux"
)

// GetUserByID godoc
type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Start() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)


	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}