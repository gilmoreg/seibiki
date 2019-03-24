package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gilmoreg/seibiki/internal/connectors/mongodb"
	"github.com/gilmoreg/seibiki/internal/connectors/redis"
	"github.com/gilmoreg/seibiki/internal/dictionary"
	"github.com/gilmoreg/seibiki/internal/endpoints"
	"github.com/gilmoreg/seibiki/internal/service"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Server - holds deps for injection
type Server struct {
	svc    service.LookupService
	router *mux.Router
	logger *zap.SugaredLogger
}

// Routes - add routes
func (s *Server) Routes() {
	s.router.Path("/api/lookup").Methods("POST").Handler(endpoints.Handler(s.svc, s.logger))
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
	c := redis.New(os.Getenv("REDIS_URL"), l)
	m := mongodb.New(os.Getenv("MONGODB_CONNECTION_STRING"), l)
	d := dictionary.New(m, c, l)
	svc := service.New(l, d)
	s := Server{
		router: r,
		svc:    svc,
		logger: l,
	}
	s.Routes()
	url := fmt.Sprintf(":%s", os.Getenv("PORT"))
	l.Info(fmt.Sprintf("starting server at %s", url))
	http.ListenAndServe(url, r)
}
