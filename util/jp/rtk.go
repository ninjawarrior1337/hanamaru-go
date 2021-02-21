//+build jp

package jp

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"errors"
	"strings"
	"unicode/utf8"
)

//go:embed rtk.csv
var rtkCSVBytes []byte

var ErrorEntryNotFound = errors.New("kanji with the provided search params not found")

type RtkDatabase struct {
	Entries []RtkEntry
}

type RtkEntry struct {
	Kanji   rune
	Number  int
	Keyword string
	Story   string
	Comment string
	Koohi   []string
}

func NewRtkDatabase() *RtkDatabase {
	r := &RtkDatabase{
		Entries: []RtkEntry{},
	}
	r.parseDict()
	return r
}

func (r *RtkDatabase) parseDict() {
	rd := csv.NewReader(bytes.NewReader(rtkCSVBytes))
	rd.Comma = ';'
	rd.LazyQuotes = true
	rd.FieldsPerRecord = -1
	records, err := rd.ReadAll()
	if err != nil {
		panic(err)
	}
	for i, rec := range records {
		re := RtkEntry{
			Kanji:   bytes.Runes([]byte(rec[0]))[0],
			Number:  i + 1,
			Keyword: rec[1],
			Story:   rec[2],
			Comment: rec[3],
			Koohi:   []string{},
		}
		for j := 4; j < len(rec); j++ {
			re.Koohi = append(re.Koohi, rec[j])
		}
		r.Entries = append(r.Entries, re)
	}
}

func (r *RtkDatabase) Search(search string) (RtkEntry, error) {
	if utf8.RuneCountInString(search) == 1 && (int(search[0]) >= 19968 && int(search[0]) <= 40895) {
		return r.searchRune(rune(search[0]))
	} else {
		return r.searchName(search)
	}
}

func (r *RtkDatabase) GetByNumber(index int) RtkEntry {
	return r.Entries[index+1]
}

func (r *RtkDatabase) searchRune(rn rune) (RtkEntry, error) {
	for _, e := range r.Entries {
		if e.Kanji == rn {
			return e, nil
		}
	}
	return RtkEntry{}, ErrorEntryNotFound
}

func (r *RtkDatabase) searchName(name string) (RtkEntry, error) {
	name = strings.ToLower(name)
	for _, e := range r.Entries {
		if strings.Contains(strings.ToLower(e.Keyword), name) {
			return e, nil
		}
	}
	return RtkEntry{}, ErrorEntryNotFound
}
