package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type MockDictionaryRepository struct {
	entries map[string][]Entry
}

func NewMockDictRepo(filename string) MockDictionaryRepository {
	jsonFile, _ := os.Open(filename)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result = MockDictionaryRepository{}
	json.Unmarshal([]byte(byteValue), &result.entries)
	return result
}

func (m MockDictionaryRepository) Lookup(query string) []Entry {
	return m.entries[query]
}

func TestIntegration(t *testing.T) {
	tt := []struct {
		query    string
		expected string
	}{
		{query: "とても良かったです。", expected: "../test/testdata/example1.json"},
		{query: "すもももももももものうち", expected: "../test/testdata/example2.json"},
		{query: "デジカメを買った", expected: "../test/testdata/example3.json"},
	}

	mockRepo := NewMockDictRepo("../test/testdata/entries.json")
	mockServer := Server{
		dictionary: mockRepo,
		router:     mux.NewRouter(),
		logger:     zap.NewExample().Sugar(),
	}
	mockServer.Routes()

	for _, tc := range tt {
		t.Run(tc.query, func(t *testing.T) {
			server := httptest.NewServer(mockServer.router)
			defer server.Close()
			data := url.Values{}
			data.Set("query", tc.query)
			formData := url.Values{
				"query": {tc.query},
			}
			resp, _ := http.PostForm(server.URL+"/api/lookup", formData)
			var words []Word
			json.NewDecoder(resp.Body).Decode(&words)

			jsonFile, _ := os.Open(tc.expected)
			defer jsonFile.Close()
			byteValue, _ := ioutil.ReadAll(jsonFile)
			var result []Word
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
