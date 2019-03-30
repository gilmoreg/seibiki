package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	tests := []struct {
		sentence string
		count    int
	}{
		{"寒いです。", 3},
		{"これは陽子の財布ですか。", 8},
		{"思いのほかよく描けた。自分の部屋に飾ろう", 11},
		{"とてもよかったです。。。ありがとうございます。", 9},
	}

	for _, test := range tests {
		t.Run(test.sentence, func(t *testing.T) {
			res := Tokenize(test.sentence)
			assert.Equal(t, test.count, len(res))
		})
	}
}
