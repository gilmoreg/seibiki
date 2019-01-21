package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestTokenize(t *testing.T) {
	tt := []struct {
		query    string
		expected string
	}{
		{query: "とても良かったです。", expected: "../test/testdata/example1.json"},
		{query: "すもももももももものうち", expected: "../test/testdata/example2.json"},
		{query: "デジカメを買った", expected: "../test/testdata/example3.json"},
	}

	mockRepo := NewMockDictRepo("../test/testdata/entries.json")

	for _, tc := range tt {
		t.Run(tc.query, func(t *testing.T) {
			words := Tokenize(tc.query)
			wordsEntries := make([]Word, 0)
			for _, word := range words {
				wordsEntries = append(wordsEntries, word.GetEntries(mockRepo))
			}
			jsonFile, _ := os.Open(tc.expected)
			defer jsonFile.Close()
			byteValue, _ := ioutil.ReadAll(jsonFile)

			var result []Word
			json.Unmarshal([]byte(byteValue), &result)
			if !cmp.Equal(wordsEntries, result) {
				fmt.Println(wordsEntries)
				fmt.Println("")
				fmt.Println(result)
				t.Fatal("Output did not match expected results")
			}
		})
	}
}
