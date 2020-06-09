package jisho

import "fmt"

type KanjiNotFound struct {
	Arg string
}

func (e *KanjiNotFound) Error() string {
	return fmt.Sprintf("kanji not found: %v", e.Arg)
}
