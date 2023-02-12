package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string
}

// function signature of our handler
type apiFunc func(http.ResponseWriter, *http.Request) error

// decorator function
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
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

	router.HandleFunc("/records", makeHTTPHandleFunc(s.handleRecord))
	log.Println("JSON Api server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleRecord(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetRecord(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateRecord(w, r)
	}

	if r.Method == "DELETE" {
		return s.handleDeleteRecord(w, r)
	}

	return fmt.Errorf("Operation type not permitted %s", r.Method)
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
