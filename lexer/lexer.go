package lexer

import "unicode"

var control = map[rune]int{
	'(': L_PAREN,
	')': R_PAREN,
	'[': L_BRACKET,
	']': R_BRACKET,
	'{': L_BRACE,
	'}': R_BRACE,
	';': SEMICOLON,
}

var keywords = map[string]bool{
	"for":   true,
	"while": true,
	"if":    true,
	"else":  true,
	"word":  true,
	"atom":  true,
	"func":  true,
}

func getToken(val string) Token {
	var rune0 rune
	for _, r := range val {
		rune0 = r
		break
	}

	if unicode.IsDigit(rune0) || rune0 == '\'' || rune0 == '"' {
		return Token{LITERAL, val}

	} else if unicode.IsLetter(rune0) {

		if keywords[val] {
			return Token{KEYWORD, val}
		} else {
			return Token{IDENTIFIER, val}
		}
	} else {
		return Token{-1, "ERROR"}
	}
}

func Lex(str string) []Token {
	in_str := false
	val := ""

	toklist := make([]Token, 0)

	for _, ch := range str {
		if ch == '"' || ch == '\'' {

			val += string(ch)

			if val == "\"" || val == "'" {
				in_str = true
			} else if ch == rune(val[0]) {
				toklist = append(toklist, getToken(val))
				in_str = false
				val = ""
			}
		} else if in_str {
			val += string(ch)
		} else if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' {
			val += string(ch)
		} else {

			if val != "" {
				toklist = append(toklist, getToken(val))
			}
			val = ""

			if !unicode.IsSpace(ch) {
				toktype, in := control[ch]
				if in {
					toklist = append(toklist, Token{toktype, ""})
				} else {
					toklist = append(toklist, Token{OPERATOR, string(ch)})
				}
			}
		}
	}

	return toklist
}
