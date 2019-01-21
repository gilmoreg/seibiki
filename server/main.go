package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func (s *server) handler(rw http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")
	if query == "" {
		http.Error(rw, "Must supply a query", 400)
		return
	}
	words := Tokenize(query)
	result := make([]Word, 0)
	for _, word := range words {
		result = append(result, word.GetEntries(s.dictionary))
	}
	j, err := json.Marshal(result)
	if err != nil {
		s.logger.Error(err)
		http.Error(rw, http.StatusText(500), 500)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(j)
}

type server struct {
	dictionary DictionaryRepository
	router     *mux.Router
	logger     *zap.SugaredLogger
}

func (s *server) routes() {
	s.router.HandleFunc("/", s.handler).Methods("POST")
}

func main() {
	l := zap.NewExample().Sugar()
	defer l.Sync()
	r := mux.NewRouter()
	c := NewRedisClient(l)
	m := &MongoDBRepository{
		cache:  c,
		logger: l,
	}
	m.Connect(os.Getenv("MONGODB_CONNECTION_STRING"))
	s := server{
		router:     r,
		dictionary: m,
		logger:     l,
	}
	s.routes()
	loggedRouter := handlers.LoggingHandler(os.Stdout, s.router)
	http.ListenAndServe(":3000", loggedRouter)
}
