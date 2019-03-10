package dictionary

import (
	"github.com/ikawaha/kagome/tokenizer"
)

func segment(tokens []tokenizer.Token) []Word {
	words := make([]Word, 0)
	currentWord := make([]Token, 0)
	for _, t := range tokens {
		if t.Class == tokenizer.DUMMY { // BOS and EOS
			continue
		}
		token := Convert(t)

		// If we are at a punctuation mark or
		// a word whose base is equal to its surface
		// (meaning it is not a conjugated helper)
		// add current word to words and start a new word
		if token.IsPunctuation() {
			// Finish adding word in progress, if any
			if len(currentWord) > 0 {
				words = append(words, NewWord(currentWord))
			}
			// Add punctuation as its own word
			words = append(words, NewWord([]Token{token}))
			// Start a new word
			currentWord = make([]Token, 0)
		} else if token.Surface == token.Base {
			// Add ending to current word
			currentWord = append(currentWord, token)
			// Add current word
			words = append(words, NewWord(currentWord))
			// Start a new word
			currentWord = make([]Token, 0)
		} else {
			// We aren't finished - add token to current word
			currentWord = append(currentWord, token)
		}
	}
	// Finish up word in progress, if any
	if len(currentWord) > 0 {
		words = append(words, NewWord(currentWord))
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
