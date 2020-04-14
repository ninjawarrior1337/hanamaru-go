// +build jp

package jp

import (
	"strings"

	"github.com/gojp/kana"
	"github.com/ikawaha/kagome/tokenizer"
)

var t = tokenizer.New()

func ParseJapanese(jp string) string {
	tokens := t.Tokenize(jp)
	var finalOutput []string

	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			continue
		}
		features := token.Features()
		inputKana := features[len(features)-1]
		finalOutput = append(finalOutput, kana.KanaToRomaji(inputKana))
	}
	return strings.Join(finalOutput, " ")
}
