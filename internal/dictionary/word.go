package dictionary

import (
	"github.com/ikawaha/kagome.ipadic/tokenizer"
)

// IsPunctuation - true if token is punctuation mark
func (t Token) IsPunctuation() bool {
	return t.POS[0] == "記号"
}

// Convert - create Token from kagome token
// Features is [0-5] POS (0-4 IPA codes, unsure what 5 is)
// [6] base form, [7] reading, [8] pronounciation
func Convert(t tokenizer.Token) Token {
	features := t.Features()
	result := Token{
		ID:      t.ID,
		Class:   t.Class.String(),
		Surface: t.Surface,
		POS:     features[0:4],
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
func (t Token) GetEntries(r Repository) Token {
	if t.IsPunctuation() {
		return t
	}
	entries, err := r.Lookup(t.Base)
	if err != nil {
		return Token{}
	}
	if len(entries) > 0 {
		t.Entries = make([]Entry, 0)
		for _, entry := range entries {
			entry.Meanings = filter(t.POS, entry.Meanings)
			if len(entry.Meanings) > 0 {
				t.Entries = append(t.Entries, entry)
			}
		}
	}
	return t
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

// GetEntries - fetch entries for tokens from DictionaryRepository
func (w Word) GetEntries(r Repository) Word {
	if w.IsPunctuation() {
		return w
	}

	// If no result for word surface, look up each token individually
	newTokens := make([]Token, 0)
	for _, token := range w.Tokens {
		newTokens = append(newTokens, token.GetEntries(r))
	}
	w.Tokens = newTokens

	return w
}

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

// Entry - dictionary entry
type Entry struct {
	Sequence int       `json:"sequence"`
	Kanji    []string  `json:"kanji"`
	Readings []string  `json:"readings"`
	Meanings []Meaning `json:"meanings"`
}

// Meaning - an English meaning with its part of speech
type Meaning struct {
	Gloss        string   `json:"gloss"`
	PartOfSpeech []string `json:"partofspeech"`
	Misc         []string `json:"misc"`
}

// Word - set of one or more Tokens comprising a single unit
type Word struct {
	Surface string  `json:"surface"`
	Tokens  []Token `json:"tokens"`
}
