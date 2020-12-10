package reverse

import "strings"

var ReverseMap = map[string]string{
	"m":  "ɯ",
	"n":  "u",
	"b":  "q",
	"v":  "ʌ",
	"c":  "ɔ",
	"x":  "x",
	"z":  "z",
	"l":  "l",
	"k":  "ʞ",
	"j":  "ɾ",
	"h":  "ɥ",
	"g":  "ƃ",
	"f":  "ɟ",
	"d":  "p",
	"s":  "s",
	"a":  "ɐ",
	"p":  "d",
	"o":  "o",
	"i":  "ı",
	"u":  "n",
	"y":  "ʎ",
	"t":  "ʇ",
	"r":  "ɹ",
	"e":  "ǝ",
	"w":  "ʍ",
	"q":  "b",
	"(":  ")",
	")":  "(",
	"?":  "¿",
	"!":  "¡",
	"<":  ">",
	",":  "`",
	"`":  ",",
	"[":  "]",
	"{":  "}",
	"}":  "{",
	"]":  "[",
	"\\": "/",
	"/":  "\\",
}

func Reverse(org string) (result string) {
	b := strings.Split(org, "")
	var a = make([]string, 0)
	for i := len(b) - 1; i >= 0; i = i - 1 {
		if v, ok := ReverseMap[b[i]]; ok {
			a = append(a, v)
		} else {
			a = append(a, b[i])
		}
	}
	return strings.Join(a, "")
}
