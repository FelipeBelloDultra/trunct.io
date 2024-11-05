package controllers

import "net/http"

func (c Controller) ShortenURL(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("method not implemented"))
}

func (c Controller) FetchURLs(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("method not implemented"))
}

func (c Controller) ShowURLById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("method not implemented"))
}

func (c Controller) ShowURLStatsById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("method not implemented"))
}

func (c Controller) RedirectToURLByCode(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("method not implemented"))
}
