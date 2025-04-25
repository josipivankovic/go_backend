package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Oglas struct {
	ID         int    `json:"id"`
	Naslov     string `json:"naslov"`
	Cijena     string `json:"cijena"`
	Lokacija   string `json:"lokacija"`
	Opis       string `json:"opis"`
	Slika      string `json:"slika"`
	Kategorija string `json:"kategorija"`
}

var oglasi = []Oglas{}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func getOglasi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(oglasi)
}

func dodajOglas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var noviOglas Oglas
	if err := json.NewDecoder(r.Body).Decode(&noviOglas); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	noviOglas.ID = len(oglasi) + 1
	oglasi = append(oglasi, noviOglas)
	json.NewEncoder(w).Encode(noviOglas)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/oglasi", getOglasi).Methods("GET", "OPTIONS")
	r.HandleFunc("/dodaj", dodajOglas).Methods("POST", "OPTIONS")

	handler := corsMiddleware(r)

	log.Println("✅ Server pokrenut na http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
