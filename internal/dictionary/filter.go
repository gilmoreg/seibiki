package dictionary

import (
	"fmt"
	"strings"
)

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

// match compares the IPA part of speech tags to JEDict codes
// to see if this entry matches the token in context
func match(pos []string, meaning Meaning) bool {
	// If something went wrong with pos, just return everything
	// Better to show something than nothing
	if len(pos) < 1 {
		return true
	}

	partOfSpeech := strings.Join(pos, ",")

	if edictTypes, ok := ipaToEdictMapping[partOfSpeech]; ok {
		for _, edict := range meaning.PartOfSpeech {
			if in(edict, edictTypes) {
				return true
			}
		}
	} else {
		// If it doesn't match any IPA type,
		// do not include (may change)
		fmt.Println("No match for " + partOfSpeech)
		return false
	}
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

// IPA: http://chasen.naist.jp/snapshot/ipadic/ipadic/doc/ipadic-ja.pdf
// EDict: http://www.edrdg.org/jmdictdb/cgi-bin/edhelp.py?svc=jmdict&sid=#kw_pos
var ipaToEdictMapping = map[string][]string{
	// Adjective Independent
	// ex 「けたたましい」、「分別臭い」、「めでたい」
	"形容詞,自立,*,*": adjectiveEDictTypes,
	// Adjective Dependent
	// ex 「づらい」、「がたい」、「よい」
	"形容詞,非自立,*,*": adjectiveEDictTypes,
	// Adjective Suffix
	// 「ったらしい」、「っぽい」
	"形容詞,接尾,*,*": append(adjectiveEDictTypes, "&suf;"),
	// Verb Independent
	// ex 「いがみ合う」、「たてつく」、「垢抜ける」
	"動詞,自立,*,*": verbEDictTypes,
	// Verb Dependent
	// ex 「しまう」、「ちゃう」、「願う」
	"動詞,非自立,*,*": verbEDictTypes,
	// Verb Suffix
	// ex 「する」、「られる」、「させる」、「がかる」
	"動詞,接尾,*,*": append(verbEDictTypes, "&suf;"),
	// Auxililary Verb
	// ex 「らしい」、「ござる」、「っす」、「じゃん」
	"助動詞,*,*,*": append(verbEDictTypes, []string{
		"&aux-adj;",
		"&aux;",
	}...),
	// Noun Adjective Verb Stem
	// ex「あからさま」、「ミステリアス」、「決定的」、「無人」
	"名詞,形容動詞語幹,*,*": []string{
		"&adj-no;",
		"&adj-na;",
		"&adj-nari;",
		"&adj-pn;",
	},
	// Noun Suffix General
	// ex 「ぎみ」、「ゆかり」、「枚」、「不可」
	"名詞,接尾,一般,*": append(nounEDictTypes, "&suf;"),
	// Particle End
	// ex 「かしら」、「ぞ」、「っけ」、「わい」
	"助詞,終助詞,*,*": []string{
		"&prt;",
	},
	// Particle Related
	// ex 「は」、「こそ」、「も」、「や」
	"助詞,係助詞,*,*": []string{
		"&prt;",
	},
	// Particle Parallel
	// ex 「とか」、「だの」、「やら」
	// https://www.wasabi-jpn.com/japanese-grammar/parallel-markers-to-ya-and-ka/
	"助詞,並立助詞,*,*": []string{
		"&prt;",
		"&conj;",
	},
	// Particle JC( * )
	// ex 「て」、「つつ」、「および」、「ので」
	"助詞,接続助詞,*,*": []string{
		"&prt;",
	},
	// Prenominal or attributive adjective
	// ex 「この」、「いろんな」、「おっきな」、「堂々たる」
	// https://en.wiktionary.org/wiki/%E9%80%A3%E4%BD%93%E8%A9%9E
	"連体詞,*,*,*": []string{
		"&adj-pn;",
	},
	// Noun General
	// ex 「大根」、「シエスタ」、「加速度」、「ありさま」
	"名詞,一般,*,*": nounEDictTypes,
	// Noun Number
	// ex 「ゼロ」、「億」
	"名詞,数,*,*": []string{
		"&num;",
	},
	// Noun Proper Noun Name First Name
	// ex「Ｂ作」、「アントニオ」、「右京太夫」
	"名詞,固有名詞,人名,名": []string{"&n-pr;"},
	// Noun Proper Noun Area General
	// ex 「北海道」、「やながわ工業団地」、「ラムサール」
	"名詞,固有名詞,地域,一般": []string{"&n-pr;"},
	// Noun Suffix Counter
	// ex 「オクターブ」、「％」、「ヶ国」
	"名詞,接尾,助数詞,*": []string{
		"&n-suf;",
		"&ctr;",
		"&suf;",
	},
	// Noun Proper Noun Name Last Name
	// ex 「山田」、「ビスコンティ」
	"名詞,固有名詞,人名,姓": []string{"&n-pr;"},
	// Adverb General
	// ex 「たいそう」、「人一倍」、「いけしゃあしゃあ」
	"副詞,一般,*,*": adverbEDictTypes,
	// Prefix Connection Noun
	// ex 「もと」、「アンチ」、「最」、「総」
	"接頭詞,名詞接続,*,*": []string{
		"&pref;",
	},
	// Adverb Connection Particle
	// ex 「あまり」、「いつも」、「ぱさぱさ」
	"副詞,助詞類接続,*,*": adverbEDictTypes,
	// Noun Adverb Possible
	// ex 「１０月」、「せんだって」、「当分」
	"名詞,副詞可能,*,*": []string{
		"&n-adv;",
		"&adv;",
	},
	// Noun Connection Part Verb
	// ex 「苦労」、「終了」、「アピール」、「くしゃみ」
	"名詞,サ変接続,*,*": []string{
		"&vs;",
	},
	//
	//
	"名詞,接尾,副詞可能,*": []string{
		"&n-adv;",
		"&n-pref;",
		"&n-suf;",
		"&n;",
		"&num;",
		"&pref;",
		"&suf;",
	},
	"名詞,固有名詞,一般,*": []string{
		"&adj-f;",
		"&adj-i;",
		"&adj-na;",
		"&adj-no;",
		"&adj-t;",
		"&adv-to;",
		"&adv;",
		"&ctr;",
		"&exp;",
		"&n-adv;",
		"&n-pr;",
		"&n-suf;",
		"&n;",
		"&pref;",
		"&prt;",
		"&suf;",
		"&v2h-k;",
	},
	"名詞,固有名詞,組織,*": []string{
		"&adj-f;",
		"&adj-na;",
		"&adj-no;",
		"&adv-to;",
		"&adv;",
		"&exp;",
		"&int;",
		"&n-adv;",
		"&n-pr;",
		"&n-suf;",
		"&n;",
		"&pref;",
		"&suf;",
		"&vs;",
	},
	"名詞,固有名詞,地域,国": []string{
		"&adj-na;",
		"&n-adv;",
		"&n-pr;",
		"&n-suf;",
		"&n;",
		"&suf;",
	},
	"接続詞,*,*,*": []string{
		"&adv;",
		"&aux;",
		"&conj;",
		"&exp;",
		"&int;",
		"&n-adv;",
		"&n;",
		"&prt;",
	},
	"感動詞,*,*,*": []string{
		"&adv-to;",
		"&adv;",
		"&aux-v;",
		"&aux;",
		"&conj;",
		"&exp;",
		"&int;",
		"&n;",
		"&pn;",
		"&prt;",
	},
	"助詞,副助詞,*,*": []string{
		"&aux;",
		"&conj;",
		"&exp;",
		"&int;",
		"&n-adv;",
		"&prt;",
		"&suf;",
	},
	"助詞,格助詞,一般,*": []string{
		"&aux;",
		"&conj;",
		"&int;",
		"&n;",
		"&prt;",
	},
	"助詞,格助詞,連語,*": []string{
		"&adv;",
		"&aux;",
		"&conj;",
		"&exp;",
		"&prt;",
	},
	"名詞,非自立,副詞可能,*": []string{
		"&n-adv;",
		"&n-suf;",
		"&n;",
		"&pref;",
		"&suf;",
		"&vs-c;",
	},
	"名詞,代名詞,一般,*": []string{
		"&adv;",
		"&conj;",
		"&exp;",
		"&int;",
		"&n-adv;",
		"&n;",
		"&pn;",
		"&prt;",
		"&suf;",
	},
	"名詞,固有名詞,人名,一般": []string{
		"&adj-na;",
		"&exp;",
		"&n-adv;",
		"&n-pref;",
		"&n;",
		"&vs;",
	},
	"名詞,ナイ形容詞語幹,*,*": []string{
		"&n-adv;",
		"&n;",
		"&suf;",
	},
	"名詞,非自立,一般,*": []string{
		"&exp;",
		"&int;",
		"&n-adv;",
		"&n-pref;",
		"&n-suf;",
		"&n;",
		"&pref;",
		"&prt;",
		"&suf;",
	},
	"接頭詞,数接続,*,*": []string{
		"&adj-no;",
		"&adv;",
		"&ctr;",
		"&n-adv;",
		"&n-pref;",
		"&n-t;",
		"&n;",
		"&pref;",
		"&suf;",
	},
	"フィラー,*,*,*": []string{
		"&adv;",
		"&exp;",
		"&int;",
		"&prt;",
	},
	"名詞,動詞非自立的,*,*": []string{
		"&exp;",
		"&int;",
		"&n;",
	},
	"名詞,非自立,助動詞語幹,*": []string{
		"&aux-v;",
		"&int;",
		"&n-suf;",
		"&n;",
		"&prt;",
		"&suf;",
	},
	"その他,間投,*,*": []string{
		"&int;",
		"&prt;",
	},
	"名詞,接尾,特殊,*": []string{
		"&int;",
		"&n;",
		"&prt;",
		"&suf;",
	},
	"接頭詞,形容詞接続,*,*": []string{
		"&int;",
		"&n;",
		"&pn;",
		"&pref;",
	},
	"助詞,特殊,*,*": []string{
		"&aux-v;",
		"&exp;",
		"&int;",
		"&prt;",
		"&suf;",
	},
	"助詞,格助詞,引用,*": []string{
		"&conj;",
		"&n;",
		"&prt;",
	},
	"助詞,副詞化,*,*": []string{
		"&conj;",
		"&n;",
		"&prt;",
	},
	"名詞,接続詞的,*,*": []string{
		"&conj;",
		"&n;",
	},
	"名詞,接尾,地域,*": []string{
		"&n-suf;",
		"&n;",
		"&suf;",
	},
	"名詞,接尾,サ変接続,*": []string{
		"&ctr;",
		"&n-suf;",
		"&n;",
		"&suf;",
	},
	"鐃緒申鐃銃誌申,鐃緒申立,*,*": []string{
		"&n-suf;",
		"&n;",
		"&pn;",
		"&pref;",
		"&suf;",
	},
	"名詞,接尾,人名,*": []string{
		"&n-suf;",
		"&n;",
		"&pn;",
		"&suf;",
	},
	"名詞,接尾,形容動詞語幹,*": []string{
		"&adj-na;",
		"&n-pref;",
		"&n-suf;",
		"&n;",
		"&pref;",
		"&suf;",
	},
	"名詞,特殊,助動詞語幹,*": []string{
		"&adj-na;",
		"&adv;",
		"&suf;",
	},
	"名詞,接尾,助動詞語幹,*": []string{
		"&adj-na;",
		"&adv;",
		"&suf;",
	},
	"助詞,副助詞／並立助詞／終助詞,*,*": []string{
		"&adv;",
		"&pref;",
		"&prt;",
		"&suf;",
	},
	"名詞,非自立,形容動詞語幹,*": []string{
		"&adv-to;",
		"&adv;",
		"&suf;",
	},
	"記号,アルファベット,*,*": []string{
		"&n;",
		"&pref;",
	},
	"接頭詞,動詞接続,*,*": []string{
		"&n;",
		"&pref;",
		"&suf;",
	},
	"助詞,連体化,*,*": []string{
		"&prt;",
	},
}

var adjectiveEDictTypes = []string{
	"&adj-f;",
	"&adj-i;",
	"&adj-na;",
	"&adj-ix;",
	"&adj-kari;",
	"&adj-ku;",
	"&adj-na;",
	"&adj-nari;",
}

var verbEDictTypes = []string{
	"&aux-v;",
	"&v1-s;",
	"&v1;",
	"&v2a-s;",
	"&v2g-s;",
	"&v2h-k;",
	"&v2m-s;",
	"&v2n-s;",
	"&v2r-k;",
	"&v2r-s;",
	"&v2s-s;",
	"&v2t-k;",
	"&v2z-s;",
	"&v4k;",
	"&v4m;",
	"&v4r;",
	"&v4s;",
	"&v4t;",
	"&v5aru;",
	"&v5b;",
	"&v5g;",
	"&v5k-s;",
	"&v5k;",
	"&v5m;",
	"&v5n;",
	"&v5r-i;",
	"&v5r;",
	"&v5s;",
	"&v5t;",
	"&v5u-s;",
	"&v5u;",
	"&vi;",
	"&vk;",
	"&vr;",
	"&vs-c;",
	"&vs-i;",
	"&vs-s;",
	"&vz;",
}

var nounEDictTypes = []string{
	"&n-adv;",
	"&n-pr;",
	"&n-pref;",
	"&n-suf;",
	"&n-t;",
	"&n;",
	"&vs;",
}

var adverbEDictTypes = []string{
	"&n-adv;",
	"&adv;",
	"&adv-to;",
}
