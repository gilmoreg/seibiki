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

func (s *server) handler(rw http.ResponseWriter, r *http.Request) {
	var body postBody
	json.NewDecoder(r.Body).Decode(&body)
	if body.Query == "" {
		fmt.Println("could not decode post body")
		http.Error(rw, http.StatusText(500), 500)
		return
	}
	query := body.Query
	words := Tokenize(query)
	result := make([]Word, 0)
	for _, word := range words {
		result = append(result, word.GetEntries(s.dictionary))
	}
	j, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, http.StatusText(500), 500)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(j)
}

type server struct {
	dictionary DictionaryRepository
	router     *mux.Router
}

func (s *server) routes() {
	s.router.HandleFunc("/", s.handler).Methods("POST")
}

func main() {
	r := mux.NewRouter()
	c := NewRedisClient()
	m := MongoDBRepository{}.New(os.Getenv("MONGODB_CONNECTION_STRING"), c)
	s := server{
		router:     r,
		dictionary: m,
	}
	s.routes()
	loggedRouter := handlers.LoggingHandler(os.Stdout, s.router)
	http.ListenAndServe(":3000", loggedRouter)
}
