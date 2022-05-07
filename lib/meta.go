package lib

import (
	"fmt"
	"sort"
	"strings"
)

func Meta(contents string) []string {
	rep1 := strings.NewReplacer("\b", " BACK_SPACE ", "\f", " FORM_STR ", "\n", " NEW_LINE ", "\r", " RETURN ", "\t", " HOR_TAB ", "\v", " VAR_TAB ")
	contents1 := rep1.Replace(contents)
	rep2 := strings.NewReplacer("&", " AND ", ";", " AND ", "`", " BACK_TICK ", "'", " SINGLE_QUATE ", "\"", " DOUBLE_QUATE ", "\\", " ESCAPE_SEQ ", "|", " PIPE ", "~", " TILDE ", "<", " REV_REDIRECT ", ">", " REDIRECT ", "(", " LEFT_BRACKET1 ", ")", " RIGHT_BRACKET1 ", "[", " LEFT_BRACKET2 ", "]", " RIGHT_BRACKET2 ", "{", " LEFT_BRACKET3 ", "}", " RIGHT_BRACKET3 ", "!", " EXC ", "#", " SHAPE ", "$", " DULL ", "%", " PERCENT ", "^", " UNKNOWN ", "≠", " UNKNOWN ", "¥", " EN ", "+", " PLUS ", "…", " UNKNOWN ", ":", " COLON ", "*", " ASTERISK ", "æ", " UNKNOWN ", "«", " UNKNOWN ", "=", " EQUAL ")
	contents2 := rep2.Replace(contents1)
	res := strings.Split(contents2, " ")
	return res
}

func EscMeta(contents string) []string {
	rep1 := strings.NewReplacer("\b", "", "\f", "", "\n", "", "\r", "", "\t", "", "\v", "")
	contents1 := rep1.Replace(contents)
	rep2 := strings.NewReplacer("&", " AND ", ";", " AND ", "`", " BACK_TICK ", "'", " SINGLE_QUATE ", "\"", " DOUBLE_QUATE ", "\\", " ESCAPE_SEQ ", "|", " PIPE ", "~", " TILDE ", "<", " REV_REDIRECT ", ">", " REDIRECT ", "(", " LEFT_BRACKET1 ", ")", " RIGHT_BRACKET1 ", "[", " LEFT_BRACKET2 ", "]", " RIGHT_BRACKET2 ", "{", " LEFT_BRACKET3 ", "}", " RIGHT_BRACKET3 ", "!", " EXC ", "#", " SHAPE ", "$", " DULL ", "%", " PERCENT ", "^", " UNKNOWN ", "≠", " UNKNOWN ", "¥", " EN ", "+", " PLUS ", "…", " UNKNOWN ", ":", " COLON ", "*", " ASTERISK ", "æ", " UNKNOWN ", "«", " UNKNOWN ", "=", " EQUAL ")
	contents2 := rep2.Replace(contents1)
	res := strings.Split(contents2, " ")
	return res
}

func Meta2(contents string) {
	meta := [...]string{
		// "\a",
		"\b",
		"\f",
		"\n",
		"\r",
		"\t",
		"\v",
		"&",
		";",
		"`",
		"'",
		"\"",
		"\\",
		"|",
		"~",
		"<",
		">",
		"(",
		")",
		"[",
		"]",
		"{",
		"}",
		"!",
		"#",
		"$",
		"%",
		"~",
		"^",
		"≠",
		"¥",
		"+",
		"…",
		":",
		"*",
		"æ",
		"«",
	}

	fmt.Println(meta[0])

	metaMap := map[string]string{
		"\b": "BACK_SPACE",
		"\f": "FORM_STR",
		"\n": "NEW_LINE",
		"\r": "RETURN",
		"\t": "HOR_TAB",
		"\v": "VAR_TAB",
		"&":  "AND1",
		";":  "AND2",
		"`":  "BACK_TICK",
		"'":  "SINGLE_QUATE",
		"\"": "DOUBLE_QUATE",
		"\\": "ESCAPE_SEQ",
		"|":  "PIPE",
		"~":  "TILDE",
		"<":  "REV_REDIRECT",
		">":  "REDIRECT",
		"(":  "LEFT_BRACKET1",
		")":  "RIGHT_BRACKET1",
		"[":  "LEFT_BRACKET2",
		"]":  "RIGHT_BRACKET2",
		"{":  "LEFT_BRACKET3",
		"}":  "RIGHT_BRACKET3",
		"!":  "EXC",
		"#":  "SHAPE",
		"$":  "DULL",
		"%":  "PERCENT",
		"^":  "UNKNOWN",
		"≠":  "UNKNOWN",
		"¥":  "EN",
		"+":  "PLUS",
		"…":  "UNKNOWN",
		":":  "COLON",
		"*":  "ASTERISK",
		"æ":  "UNKNOWN",
		"«":  "UNKNOWN",
	}

	fmt.Println(metaMap["*"])

	meta1 := map[string]string{
		"\b": " BACK_SPACE ",
		"\f": " FORM_STR ",
		"\n": " NEW_LINE ",
		"\r": " RETURN ",
		"\t": " HOR_TAB ",
		"\v": " VAR_TAB ",
	}

	meta2 := map[string]string{
		"&":  " AND1 ",
		";":  " AND2 ",
		"`":  " BACK_TICK ",
		"'":  " SINGLE_QUATE ",
		"\"": " DOUBLE_QUATE ",
		"\\": " ESCAPE_SEQ ",
		"|":  " PIPE ",
		"~":  " TILDE ",
		"<":  " REV_REDIRECT ",
		">":  " REDIRECT ",
		"(":  " LEFT_BRACKET1 ",
		")":  " RIGHT_BRACKET1 ",
		"[":  " LEFT_BRACKET2 ",
		"]":  " RIGHT_BRACKET2 ",
		"{":  " LEFT_BRACKET3 ",
		"}":  " RIGHT_BRACKET3 ",
		"!":  " EXC ",
		"#":  " SHAPE ",
		"$":  " DULL ",
		"%":  " PERCENT ",
		"^":  " UNKNOWN ",
		"≠":  " UNKNOWN ",
		"¥":  " EN ",
		"+":  " PLUS ",
		"…":  " UNKNOWN ",
		":":  " COLON ",
		"*":  " ASTERISK ",
		"æ":  " UNKNOWN ",
		"«":  " UNKNOWN ",
		"=":  " EQUAL ",
	}

	// fmt.Println(meta1)
	fmt.Println(meta2["*"])

	fmt.Println("Before2:", contents)

	var meta1Keys []string
	for key := range meta1 {
		meta1Keys = append(meta1Keys, key)
	}
	sort.Strings(meta1Keys)

	for _, meta1Key := range meta1Keys {
		// fmt.Println(meta1Key, meta1[meta1Key])
		fmt.Println("before1:", contents)
		contents := strings.Replace(contents, meta1Key, meta1[meta1Key], -1)
		fmt.Println("after1:", contents)
	}

	var meta2Keys []string
	for key := range meta2 {
		meta2Keys = append(meta2Keys, key)
	}
	sort.Strings(meta2Keys)

	for _, meta2Key := range meta2Keys {
		// fmt.Println("Before2:", contents)
		contents := strings.Replace(contents, meta2Key, meta1[meta2Key], -1)
		fmt.Println("after2:", contents)

	}
	fmt.Println("after:", contents)

}
