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

func SendDataHandler(w http.ResponseWriter, r *http.Request) {
	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Printf("%#v\n", data)
	w.WriteHeader(http.StatusCreated)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/v1/scores", http.HandlerFunc(SendDataHandler))
	handler := cors.AllowAll().Handler(mux)
	http.ListenAndServe(":8080", handler)
}
