package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type PingResponse struct {
	Message string `json:"message"`
}

type LookupResponse struct {
	Translation string `json:"translation"`
	Knowledge   string `json:"knowledge"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	res := PingResponse{Message: "pong"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func lookupHandler(w http.ResponseWriter, r *http.Request) {
	res := LookupResponse{Translation: "not implemented", Knowledge: "unknown"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

type ImportResponse struct {
	Status string `json:"status"`
}

func importHandler(w http.ResponseWriter, r *http.Request) {
	res := ImportResponse{Status: "not implemented"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

type ContentResponse struct {
	Data string `json:"data"`
}

func contentHandler(w http.ResponseWriter, r *http.Request) {
	res := ContentResponse{Data: "not implemented"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/lookup", lookupHandler)
	http.HandleFunc("/import", importHandler)
	http.HandleFunc("/content", contentHandler)

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
