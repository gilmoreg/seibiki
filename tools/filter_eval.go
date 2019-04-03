// Inspired by https://www.thepolyglotdeveloper.com/2017/03/parse-csv-data-go-programming-language/
package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gilmoreg/seibiki/internal/connectors/mongodb"
	"github.com/gilmoreg/seibiki/internal/connectors/redis"
	"github.com/gilmoreg/seibiki/internal/dictionary"
	"github.com/mongodb/mongo-go-driver/bson"
	"go.uber.org/zap"
)

var ipaPath = "../ipa-edict-mapping/ipa"

var result map[string][]edictEntry

func main() {
	start := time.Now()
	result = make(map[string][]edictEntry)
	l := zap.NewExample().Sugar()
	defer l.Sync()
	m := newMongo(l)
	files := loadIPAFiles()
	entries := loadIPAEntries(files)
	jedict := loadJEDict(m)
	fmt.Printf("Read %v IPA entries and %v JEDict entries.\n", len(entries), len(jedict))
	for i, entry := range entries {
		checkEntry(entry, jedict)
		if i%10000 == 0 {
			progress := float64(i) / float64(len(entries)) * 100.0
			fmt.Println(fmt.Sprintf("%.0f%% done", progress))
		}
	}
	jb, _ := json.Marshal(result)
	err := ioutil.WriteFile("log.json", jb, 0644)
	if err != nil {
		panic(err)
	}
	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("done. %s elapsed", elapsed))
}

func checkEntry(token dictionary.Token, jedict []dictionary.Entry) {
	if token.POS[0] == "鐃緒申鐃銃誌申" {
		return
	}
	if token.IsPunctuation() {
		return
	}
	matches := make([]dictionary.Entry, 0)
	for _, entry := range jedict {
		if in(token.Base, entry.Kanji) || in(token.Base, entry.Readings) {
			matches = append(matches, entry)
		}
	}

	// Ignore no matches entirely
	if len(matches) > 0 {
		meanings := 0
		for _, entry := range matches {
			meanings += len(dictionary.Filter(token.POS, entry.Meanings))
		}
		// If no meanings remain after filtering each entry...
		if meanings < 1 {
			if token.POS[3] == "国" {
				fmt.Println("trouble ahead")
			}

			addToMapping(token, matches)
		}
	}
}

func loadJEDict(m mongodb.Client) []dictionary.Entry {
	var result []dictionary.Entry
	raw, err := m.Get(bson.M{})
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(raw, &result)
	if err != nil {
		panic(err)
	}
	return result
}

func loadIPAEntries(files []string) []dictionary.Token {
	var tokens []dictionary.Token
	for _, file := range files {
		if file == ipaPath {
			continue
		}
		csvFile, _ := os.Open(file)
		reader := csv.NewReader(bufio.NewReader(csvFile))
		for {
			line, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			tokens = append(tokens, dictionary.Token{
				Surface: line[0],
				POS:     line[4:8],
				Base:    line[10],
			})
		}
	}
	return tokens
}

func loadIPAFiles() []string {
	var files []string
	err := filepath.Walk(ipaPath, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func newDictionary(l *zap.SugaredLogger, m mongodb.Client) dictionary.Repository {
	c := redis.New("redis://localhost:6379", l)
	return dictionary.New(m, c, l)
}

func newMongo(l *zap.SugaredLogger) mongodb.Client {
	m, err := mongodb.New("mongodb://reader:password@localhost:27017/jedict", l)
	if err != nil {
		panic(err)
	}
	return m
}

type edictEntry struct {
	Surface string   `json:"surface"`
	POS     []string `json:"pos"`
}

// AddPOS -
func (e *edictEntry) AddPOS(str string) {
	for _, s := range e.POS {
		if s == str {
			return
		}
	}
	e.POS = append(e.POS, str)
}

func addToMapping(token dictionary.Token, entries []dictionary.Entry) {
	posString := strings.Join(token.POS, ",")
	_, ok := result[posString]
	if !ok {
		result[posString] = make([]edictEntry, 0)
	}
	toAdd := edictEntry{
		Surface: token.Surface,
		POS:     make([]string, 0),
	}
	for _, entry := range entries {
		for _, meaning := range entry.Meanings {
			for _, pos := range meaning.PartOfSpeech {
				toAdd.AddPOS(pos)
			}
		}
	}
	result[posString] = append(result[posString], toAdd)
}

func in(str string, slice []string) bool {
	for _, t := range slice {
		if str == t {
			return true
		}
	}
	return false
}
