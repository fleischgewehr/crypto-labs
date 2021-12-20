package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	Text string `json:"text"`
}

func Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{Text: "This is fine."})
}

func main() {
	http.HandleFunc("/", Handler)
	fmt.Println("server started at localhost:443...")
	err := http.ListenAndServeTLS("127.0.0.1:443", "resources/server.crt", "resources/server.key", nil)
	if err != nil {
		panic(err)
	}
}
