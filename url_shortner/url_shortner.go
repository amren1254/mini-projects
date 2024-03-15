package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// URLMap stores the mapping of short URLs to long URLs
type URLMap struct {
	mu      sync.Mutex
	entries map[string]string
}

func NewURLMap() *URLMap {
	return &URLMap{
		entries: make(map[string]string),
	}
}

// Shorten generates a short URL for the given long URL
func (u *URLMap) Shorten(longURL string) string {
	u.mu.Lock()
	defer u.mu.Unlock()

	shortURL := fmt.Sprintf("/%d", len(u.entries))
	u.entries[shortURL] = longURL
	return shortURL
}

// Expand retrieves the original long URL for the given short URL
func (u *URLMap) Expand(shortURL string) (string, bool) {
	u.mu.Lock()
	defer u.mu.Unlock()

	longURL, ok := u.entries[shortURL]
	return longURL, ok
}

func handleShortenURL(w http.ResponseWriter, r *http.Request) {
	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortURL := urlMap.Shorten(longURL)
	fmt.Fprintf(w, "Shortened URL: %s%s\n", r.Host, shortURL)
}

func handleExpandURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := "/" + vars["shortURL"]

	longURL, ok := urlMap.Expand(shortURL)
	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

var urlMap *URLMap

func main() {
	urlMap = NewURLMap()

	r := mux.NewRouter()
	r.HandleFunc("/shorten", handleShortenURL).Methods("POST")
	r.HandleFunc("/{shortURL:[0-9]+}", handleExpandURL).Methods("GET")

	http.Handle("/", r)

	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
