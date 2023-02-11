package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			//Handle error
		}
	}
}

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/records", s.handleRecord())
}

func (s *APIServer) handleRecord(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetRecord(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateRecord(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteRecord(w http.ResponseWriter, r *http.Request) error {
	return nil
}
