package controllers

import "net/http"

func (c Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("method not implemented"))
}

func (c Controller) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("method not implemented"))
}

func (c Controller) ShowAuthenticatedUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("method not implemented"))
}
