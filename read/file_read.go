package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ans, err := readLines("datafile.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, ans[0])
}

func readLines(filename string) ([]string, error) {
	ans := make([]string, 10)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ans, fmt.Errorf(filename + " can't be opened")
	}
	ans = strings.Split(string(data), "\n")

	return ans, err
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
