package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", get)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	path := "https://" + strings.TrimPrefix(r.URL.String(), "/")

	fmt.Println("Path: " + path)

	resp, err := http.Get(path)
	if err != nil {
		http.Error(w, "Failed to GET "+path, http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	jsonBody := json.RawMessage(string(body))

	json.NewEncoder(w).Encode(jsonBody)
}
