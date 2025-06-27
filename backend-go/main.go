package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type PingResponse struct {
	Message string `json:"message"`
}

type LookupResponse struct {
	Translation string `json:"translation"`
	Knowledge   string `json:"knowledge"`
}

var (
	contents []string
	mtx      sync.Mutex
)

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
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("text")
	if text == "" {
		http.Error(w, "missing text", http.StatusBadRequest)
		return
	}
	mtx.Lock()
	contents = append(contents, text)
	mtx.Unlock()
	res := ImportResponse{Status: "ok"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

type ContentResponse struct {
	Data string `json:"data"`
}

func contentHandler(w http.ResponseWriter, r *http.Request) {
	mtx.Lock()
	var text string
	if len(contents) > 0 {
		text = contents[0]
	}
	mtx.Unlock()
	res := ContentResponse{Data: text}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	contents = append(contents, "Hello World! This is sample content.")

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/lookup", lookupHandler)
	http.HandleFunc("/import", importHandler)
	http.HandleFunc("/content", contentHandler)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
