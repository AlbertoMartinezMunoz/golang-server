package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func main() {
	fmt.Printf("************ Native HTTP REST Web Service ************\n")

	mux := http.NewServeMux()

	mux.Handle("/", &homeHandler{})
	mux.Handle("/catalogue", &catalogueHandler{})
	mux.Handle("/catalogue/", &catalogueHandler{})

	http.ListenAndServe(":8080", mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page."))
}

type catalogueHandler struct{}

func (h *catalogueHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received Request: %v, '%v'\n", r.Method, r.URL.Path)
	switch {
	case r.Method == http.MethodPost && RecipeRe.MatchString(r.URL.Path):
		h.CreateAlbum(w, r)
		return
	case r.Method == http.MethodGet && RecipeRe.MatchString(r.URL.Path):
		h.ListAlbums(w, r)
		return
	case r.Method == http.MethodGet && RecipeReWithID.MatchString(r.URL.Path):
		h.GetAlbum(w, r)
		return
	case r.Method == http.MethodGet && RecipeReWithID.MatchString(r.URL.Path):
		h.GetAlbum(w, r)
		return
	case r.Method == http.MethodPut && RecipeReWithID.MatchString(r.URL.Path):
		h.UpdateAlbum(w, r)
		return
	case r.Method == http.MethodDelete && RecipeReWithID.MatchString(r.URL.Path):
		h.DeleteAlbum(w, r)
		return
	default:
		w.Write([]byte("This is the catalogue root page."))
		return
	}
}

func (h *catalogueHandler) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Album Request."))
}

func (h *catalogueHandler) ListAlbums(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List Albums Request."))
}

func (h *catalogueHandler) GetAlbum(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Album Request."))
}

func (h *catalogueHandler) UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Album Request."))
}

func (h *catalogueHandler) DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Album Request."))
}

var (
	RecipeRe       = regexp.MustCompile(`^/catalogue/*$`)
	RecipeReWithID = regexp.MustCompile(`^/catalogue/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
)
