package main

import (
	"github.com/ikawaha/kagome/tokenizer"
)

// Features is [0-5] POS, [6] base form, [7] reading, [8] pronounciation
// 一段 = ichidan, 一般 = common

// JSONToken - json friendly representation of tokenizer.Token
type JSONToken struct {
	ID      int      `json:"id"`
	Class   string   `json:"class"` // DUMMY, KNOWN, UNKNOWN, USER
	Surface string   `json:"surface"`
	POS     []string `json:"pos"`
	Base    string   `json:"base"`
	Reading string   `json:"reading"`
	Pron    string   `json:"pron"`
}

// Word - one or more JSONTokens
type Word []JSONToken

func newToken(t tokenizer.Token) JSONToken {
	features := t.Features()
	result := JSONToken{
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

func isPunctuation(pos []string) bool {
	if len(pos) < 2 {
		return false
	}
	return pos[0] == "記号" && pos[1] == "句点"
}

func segment(tokens []tokenizer.Token) []Word {
	words := make([]Word, 0)
	word := make([]JSONToken, 0)
	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY { // BOS and EOS
			continue
		}
		jToken := newToken(token)

		// If we are at a punctuation mark or
		// a word whose base is equal to its surface
		// (meaning it is not a conjugated helper)
		// add current word to words and start a new word
		if isPunctuation(jToken.POS) {
			// Finish adding word in progress, if any
			if len(word) > 0 {
				words = append(words, word)
			}
			// Add punctuation as its own word
			words = append(words, []JSONToken{jToken})
			// Start a new word
			word = make([]JSONToken, 0)
		} else if jToken.Surface == jToken.Base {
			// Add ending to current word
			word = append(word, jToken)
			// Add current word
			words = append(words, word)
			// Start a new word
			word = make([]JSONToken, 0)
		} else {
			// We aren't finished - add token to current word
			word = append(word, jToken)
		}
	}
	// Finish up word in progress, if any
	if len(word) > 0 {
		words = append(words, word)
	}
	return words
}

// Tokenize - use Kagome to tokenize input string, collect into words
func Tokenize(query string) []Word {
	t := tokenizer.New()
	tokens := t.Analyze(query, tokenizer.Search)
	words := segment(tokens)
	return words
}
