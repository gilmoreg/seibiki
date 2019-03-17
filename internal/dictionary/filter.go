package dictionary

// filter returns a slice of meanings that are deemed relevant
func filter(pos []string, meanings []Meaning) []Meaning {
	res := make([]Meaning, 0)
	for _, m := range meanings {
		if match(pos, m) {
			res = append(res, m)
		}
	}
	return res
}

// match compares the ChaSen part of speech tags to JEDict codes
// to see if this entry matches the token in context
func match(pos []string, meaning Meaning) bool {
	// If something went wrong with pos, just return everything
	// Better to show something than nothing
	if len(pos) < 1 {
		return true
	}
	// Chasen Stage 1
	// Since this stage is so broad, if there is no match, virtually no chance this
	// is a relevant entry
	switch pos[0] {
	// Noun
	case "名詞":
		return in(meaning.PartOfSpeech, nounTypes)
	// Prefix
	case "接頭":
		return in(meaning.PartOfSpeech, prefixTypes)
	// Verb
	case "動詞":
		return in(meaning.PartOfSpeech, verbTypes)
	// Adjective
	case "形容詞":
		return in(meaning.PartOfSpeech, adjectiveTypes)
	// Adverb
	case "副詞":
		return in(meaning.PartOfSpeech, adverbTypes)
	// Adnominal  (attached to or modifying a noun)
	case "連体詞":
		return in(meaning.PartOfSpeech, adnominalTypes)
	// Conjunction
	case "接続詞":
		return in(meaning.PartOfSpeech, conjunctionTypes)
	// Particle
	case "助詞":
		return in(meaning.PartOfSpeech, particleTypes)
	// Auxilliary
	case "助動詞":
		return in(meaning.PartOfSpeech, auxilliaryTypes)
	// Interjection
	case "感動詞":
		return in(meaning.PartOfSpeech, interjectionTypes)
	}

	// If it doesn't match any ChaSen root type,
	// do not include (may change)
	return false
}

func in(pos string, posTypes []string) bool {
	for _, t := range posTypes {
		if pos == t {
			return true
		}
	}
	return false
}

var nounTypes = []string{
	"&n;",
	"noun (common) (futsuumeishi)",
	"&pn;",
	"pronoun",
	"&n-adv;",
	"adverbial noun (fukushitekimeishi)",
	"&vs;",
	"noun or participle which takes the aux. verb suru",
	"&adj-na;",
	"adjectival nouns or quasi-adjectives (keiyodoshi)",
	"&num;",
	"numeric",
	"&aux;",
	"auxiliary",
	"&n-suf;",
	"noun, used as a suffix",
	"&suf;",
	"suffix",
	"&conj;",
	"conjunction",
	"&vs;",
	"noun or participle which takes the aux. verb suru",
	"&exp;",
	"expressions (phrases, clauses, etc.)",
}

var prefixTypes = []string{
	"&pref;",
	"prefix",
	"&n-pref;",
	"noun, used as a prefix",
}

var verbTypes = []string{
	"&v",
	"verb",
	"&v1;",
	"&v1-s;",
	"&v1-s;",
	"&v2b-k;",
	"&v2h-k;",
	"&v2h-s;",
	"&v2m-s;",
	"&v2r-s;",
	"&v2t-k;",
	"&v5k-s;",
	"&v5r-i;",
	"&v5u-s;",
	"&vs-c;",
	"&vs-i;",
	"&vs-s;",
	"Godan verb - -aru special class",
	"Godan verb with `bu' ending",
	"Godan verb with `gu' ending",
	"Godan verb with `ku' ending",
	"Godan verb with `mu' ending",
	"Godan verb with `nu' ending",
	"Godan verb with `ru' ending",
	"Godan verb with `su' ending",
	"Godan verb with `tsu' ending",
	"Godan verb with `u' ending",
	"Ichidan verb",
	"Ichidan verb - zuru verb (alternative form of -jiru verbs)",
	"Kuru verb - special class",
	"Yodan verb with `hu/fu' ending (archaic)",
	"Yodan verb with `ku' ending (archaic)",
	"Yodan verb with `ru' ending (archaic)",
	"Yodan verb with `su' ending (archaic)",
	"intransitive verb",
	"irregular ru verb, plain form ends with -ri",
	"transitive verb",
	"&cop-da;",
}

var adjectiveTypes = []string{
	"&adj-f;",
	"&adj-i;",
	"&adj-ix;",
	"&adj-ku;",
	"&adj-na;",
	"&adj-nari;",
	"&adj-no;",
	"&adj-pn;",
	"&adj-shiku;",
	"&adj-t;",
}

var adverbTypes = []string{
	"&adv;",
	"&adv-to;",
	"adverb (fukushi)",
	"&n-adv;",
	"adverbial noun (fukushitekimeishi)",
}

var adnominalTypes = []string{
	"&n;",
	"noun (common) (futsuumeishi)",
	"&adj-pn;",
	"pre-noun adjectival (rentaishi)",
	"&adj-f;",
	"noun or verb acting prenominally (other than the above)",
}

var conjunctionTypes = []string{
	"&conj;",
	"conjunction",
}

var particleTypes = []string{
	"&prt;",
	"particle",
	// TODO is this correct?
	"&conj;",
	"conjunction",
}

var auxilliaryTypes = []string{
	"&aux;",
	"auxiliary",
}

var interjectionTypes = []string{
	"&int;",
	"interjection (kandoushi)",
}
