package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ikawaha/kagome/tokenizer"
)

// Features is [0-5] POS, [6] base form, [7] reading, [8] pronounciation
// 一段 = ichidan, 一般 = common

type jsonToken struct {
	ID      int      `json:"id"`
	Class   string   `json:"class"` // DUMMY, KNOWN, UNKNOWN, USER
	Surface string   `json:"surface"`
	POS     []string `json:"pos"`
	Base    string   `json:"base"`
	Reading string   `json:"reading"`
	Pron    string   `json:"pron"`
}

func newToken(t tokenizer.Token) jsonToken {
	features := t.Features()
	result := jsonToken{
		ID:      t.ID,
		Class:   t.Class.String(),
		Surface: t.Surface,
		POS:     features[0:5],
		Base:    features[6],
		Reading: features[7],
		Pron:    features[8],
	}
	return result
}

func handler(rw http.ResponseWriter, r *http.Request) {
	t := tokenizer.New()
	tokens := t.Analyze("寿司が食べたい。", tokenizer.Search)
	jTokens := make([]jsonToken, 0)
	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY { // BOS and EOS
			continue
		}
		jToken := newToken(token)
		jTokens = append(jTokens, jToken)
	}
	if j, err := json.Marshal(jTokens); err != nil {
		http.Error(rw, http.StatusText(500), 500)
		return
	} else {
		rw.Write(j)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":3000", loggedRouter)
}
