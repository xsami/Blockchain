package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func handleGetBlockchain(w http.ResponseWriter, r *http.Request, Blockchain []Block) {
	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}
