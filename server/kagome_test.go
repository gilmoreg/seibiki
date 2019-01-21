package main

import (
	"testing"
)

type MockDictionaryRepository struct{}

func (m MockDictionaryRepository) Lookup(query string) []Entry {
	kanji := ""
	meaning := ""
	reading := ""
	switch query {
	case "とても":
		reading = "とても"
	case "です":
		reading = "です"
		meaning = "to be"
	case "良かった":
		return make([]Entry, 0)
	case "良い":
		reading = "よい"
		meaning = "good"
		kanji = "良い"
	}
	entry := Entry{
		Sequence:     1,
		Kanji:        []string{kanji},
		Meanings:     []string{meaning},
		Readings:     []string{reading},
		PartOfSpeech: "",
	}
	return []Entry{entry}
}

func TestTokenize(t *testing.T) {
	words := Tokenize("とても良かったです。")
	if len(words) < 4 {
		t.Fatal("Not enough words")
	}
	if len(words[1].Tokens) != 2 {
		t.Fatal("良かった should be split into two tokens")
	}
	if words[2].Surface != "です" {
		t.Fatal("Surface must be computed and be accurate")
	}
	if !words[3].IsPunctuation() {
		t.Fatal("Fourth word must be punctuation")
	}
	mockRepo := MockDictionaryRepository{}
	result := make([]Word, 0)
	for _, word := range words {
		result = append(result, word.GetEntries(mockRepo))
	}
	if result[0].Entries[0].Readings[0] != "とても" {
		t.Fatal("とても should have an accurate root entry")
	}
	if result[1].Tokens[0].Entries[0].Meanings[0] != "good" {
		t.Fatal("良かった should have accurate entries on tokens")
	}
}
