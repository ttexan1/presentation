package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	mux.Handle("/scores", http.HandlerFunc(SendDataHandler))
	handler := cors.AllowAll().Handler(mux)
	http.ListenAndServe(":8080", handler)
}

func addDataToRanking(data Data) (err error) {
	var ranking []Data
	bytes, err := ioutil.ReadFile("ranking.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &ranking)
	if err != nil {
		return
	}
	ranking = append(ranking, data)
	bytes, err = json.Marshal(ranking)
	if err != nil {
		return
	}
	err = ioutil.WriteFile("ranking.json", bytes, 0600)
	return
}
