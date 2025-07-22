package main

import (
	"encoding/json"
	"net/http"
)

type parameters struct {
	Body  string `json:"body"`
	Error string `json:"error"`
	Valid bool   `json:"valid"`
}

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		params.Error = "Something went wrong"
		w.WriteHeader(500)
		resp, _ := json.Marshal(params)
		w.Write(resp)
		return
	}

	if len(params.Body) > 140 {
		params.Error = "Chirp is too long"
		w.WriteHeader(400)
		resp, _ := json.Marshal(params)
		w.Write(resp)
		return
	}

	params.Valid = true
	w.WriteHeader(200)
	resp, _ := json.Marshal(params)
	w.Write(resp)

}
