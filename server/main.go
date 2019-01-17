package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type postBody struct {
	Query string `json:"query"`
}

func handler(rw http.ResponseWriter, r *http.Request) {
	var body postBody
	json.NewDecoder(r.Body).Decode(&body)
	if body.Query == "" {
		fmt.Println("could not decode post body")
		http.Error(rw, http.StatusText(500), 500)
		return
	}
	query := body.Query
	tokens := Tokenize(query)
	j, err := json.Marshal(tokens)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, http.StatusText(500), 500)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(j)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("POST")
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":3000", loggedRouter)
}
