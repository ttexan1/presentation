package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

type Data struct {
	Name  string `json:"name"`
	Point int    `json:"point"`
}

var recentData Data

func SendDataHandler(w http.ResponseWriter, r *http.Request) {
	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	recentData = data
	fmt.Printf("%#v\n", data)
	w.WriteHeader(http.StatusCreated)
}

func saveDataHandler(w http.ResponseWriter, r *http.Request) {
	ping := recentData
	res, err := json.Marshal(ping)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/v1/scores", http.HandlerFunc(SendDataHandler))
	mux.Handle("/api/v1/score", http.HandlerFunc(saveDataHandler))
	handler := cors.AllowAll().Handler(mux)
	http.ListenAndServe(":8080", handler)
}
