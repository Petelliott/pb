package lexer

import "unicode"
import "fmt"

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

var operators = map[string]bool{
	"+":  true,
	"-":  true,
	"*":  true,
	"/":  true,
	"&":  true,
	"|":  true,
	"^":  true,
	"~":  true,
	"%":  true,
	"!":  true,
	"!=": true,
	"=":  true,
	"&&": true,
	"||": true,
	"^^": true,
	"==": true,
	"+=": true,
	"-=": true,
	">=": true,
	"<=": true,
	">":  true,
	"<":  true,
	">>": true,
	"<<": true,
}

var multi_operator = map[string]bool{
	"+": true,
	"-": true,
	"=": true,
	"&": true,
	"|": true,
	"^": true,
	">": true,
	"<": true,
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
	} else if operators[val] {
		return Token{OPERATOR, val}
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
			if in_str && ch == rune(val[0]) {
				val += string(ch)
				if ch == rune(val[0]) {
					in_str = false
				}
			} else {
				if val != "" {
					toklist = append(toklist, getToken(val))
					val = ""
				}

				val = string(ch)
				in_str = true
			}
		} else if in_str {
			val += string(ch)
		} else if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' {
			val += string(ch)

		} else if multi_operator[string(ch)] || operators[string(ch)] {
			op := val + string(ch)
			if multi_operator[op] {
				val = op
			} else if operators[op] {
				val = op
			} else {
				toklist = append(toklist, getToken(val))
				val = string(ch)
			}
		} else if toktype, in := control[ch]; in {
			if val != "" {
				toklist = append(toklist, getToken(val))
				val = ""
			}
			toklist = append(toklist, Token{toktype, ""})
		} else if unicode.IsSpace(ch) {
			if val != "" {
				toklist = append(toklist, getToken(val))
				val = ""
			}
		} else {
			return nil
		}
		fmt.Println(val)
	}

	if val != "" {
		toklist = append(toklist, getToken(val))
	}

	return toklist
}
