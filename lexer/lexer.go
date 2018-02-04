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
	',': COMMA,
	'.': DOT,
}

var keywords = map[string]int{
	"for":    KW_FOR,
	"while":  KW_WHILE,
	"if":     KW_IF,
	"else":   KW_ELSE,
	"func":   KW_FUNC,
	"return": KW_RETURN,
}

var types = map[string]bool{
	"word": true,
	"atom": true,
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
		if types[val] {
			return Token{TYPE, val}
		} else if tok, ok := keywords[val]; ok {
			return Token{tok, ""}
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
			if in_str {
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
			var rune0 rune
			for _, r := range val {
				rune0 = r
				break
			}

			if !(unicode.IsLetter(rune0) || unicode.IsDigit(rune0) || rune0 == '_') {
				if val != "" {
					toklist = append(toklist, getToken(val))
					val = ""
				}
			}
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
	}

	if val != "" {
		toklist = append(toklist, getToken(val))
	}

	return toklist
}
