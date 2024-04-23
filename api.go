package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (s *APIServer) ListenAndServe() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// registering our services
	taskService := NewTasksService(s.store)
	taskService.RegisterRoutes(subrouter)

	userService := NewUserService(s.store)
	userService.RegisterRoutes(subrouter)

	log.Printf("Listening api server on %s\n", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
