import { WordData } from '../types'

const words = [
    {
        "surface": "とても",
        "entries": [
            {
                "sequence": 1008630,
                "kanji": [
                    "迚も"
                ],
                "readings": [
                    "とても",
                    "とっても"
                ],
                "meanings": [
                    "very",
                    "awfully",
                    "exceedingly",
                    "(not) at all",
                    "by no means",
                    "simply (cannot)"
                ],
                "partofspeech": "adverb (fukushi)"
            }
        ],
        "tokens": [
            {
                "id": 48613,
                "class": "KNOWN",
                "surface": "とても",
                "pos": [
                    "副詞",
                    "助詞類接続",
                    "*",
                    "*",
                    "*"
                ],
                "base": "とても",
                "reading": "トテモ",
                "pron": "トテモ",
                "entries": null
            }
        ]
    },
    {
        "surface": "良かった",
        "entries": null,
        "tokens": [
            {
                "id": 327871,
                "class": "KNOWN",
                "surface": "良かっ",
                "pos": [
                    "形容詞",
                    "自立",
                    "*",
                    "*",
                    "形容詞・アウオ段"
                ],
                "base": "良い",
                "reading": "ヨカッ",
                "pron": "ヨカッ",
                "entries": [
                    {
                        "sequence": 1605820,
                        "kanji": [
                            "良い",
                            "善い",
                            "好い",
                            "佳い",
                            "吉い",
                            "宜い"
                        ],
                        "readings": [
                            "よい"
                        ],
                        "meanings": [
                            "good",
                            "excellent",
                            "fine",
                            "nice",
                            "pleasant",
                            "agreeable",
                            "sufficient (can be used to turn down an offer)",
                            "ready",
                            "prepared",
                            "profitable (e.g. deal, business offer, etc.)",
                            "beneficial",
                            "OK"
                        ],
                        "partofspeech": "&adj-i;"
                    }
                ]
            },
            {
                "id": 39233,
                "class": "KNOWN",
                "surface": "た",
                "pos": [
                    "助動詞",
                    "*",
                    "*",
                    "*",
                    "特殊・タ"
                ],
                "base": "た",
                "reading": "タ",
                "pron": "タ",
                "entries": [
                    {
                        "sequence": 1407450,
                        "kanji": [
                            "多"
                        ],
                        "readings": [
                            "た"
                        ],
                        "meanings": [
                            "multi-"
                        ],
                        "partofspeech": "prefix"
                    },
                    {
                        "sequence": 1416830,
                        "kanji": [
                            "誰"
                        ],
                        "readings": [
                            "だれ",
                            "たれ",
                            "た"
                        ],
                        "meanings": [
                            "who"
                        ],
                        "partofspeech": "pronoun"
                    },
                    {
                        "sequence": 1442730,
                        "kanji": [
                            "田"
                        ],
                        "readings": [
                            "た"
                        ],
                        "meanings": [
                            "rice field"
                        ],
                        "partofspeech": "noun (common) (futsuumeishi)"
                    },
                    {
                        "sequence": 1949190,
                        "kanji": [
                            "他"
                        ],
                        "readings": [
                            "た"
                        ],
                        "meanings": [
                            "other (esp. people and abstract matters)"
                        ],
                        "partofspeech": "&adj-no;"
                    },
                    {
                        "sequence": 2654250,
                        "kanji": [],
                        "readings": [
                            "た"
                        ],
                        "meanings": [
                            "did",
                            "(have) done",
                            "(please) do"
                        ],
                        "partofspeech": "&aux-v;"
                    },
                    {
                        "sequence": 2243700,
                        "kanji": [
                            "咫",
                            "尺"
                        ],
                        "readings": [
                            "あた",
                            "た"
                        ],
                        "meanings": [
                            "distance between outstretched thumb and middle finger (approx. 18 cm)"
                        ],
                        "partofspeech": "counter"
                    }
                ]
            }
        ]
    },
    {
        "surface": "です",
        "entries": [
            {
                "sequence": 1628500,
                "kanji": [],
                "readings": [
                    "です"
                ],
                "meanings": [
                    "be",
                    "is"
                ],
                "partofspeech": "expressions (phrases, clauses, etc.)"
            },
            {
                "sequence": 2701430,
                "kanji": [
                    "出洲",
                    "出州"
                ],
                "readings": [
                    "でず",
                    "です"
                ],
                "meanings": [
                    "spit (of land)"
                ],
                "partofspeech": "noun (common) (futsuumeishi)"
            }
        ],
        "tokens": [
            {
                "id": 47492,
                "class": "KNOWN",
                "surface": "です",
                "pos": [
                    "助動詞",
                    "*",
                    "*",
                    "*",
                    "特殊・デス"
                ],
                "base": "です",
                "reading": "デス",
                "pron": "デス",
                "entries": null
            }
        ]
    },
    {
        "surface": "。",
        "entries": null,
        "tokens": [
            {
                "id": 98,
                "class": "KNOWN",
                "surface": "。",
                "pos": [
                    "記号",
                    "句点",
                    "*",
                    "*",
                    "*"
                ],
                "base": "。",
                "reading": "。",
                "pron": "。",
                "entries": null
            }
        ]
    }
] as Array<WordData>;

export default words;