package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gilmoreg/seibiki/internal/connectors/mongodb"
	"github.com/gilmoreg/seibiki/internal/dictionary"
	"github.com/gilmoreg/seibiki/internal/service"
	"github.com/gomodule/redigo/redis"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestLookupEndpoint(t *testing.T) {
	tt := []struct {
		query    string
		expected string
	}{
		{query: "とても良かったです。", expected: "../../testdata/testdata/example1.json"},
		{query: "すもももももももものうち", expected: "../../testdata/testdata/example2.json"},
		{query: "デジカメを買った", expected: "../../testdata/testdata/example3.json"},
	}

	ep := createTestEndpoint()
	server := httptest.NewServer(ep)
	defer server.Close()

	for _, tc := range tt {
		t.Run(tc.query, func(t *testing.T) {
			body := []byte(fmt.Sprintf(`{ "query": "%s" }`, tc.query))
			resp, err := http.Post(server.URL+"/api/lookup", "application/json", bytes.NewBuffer(body))
			assert.Nil(t, err)
			var words []dictionary.Word
			json.NewDecoder(resp.Body).Decode(&words)

			jsonFile, err := os.Open(tc.expected)
			assert.Nil(t, err)
			defer jsonFile.Close()
			byteValue, err := ioutil.ReadAll(jsonFile)
			assert.Nil(t, err)
			var result []dictionary.Word
			json.Unmarshal([]byte(byteValue), &result)
			if !cmp.Equal(words, result) {
				fmt.Println(words)
				fmt.Println("")
				fmt.Println(result)
				t.Fatal("Output did not match expected results")
			}
		})
	}
}

func createTestEndpoint() http.Handler {
	l := zap.NewExample().Sugar()
	c := &mockCacheClient{}
	m := newMockDBClient()
	d := dictionary.New(m, c, l)
	svc := service.New(l, d)
	return Handler(svc, l)
}

type mockCacheClient struct{}

func (m *mockCacheClient) Get(key string) ([]byte, error) {
	return nil, redis.ErrNil
}

func (m *mockCacheClient) Set(key string, value []byte) error {
	return nil
}

type mockDBClient struct {
	entries map[string][]dictionary.Entry
}

func newMockDBClient() mongodb.Client {
	jsonFile, _ := os.Open("../../testdata/testdata/entries.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result = mockDBClient{}
	json.Unmarshal([]byte(byteValue), &result.entries)
	return &result
}

func (m *mockDBClient) Get(query string) ([]byte, error) {
	return json.Marshal(m.entries[query])
}
