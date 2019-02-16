package main

import (
	"github.com/ikawaha/kagome/tokenizer"
)

// Token - kagome token plus dictionary entries
type Token struct {
	ID      int      `json:"id"`
	Class   string   `json:"class"` // DUMMY, KNOWN, UNKNOWN, USER
	Surface string   `json:"surface"`
	POS     []string `json:"pos"`
	Base    string   `json:"base"`
	Reading string   `json:"reading"`
	Pron    string   `json:"pron"`
	Entries []Entry  `json:"entries"`
}

// IsPunctuation - true if token is punctuation mark
func (t Token) IsPunctuation() bool {
	if len(t.POS) < 2 {
		return false
	}
	return t.POS[0] == "記号" && t.POS[1] == "句点"
}

// Convert - create Token from kagome token
// Features is [0-5] POS, [6] base form, [7] reading, [8] pronounciation
// 一段 = ichidan, 一般 = common
func Convert(t tokenizer.Token) Token {
	features := t.Features()
	result := Token{
		ID:      t.ID,
		Class:   t.Class.String(),
		Surface: t.Surface,
		POS:     features[0:5],
		Base:    features[6],
	}
	if len(features) > 7 {
		result.Reading = features[7]
	}
	if len(features) > 8 {
		result.Pron = features[8]
	}
	return result
}

// GetEntries - fetch entries for Token from DictionaryRepository
func (t Token) GetEntries(d DictionaryRepository) Token {
	if t.IsPunctuation() {
		return t
	}
	entries := d.Lookup(t.Base)
	if len(entries) > 0 {
		t.Entries = make([]Entry, 0)
		for _, entry := range entries {
			t.Entries = append(t.Entries, entry)
		}
	}
	return t
}

// Word - set of one or more Tokens comprising a single unit
type Word struct {
	Surface string  `json:"surface"`
	Entries []Entry `json:"entries"`
	Tokens  []Token `json:"tokens"`
}

// NewWord - create Word from tokens
func NewWord(tokens []Token) Word {
	result := Word{Tokens: tokens}
	for _, token := range result.Tokens {
		result.Surface += token.Surface
	}
	return result
}

// IsPunctuation - true if first token is punctuation mark
func (w Word) IsPunctuation() bool {
	return w.Tokens[0].IsPunctuation()
}

// GetEntries - fetch entries for word/tokens from DictionaryRepository
func (w Word) GetEntries(d DictionaryRepository) Word {
	if w.IsPunctuation() {
		return w
	}

	// Try looking up word surface (won't always succeed)
	entries := d.Lookup(w.Surface)
	if len(entries) > 0 {
		w.Entries = entries
		return w
	}
	// If no result for word surface, look up each token individually
	newTokens := make([]Token, 0)
	for _, token := range w.Tokens {
		newTokens = append(newTokens, token.GetEntries(d))
	}
	w.Tokens = newTokens

	return w
}
