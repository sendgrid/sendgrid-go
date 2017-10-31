package controllers

import (
	"log"
	"net/http"
)

func ServeInternalServerError(w http.ResponseWriter, err error) {
	log.Print(err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func ServeStatusForbiddenError(w http.ResponseWriter, err error) {
	log.Print(err)
	http.Error(w, "Forbidden", http.StatusForbidden)
}

func RecordNotFound(w http.ResponseWriter, err error) {
	log.Print(err)
	http.Error(w, "Record not found", http.StatusNotFound)
}

func ServeBadRequest(w http.ResponseWriter, err error) {
	log.Print(err)
	http.Error(w, "Bad Request", http.StatusBadRequest)
}
