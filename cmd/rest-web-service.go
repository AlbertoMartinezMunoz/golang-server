package main

import (
	"cmd/rest-web-service/pkg/catalogue"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gosimple/slug"
)

func main() {
	fmt.Printf("************ Native HTTP REST Web Service ************\n")

	// Create the Store and Recipe Handler
	store := catalogue.NewMemStore()
	catalogueHandler := NewCatalogueHandler(store)

	mux := http.NewServeMux()

	mux.Handle("/", &homeHandler{})
	mux.Handle("/catalogue", catalogueHandler)
	mux.Handle("/catalogue/", catalogueHandler)

	http.ListenAndServe(":8080", mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page."))
}

type catalogueHandler struct {
	store catalogue.CatalogueStore
}

func NewCatalogueHandler(s catalogue.CatalogueStore) *catalogueHandler {
	return &catalogueHandler{store: s}
}

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
	var album catalogue.Album

	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		InternalServerErrorHandler(w, r)
		fmt.Printf("Error decoding json: '%v'", err)
		return
	}

	albumID := slug.Make(album.Title)

	if err := h.store.Add(albumID, album); err != nil {
		fmt.Printf("Error adding album: '%v' - '%v'", err, album)
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *catalogueHandler) ListAlbums(w http.ResponseWriter, r *http.Request) {
	resources, err := h.store.List()

	jsonBytes, err := json.Marshal(resources)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *catalogueHandler) GetAlbum(w http.ResponseWriter, r *http.Request) {
	// Extract the resource ID/slug using a regex
    matches := RecipeReWithID.FindStringSubmatch(r.URL.Path)
    // Expect matches to be length >= 2 (full string + 1 matching group)
    if len(matches) < 2 {
        InternalServerErrorHandler(w, r)
        return
    }

    // Retrieve recipe from the store
    recipe, err := h.store.Get(matches[1])
    if err != nil {
        // Special case of NotFound Error
        if err == catalogue.NotFoundErr {
			fmt.Printf("Error getting album: '%v' - '%v'", err, matches[1])
            NotFoundHandler(w, r)
            return
        }

        // Every other error
        InternalServerErrorHandler(w, r)
        return
    }

    // Convert the struct into JSON payload
    jsonBytes, err := json.Marshal(recipe)
    if err != nil {
        InternalServerErrorHandler(w, r)
        return
    }

    // Write the results
    w.WriteHeader(http.StatusOK)
    w.Write(jsonBytes)
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

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}
