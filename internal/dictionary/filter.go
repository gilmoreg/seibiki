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
// Numbers correpsond to IPADIC manual
var ipaToEdictMapping = map[string][]string{
	/*
	  5.1 名詞 - Nouns
	*/

	// 5.1.1 Noun General
	// ex 「大根」、「シエスタ」、「加速度」、「ありさま」
	"名詞,一般,*,*": nounEDictTypes,

	// Proper Nouns

	// 5.1.2 Proper Noun General
	// ex 「北穂高岳」、「電通銀座ビル」、「Ｇ１」
	"名詞,固有名詞,一般,*": []string{"&n-pr;"},

	// 5.1.3 Proper Noun Name General
	// ex 「グッチ裕三」、「紫式部」
	"名詞,固有名詞,人名,一般": []string{"&n-pr;"},

	// 5.1.4 Proper Noun Name Last Name
	// ex 「山田」、「ビスコンティ」
	"名詞,固有名詞,人名,姓": []string{"&n-pr;"},

	// 5.1.5 Proper Noun Name First Name
	// ex「Ｂ作」、「アントニオ」、「右京太夫」
	"名詞,固有名詞,人名,名": []string{"&n-pr;"},

	// 5.1.6 Proper Noun Name Organization
	// ex 「いすゞ自動車」、「ニチレイ」、「統一アイルランド党」
	"名詞,固有名詞,組織,*": []string{"&n-pr;"},

	// 5.1.7 Proper Noun Area General
	// ex 「北海道」、「やながわ工業団地」、「ラムサール」
	"名詞,固有名詞,地域,一般": []string{"&n-pr;"},

	// 5.1.8 Proper Noun Country
	// ex 「露西亜」、「バングラデシュ」
	"名詞,固有名詞,地域,国": []string{"&n-pr;"},

	// Pronouns

	// 5.1.9 Pronoun General
	// ex 「そこ」、「俺」、「こんちくしょう」
	"名詞,代名詞,一般,*": []string{"&pn;"},

	// 5.1.10 Pronoun Contraction
	// ex 「わたしゃ」、「そりゃあ」
	"名詞,代名詞,縮約,*": []string{"&exp;"},

	// 5.1.11 Adverb Possible
	// ex 「１０月」、「せんだって」、「当分」
	"名詞,副詞可能,*,*": []string{
		"&n-adv;",
		"&adv;",
	},

	// 5.1.12 Verbal Connection (i.e. -suru noun/participle)
	// ex 「苦労」、「終了」、「アピール」、「くしゃみ」
	"名詞,サ変接続,*,*": []string{"&vs;"},

	// 5.1.13 Adjective verb stem
	// ex「あからさま」、「ミステリアス」、「決定的」、「無人」
	// TODO unclear on this one
	// Desc A so-called adjective verb stem that appears before "na".
	"名詞,形容動詞語幹,*,*": []string{
		"&adj-no;",
		"&adj-na;",
		"&adj-nari;",
		"&adj-pn;",
	},

	// 5.1.14 Nai adjective stem
	// ex 「申し訳」、「とんでも」、「おとなげ」
	// TODO unclear on this one
	// Desc An adjective word that appears immediately before the auxiliary verb "nai"
	"名詞,ナイ形容詞語幹,*,*": []string{
		"&n;",
		"&n-pref;",
	},

	// 5.1.15 Number
	// ex 「ゼロ」、「億」
	"名詞,数,*,*": []string{"&num;"},

	// Dependent Nouns

	// 5.1.16 Noun Dependent General
	// ex 「こと」、「きらい」、「くせ」、「もの」
	"名詞,非自立,一般,*": nounEDictTypes,

	// 5.1.17 Noun Dependent Possible
	// ex 「限り」、「さなか」、「うち」
	"名詞,非自立,副詞可能,*": nounEDictTypes,

	// 5.1.18 Auxilliary Verb Stem
	// ex 「よ」、「よう」のみ
	"名詞,非自立,助動詞語幹,*": []string{"&n;"},

	// 5.1.19 Adjective Verb Stem
	// ex 「ふう」、「みたい」のみ
	// TODO only two examples in IPA are listed above, and both are suffixes
	// though the translation of the type suggests a prefix/prenominal?
	"名詞,非自立,形容動詞語幹,*": []string{
		"&n-suf;",
		"&suf;",
	},

	// 5.1.20 Special Auxilliary Verb Stem
	// ex 「そ」、「そう」のみ (only two examples in the IPA)
	// TODO all examples in JEDict are nouns
	// possibility for そう: https://jisho.org/word/%E3%81%9D%E3%81%86
	// no possibilities for そ alone
	// Desc The stem part of "そうだ" that is connected to the basic form and is an auxiliary verb in the school grammar.
	"名詞,特殊,助動詞語幹,*": []string{
		"&adj-na;",
		"&suf;",
	},

	// 5.1.21 Suffix General
	// ex 「ぎみ」、「ゆかり」、「枚」、「不可」
	"名詞,接尾,一般,*": []string{
		"&n-suf;",
		"&suf;",
	},

	// 5.1.22 Suffix Name
	// ex 「君」、「はん」
	"名詞,接尾,人名,*": []string{
		"&n-suf;",
		"&n;",
		"&pn;",
		"&suf;",
	},

	// 5.1.23 Suffix Region
	// ex 「州」、「市内」、「港」
	"名詞,接尾,地域,*": []string{
		"&n-suf;",
		"&n;",
		"&pn;",
		"&suf;",
	},

	// 5.1.24 Suffix Verbal Connection
	// ex 「化」、「入り」
	"名詞,接尾,サ変接続,*": []string{
		"&n-suf;",
		"&n;",
		"&suf;",
	},

	// 5.1.25 Suffix Auxilliary Verb Stem
	// ex 「そ」、「そう」のみ
	// TODO see 5.1.20
	"名詞,接尾,助動詞語幹,*": []string{
		"&adj-na;",
		"&suf;",
	},

	// 5.1.26 Suffix Adjective Verb Stem
	// ex 「がち」、「的」、「同然」
	"名詞,接尾,形容動詞語幹,*": []string{
		"&adj-na;",
		"&n-suf;",
		"&suf;",
	},

	// 5.1.27 Suffix Adverb Possible
	// ex 「いっぱい」、「前後」、「次第」
	"名詞,接尾,副詞可能,*": []string{
		"&adv;",
		"&n-adv;",
		"&n-suf;",
		"&suf;",
	},

	// 5.1.28 Suffix Counter
	// ex 「オクターブ」、「％」、「ヶ国」
	"名詞,接尾,助数詞,*": []string{
		"&n-suf;",
		"&ctr;",
		"&suf;",
	},

	// 5.1.29 Suffix Special
	// ex 「ぶり」、「み」、「方」
	"名詞,接尾,特殊,*": []string{
		"&n-suf;",
		"&suf;",
	},

	// 5.1.30 Suffix Conjunctional
	// ex 「VS」、「対」、「兼」のみ (all examples in IPA)
	"名詞,接続詞的,*,*": []string{"&conj;"},

	// 5.1.31 Noun Verb Depdendent
	// ex「ごらん」、「ちょうだい」のみ
	// Most examples match the expression or interjection "please do this"
	// Desc It is connected to "te" of [助詞-接続助詞] and is semantically verbal.
	"名詞,動詞非自立的,*,*": []string{
		"&exp;",
		"&int;",
	},

	/*
	  5.2 Prefixes
	*/

	// 5.2.1 Noun Connection
	// ex 「もと」、「アンチ」、「最」、「総」
	"接頭詞,名詞接続,*,*": []string{
		"&n-pref;",
		"&pref;",
	},

	// 5.2.2 Counter connection
	// ex 「No.」、「計」、「毎分」
	// Many odd classifications in JEDict, might not get them all
	"接頭詞,数接続,*,*": []string{
		"&n-pref;",
		"&pref;",
		"&n-t;",
		"&ctr;",
		"&adj-no;",
		"&n;",
	},

	// 5.2.3 Verb Connection
	// ex 「ぶっ」、「引き」
	// Under prefix category, but 引き is only ever a suffix
	// Desc The verb's imperative form or [verb syntactic form] + a prefix that precedes なる／なさる／くださる
	"接頭詞,動詞接続,*,*": []string{
		"&n;",
		"&pref;",
		"&suf;",
	},

	// 5.2.4 Adjective Connection
	// ex 「お」、「まっ」、「クソ」
	// Desc Prefix prefixed to adjectives.
	"接頭詞,形容詞接続,*,*": []string{"&pref;"},

	/*
	   5.3 Verbs
	*/

	// 5.3.1,3-6,9,13-14,16-19,21,24,26,28,30-34 Verb Independent
	// ex 「いがみ合う」、「たてつく」、「垢抜ける」
	// 「くる」「来る」「やってくる」「やって来る」
	"動詞,自立,*,*": verbEDictTypes,

	// 5.3.2,7,10-12,15,20,22,25,27,29 Verb Depdendent
	// ex 「（て）くる」「（て）来る」
	// 「しまう」、「ちゃう」、「願う」
	"動詞,非自立,*,*": verbEDictTypes,

	// 5.3.8,23 Verb Suffix
	// ex 「する」、「られる」、「させる」、「がかる」
	"動詞,接尾,*,*": append(verbEDictTypes, "&suf;"),

	/*
	   5.4 Adjectives
	*/

	// Adjective Independent
	// ex 「けたたましい」、「分別臭い」、「めでたい」
	"形容詞,自立,*,*": adjectiveEDictTypes,
	// Adjective Dependent
	// ex 「づらい」、「がたい」、「よい」
	"形容詞,非自立,*,*": adjectiveEDictTypes,
	// Adjective Suffix
	// 「ったらしい」、「っぽい」
	"形容詞,接尾,*,*": append(adjectiveEDictTypes, "&suf;"),

	// Auxililary Verb
	// ex 「らしい」、「ござる」、「っす」、「じゃん」
	"助動詞,*,*,*": append(verbEDictTypes, []string{
		"&aux-adj;",
		"&aux;",
	}...),
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
	// Adverb General
	// ex 「たいそう」、「人一倍」、「いけしゃあしゃあ」
	"副詞,一般,*,*": adverbEDictTypes,

	// Adverb Connection Particle
	// ex 「あまり」、「いつも」、「ぱさぱさ」
	"副詞,助詞類接続,*,*": adverbEDictTypes,

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

	"フィラー,*,*,*": []string{
		"&adv;",
		"&exp;",
		"&int;",
		"&prt;",
	},

	"その他,間投,*,*": []string{
		"&int;",
		"&prt;",
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

	"鐃緒申鐃銃誌申,鐃緒申立,*,*": []string{
		"&n-suf;",
		"&n;",
		"&pn;",
		"&pref;",
		"&suf;",
	},

	"助詞,副助詞／並立助詞／終助詞,*,*": []string{
		"&adv;",
		"&pref;",
		"&prt;",
		"&suf;",
	},

	"記号,アルファベット,*,*": []string{
		"&n;",
		"&pref;",
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
