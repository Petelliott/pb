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

func Lex(str string) []Token {
	state := -1
	val := ""

	toklist := make([]Token, 0)

	for _, ch := range str {
		if ch == '"' || ch == '\'' {
			val += string(ch)

			if val == "\"" || val == "'" {
				state = STR_LITERAL
			} else if ch == rune(val[0]) {
				toklist = append(toklist, Token{state, val})
				state = -1
				val = ""
			}
		} else if state == STR_LITERAL {
			val += string(ch)
		} else if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' {
			if val == "" && unicode.IsLetter(ch) {
				state = IDENTIFIER
			} else if val == "" && unicode.IsDigit(ch) {
				state = LITERAL
			}

			val += string(ch)
		} else {
			if val != "" && state != -1 {
				toklist = append(toklist, Token{state, val})
			}
			state = -1
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
