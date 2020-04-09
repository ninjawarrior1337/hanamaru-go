package util

import (
	"encoding/csv"
	"strings"
)

func ParseArgs(s string) []string {
	//Magic code from stackoverflow (https://stackoverflow.com/questions/47489745/splitting-a-string-at-space-except-inside-quotation-marks-go)
	s = strings.TrimPrefix(s, " ")
	r := csv.NewReader(strings.NewReader(s))
	r.Comma = ' '
	fields, err := r.Read()
	if err != nil {
		return []string{}
	}

	return fields
}
