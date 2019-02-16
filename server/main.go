package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func (s *Server) handler(rw http.ResponseWriter, r *http.Request) {
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

// Server - holds deps for injection
type Server struct {
	dictionary DictionaryRepository
	router     *mux.Router
	logger     *zap.SugaredLogger
}

// Routes - add route
func (s *Server) Routes() {
	s.router.HandleFunc("/api/lookup", s.handler).Methods("POST")
	s.router.
		PathPrefix("/static/js/").
		Handler(http.StripPrefix("/static/js/", http.FileServer(http.Dir("/go/bin/wwwroot/static/js/"))))
	s.router.
		PathPrefix("/static/css/").
		Handler(http.StripPrefix("/static/css/", http.FileServer(http.Dir("/go/bin/wwwroot/static/css/"))))
	s.router.
		PathPrefix("/static/media/").
		Handler(http.StripPrefix("/static/media/", http.FileServer(http.Dir("/go/bin/wwwroot/static/media/"))))
	s.router.Handle("/", http.FileServer(http.Dir("/go/bin/wwwroot/")))
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
	s := Server{
		router:     r,
		dictionary: m,
		logger:     l,
	}
	s.Routes()
	loggedRouter := handlers.LoggingHandler(os.Stdout, s.router)
	url := fmt.Sprintf(":%s", os.Getenv("PORT"))
	l.Info(fmt.Sprintf("starting server at %s", url))
	http.ListenAndServe(url, loggedRouter)
}
