package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name     string
		pos      []string
		meanings []Meaning
		expected int
	}{
		{
			"nouns",
			[]string{"名詞", "一般", "*", "*"},
			meanings("&n;"),
			1,
		},
		{
			"verbs",
			[]string{"動詞", "自立", "*", "*"},
			meanings("&v1;"),
			1,
		},
		{
			"exclude",
			[]string{"名詞", "非自立", "助動詞語幹", "*"},
			meanings("&v1;"),
			0,
		},
		{
			"no match",
			[]string{"", "", "", ""},
			meanings("&adj;"),
			0,
		},
		{
			"malformed pos",
			[]string{},
			meanings("&adj;"),
			1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := filter(test.pos, test.meanings)
			assert.Equal(t, test.expected, len(res))
		})
	}
}

func meanings(pos string) []Meaning {
	return []Meaning{Meaning{PartOfSpeech: []string{pos}}}
}
