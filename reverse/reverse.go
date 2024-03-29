package reverse

import "strings"

var ReverseMap = map[string]string{
	"a":  "ɐ",
	"b":  "q",
	"c":  "ɔ",
	"d":  "p",
	"e":  "ǝ",
	"f":  "ɟ",
	"g":  "ƃ",
	"h":  "ɥ",
	"i":  "ᴉ",
	"j":  "ɾ",
	"k":  "ʞ",
	"l":  "ʅ",
	"m":  "ɯ",
	"n":  "u",
	"o":  "o",
	"p":  "d",
	"q":  "b",
	"r":  "ɹ",
	"s":  "s",
	"t":  "ʇ",
	"u":  "n",
	"v":  "ʌ",
	"w":  "ʍ",
	"x":  "x",
	"y":  "ʎ",
	"z":  "z",
	"A":  "∀",
	"B":  "𐐒",
	"C":  "Ɔ",
	"D":  "ᗡ",
	"E":  "Ǝ",
	"F":  "Ⅎ",
	"G":  "⅁",
	"H":  "H",
	"I":  "I",
	"J":  "ſ",
	"K":  "⋊",
	"L":  "˥",
	"M":  "W",
	"N":  "N",
	"O":  "O",
	"P":  "Ԁ",
	"Q":  "Ό",
	"R":  "ᴚ",
	"S":  "S",
	"T":  "⊥",
	"U":  "∩",
	"V":  "Λ",
	"W":  "M",
	"X":  "X",
	"Y":  "⅄",
	"Z":  "Z",
	"(":  ")",
	")":  "(",
	"?":  "¿",
	"!":  "¡",
	"<":  ">",
	",":  "`",
	"`":  ",",
	"[":  "]",
	"]":  "[",
	"{":  "}",
	"}":  "{",
	"\\": "/",
	"/":  "\\",
	"~":  "~",
	"@":  "@",
	"#":  "#",
	"$":  "$",
	"%":  "%",
	"^":  "^",
	"&":  "&",
	"*":  "*",
	"-":  "-",
	"_":  "_",
	"+":  "+",
	"=":  "=",
	"|":  "|",
	":":  ":",
	";":  ";",
	"'":  ",",
	"\"": "\"",
	">":  "<",
	".":  "˙",
	" ":  " ",
}

func Reverse(org string) string {
	var result strings.Builder
	b := strings.Split(org, "")
	for i := len(b) - 1; i >= 0; i-- {
		if v, ok := ReverseMap[b[i]]; ok {
			result.WriteString(v)
		} else {
			result.WriteString(b[i])
		}
	}
	return result.String()
}
