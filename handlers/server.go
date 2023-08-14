package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouters() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/notes", GetAllNotes).Methods(http.MethodGet)
	router.HandleFunc("/notes/{id}", GetNote).Methods(http.MethodGet)
	router.HandleFunc("/notes/create", CreateNote).Methods(http.MethodPost)
	router.HandleFunc("/notes/{id}/edit", EditNotes).Methods(http.MethodPut)
	router.HandleFunc("/notes/{id}/delete", DeleteNotes).Methods(http.MethodDelete)

	return router
}
