package endpoints

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gilmoreg/seibiki/internal/connectors/mongodb"
	"github.com/gilmoreg/seibiki/internal/connectors/redis"
	"github.com/gilmoreg/seibiki/internal/dictionary"
	"github.com/gilmoreg/seibiki/internal/service"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestEndpoint(t *testing.T) {
	svc := createTestService()
	handler := Handler(svc)

	t.Run("Happy", func(t *testing.T) {
		body := []byte(`{ "query": "寒い中で飲むココアはうまいね" }`)
		req, _ := http.NewRequest(http.MethodPost, "/lookup", bytes.NewBuffer(body))
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)
		res := w.Result()
		assert.NotNil(t, res)
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("NonJSONBody", func(t *testing.T) {
		body := []byte(`!!!`)
		req, _ := http.NewRequest(http.MethodPost, "/lookup", bytes.NewBuffer(body))
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)
		res := w.Result()
		assert.NotNil(t, res)
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})

	t.Run("NoBody", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/lookup", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)
		res := w.Result()
		assert.NotNil(t, res)
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})
}

func createTestService() service.LookupService {
	log := zap.NewExample().Sugar()
	c := redis.New("redis://localhost:6379", log)
	m, err := mongodb.New("mongodb://reader:password@localhost:27017/jedict", log)
	if err != nil {
		panic(err)
	}
	d := dictionary.New(m, c, log)
	return service.New(log, d)
}
