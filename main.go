package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

/*
	for example: s200ta1--> {
	                            ID:"s200ta1",
								OriginalURL:"https://github.com/sharanyashetty013-droid",
								ShortURL:"http://localhost:8080/s200ta1",
								CreationDate:time.Now(
								}
*/
var urlDB = make(map[string]URL)

//now making a function to short thee original url toshort url
/* here we use string at end cuz it will have to
return a string value back and also we use hashing method to sort
*/
func generateShortURL(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	fmt.Println("Hasher:", hasher)
	data := hasher.Sum(nil)
	fmt.Println("Data hasher:", data)
	hash := hex.EncodeToString(data)
	fmt.Println("encode to string", hash)
	fmt.Println("final string:", hash[:8])
	return hash[:8]
}

// till gere we generated short url from og url
// now after short we have to store the value in DB
func createURL(OriginalURL string) string {
	shortURL := generateShortURL(OriginalURL)
	id := shortURL // we can use short url as id as well as DB need id
	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  OriginalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}
func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hey techy buddies server is working")
}
func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL := createURL(data.URL)
	fmt.Fprintf(w, "Short URL: %s", shortURL)
}

// redirect handler to redirect the short url to original url
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:] // Remove the leading slash
	url, err := getURL(id)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusMovedPermanently)
}

func main() {
	////fmt.Println("Starting url shortner")
	//OriginalURL := "https://github.com/sharanyashetty013-droid"
	//generateShortURL(OriginalURL)
	//register the handler function to handle the request to the root URL
	http.HandleFunc("/", RedirectHandler)
	http.HandleFunc("/shorten", ShortURLHandler)
	//start the HTTP SERVER on port 3000
	fmt.Println("Starting server on port 3000 . . . .")
	//here we are using http package to start the server
	//  and listen on port 3000
	//we use :3000 port as a string and nil as handler
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
