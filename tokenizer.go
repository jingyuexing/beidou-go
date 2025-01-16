package beidou

import (
	"strings"
)

// SpecialCharacter holds various special characters
type SpecialCharacter struct {
	BEGIN     string
	END       string
	NEWLINE   string
	CHECKSUM  string
	DELIMITER string
	ENCODE    string
	Key_1     string
	BACKSLASH string
	DELETE    string
}

var SpecialChracter = SpecialCharacter{
	BEGIN:     "\x24",
	END:       "\x0D",
	NEWLINE:   "\x0A",
	CHECKSUM:  "\x2A",
	DELIMITER: "\x2C",
	ENCODE:    "\x5E",
	Key_1:     "\x7E",
	BACKSLASH: "\x5C",
	DELETE:    "\x7F",
}

// OutInputCharacter defines possible output input characters
type OutInputChracter string

const (
	EPI    OutInputChracter = "EPI"
	TCI    OutInputChracter = "TCI"
	SHZ    OutInputChracter = "SHZ"
	SAK    OutInputChracter = "SAK"
	ACK    OutInputChracter = "ACK"
	TCK    OutInputChracter = "TCK"
	PWI    OutInputChracter = "PWI"
	ICZOut OutInputChracter = "ICZ"
	FKZOut OutInputChracter = "FKZ"
)

// InputCharacter defines possible input characters
type InputChracter string

const (
	EPQ InputChracter = "EPQ"
	TCQ InputChracter = "TCQ"
	PSQ InputChracter = "PSQ"
	RTQ InputChracter = "RTQ"
	OFQ InputChracter = "OFQ"
	KSQ InputChracter = "KSQ"
	PWS InputChracter = "PWS"
	ICZ InputChracter = "ICZ"
	FKZ InputChracter = "FKZ"
)

// TokenKind defines different kinds of tokens
type TokenKind int

const (
	Begin TokenKind = iota + 1
	End
	NewLine
	Checksum
	Spiliter
	Encode
	Number
	Text
)

// Token represents a single token with a type and value
type Token struct {
	Type  TokenKind
	Value string
}

// isValidCharacter checks if the character is a valid character
func isValidCharacter(ch string) bool {
	return (ch >= "\x20" && ch <= "\x7f" && !strings.ContainsAny(
		ch,
		SpecialChracter.BACKSLASH+SpecialChracter.BEGIN+SpecialChracter.CHECKSUM+SpecialChracter.DELETE+SpecialChracter.ENCODE+SpecialChracter.DELIMITER+SpecialChracter.NEWLINE+SpecialChracter.END)) || ch > "\xff"
}

// isLetter checks if the character is a letter
func isLetter(ch string) bool {
	return (ch >= "a" && ch <= "z") || (ch >= "A" && ch <= "Z")
}

// isNumeric checks if the character is a numeric value
func isNumeric(ch string) bool {
	return ch >= "0" && ch <= "9"
}

// Tokenizer splits the input text into tokens
func Tokenizer(text string) []Token {
	var tokens []Token
	current := 0
	for current < len(text) {
		ch := string(text[current])

		if isValidCharacter(ch) {
			val := ""
			if isNumeric(ch) {
				for current < len(text) && (isNumeric(string(text[current])) || string(text[current]) == ".") {
					val += string(text[current])
					current++
				}
				tokens = append(tokens, Token{Type: Number, Value: val})
			} else {
				for current < len(text) && isValidCharacter(string(text[current])) {
					val += string(text[current])
					current++
				}
				tokens = append(tokens, Token{Type: Text, Value: val})
			}
		} else {
			switch ch {
			case SpecialChracter.BEGIN:
				tokens = append(tokens, Token{Type: Begin, Value: ch})
				current++
			case SpecialChracter.END:
				tokens = append(tokens, Token{Type: End, Value: ch})
				current++
			case SpecialChracter.DELIMITER:
				tokens = append(tokens, Token{Type: Spiliter, Value: ch})
				current++
			case SpecialChracter.CHECKSUM:
				tokens = append(tokens, Token{Type: Checksum, Value: ch})
				current++
			case SpecialChracter.NEWLINE:
				tokens = append(tokens, Token{Type: NewLine, Value: ch})
				current++
			}
		}
	}
	return tokens
}
