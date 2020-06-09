package util

import (
	"fmt"
	"strings"
)

var emojiMap = map[string]string{
	"A": "🇦",
	"B": "🇧",
	"C": "🇨",
	"D": "🇩",
	"E": "🇪",
	"F": "🇫",
	"G": "🇬",
	"H": "🇭",
	"I": "🇮",
	"J": "🇯",
	"K": "🇰",
	"L": "🇱",
	"M": "🇲",
	"N": "🇳",
	"O": "🇴",
	"P": "🇵",
	"Q": "🇶",
	"R": "🇷",
	"S": "🇸",
	"T": "🇹",
	"U": "🇺",
	"V": "🇻",
	"W": "🇼",
	"X": "🇽",
	"Y": "🇾",
	"Z": "🇿",
}

var duplicateMap = map[string]string{
	"A": "🅰️",
	"B": "🅱️",
	"O": "🅾️",
}

type EmojiConversionError struct {
	ConversionString string
	FailIdx          int
}

func (e EmojiConversionError) Error() string {
	return fmt.Sprintf("failed to convert %v to emoji. failed on char %v", e.ConversionString, e.FailIdx)
}

//This will panic if it can't do it.
func MustMapToEmoji(in string) []string {
	emoji, err := MapToEmoji(in)
	if err != nil {
		panic(fmt.Sprintf(`cannot transform "%v" into emoji`, in))
	}
	return emoji
}

func MapToEmoji(in string) (out []string, err error) {
	usedMap := map[string]bool{}
	inUpper := strings.ToUpper(in)
	for idx, char := range strings.Split(inUpper, "") {
		if _, used := usedMap[emojiMap[char]]; used {
			dupChar, ok := duplicateMap[char]
			if !ok {
				return nil, EmojiConversionError{ConversionString: in, FailIdx: idx}
			}
			if _, used := usedMap[dupChar]; used {
				return nil, EmojiConversionError{ConversionString: in, FailIdx: idx}
			} else {
				out = append(out, duplicateMap[char])
				usedMap[duplicateMap[char]] = true
			}
		} else {
			out = append(out, emojiMap[char])
			usedMap[emojiMap[char]] = true
		}
	}
	return
}
