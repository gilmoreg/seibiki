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
			[]Meaning{meaning("&n;")},
			1,
		},
		{
			"verbs",
			[]string{"動詞", "自立", "*", "*"},
			[]Meaning{meaning("&v1;")},
			1,
		},
		{
			"exclude",
			[]string{"名詞", "非自立", "助動詞語幹", "*"},
			[]Meaning{meaning("&v1;")},
			0,
		},
		{
			"no match",
			[]string{"", "", "", ""},
			[]Meaning{meaning("&adj;")},
			0,
		},
		{
			"malformed pos",
			[]string{},
			[]Meaning{meaning("&adj;")},
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

func meaning(pos string) Meaning {
	return Meaning{PartOfSpeech: []string{pos}}
}
