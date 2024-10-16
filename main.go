package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var UrlDB = make(map[string]URL)

func generateShortUrl(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))

	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)

	return hash[:9]
}

func createURL(originalURL string) string {
	shortURL := generateShortUrl(originalURL)
	id := shortURL

	UrlDB[id] = URL{
		ID:           id,
		OriginalURL:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}

	return shortURL
}

func getURL(id string) (URL, error) {
	url, ok := UrlDB[id]

	if !ok {
		return URL{}, errors.New("URL not found")
	}

	return url, nil
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hellow world")
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {

	//Checking if the request is a GET request or not
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	//Parse the JSON request body
	var requestBody struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	shortURL := createURL(requestBody.URL)

	//Prepare a JSON response
	response := struct {
		Message     string `json:"message"`
		ShortendURL string `json:"short_url"`
	}{
		Message:     "The short url is generated",
		ShortendURL: shortURL,
	}

	//set response content type to `json:

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//Send the JSON response back to the client
	json.NewEncoder(w).Encode(response)

}

func RedirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/redirect/")
	url, err := getURL(id)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func main() {

	//Handlers for the server routes

	http.HandleFunc("/", handler)
	http.HandleFunc("/shorten", ShortURLHandler)
	http.HandleFunc("/redirect/", RedirectURLHandler)

	//Start a Server on port 3000
	fmt.Println("Server is listening on the port 3000...")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println("Error while starting the server", err)
	}

}
