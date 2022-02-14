// { type: "punc", value: "(" }           // punctuation: parens, comma, semicolon etc.
// { type: "num", value: 5 }              // numbers
// { type: "str", value: "Hello World!" } // strings
// { type: "kw", value: "lambda" }        // keywords
// { type: "var", value: "a" }            // identifiers
// { type: "op", value: "!=" }            // operators

package main

import (
	"regexp"
	"strings"
)

var separators = []string{"(", ")", "{", "}"}
var keywords = []string{"const", "if", "else", "out", "in"}
var operators = []string{"=", "==", "!=", "+", "-", "/", "*"}

type Lexer struct {
	index int
	pointer int
	stream Stream
	buffer []Token
}

type Token struct {
	category string
	value string
}

// TODO: optimize with FSM for scanning possible next token first
func evaluate(chars []byte) (Token, bool) {
	segment := string(chars)
	if (contains(segment, keywords)) {
		return Token{ value: segment, category: "keyword"}, true
	}
	if (contains(segment, separators)) {
		return Token{ value: segment, category: "separator"}, true
	}
	match, _ := regexp.MatchString("\"(.*?)\"", segment)
	if (match) {
		return Token{ value: strings.Replace(segment, "\"", "", -1), category: "string"}, true
	}


	return Token{}, false
}


func lpeek(l *Lexer) Token {
	var chars []byte

	for !seof(&l.stream) {
		char := speek(&l.stream)
		chars = append(chars, char)

		_, found := evaluate(chars)
		if (found) {
			l.pointer++
			spop(&l.stream)
			break
		}


		l.pointer++
		spop(&l.stream)
	}

	token, _ := evaluate(chars)
	return token
}



func lpop(l *Lexer) {
	l.index = l.pointer
}

func leof(l *Lexer) bool {
	return seof(&l.stream)
}



func lexer(file string) Lexer {
	streamer := stream(file)
	var buffer []Token

	return Lexer{
		index: 0,
		pointer: 0,
		stream: streamer,
		buffer: buffer,
	}
}