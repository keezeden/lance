package main

import (
	"fmt"
	"regexp"
	"strings"
)

var separators = []string{"(", ")", "{", "}"}
var keywords = []string{"const", "if", "else", "out", "in"}
var operators = []string{"=", "==", "!=", "+", "-", "/", "*"}

type Lexer struct {
	index int
	buffer []byte
	stream Stream
	tokens []Token
}

type Token struct {
	category string
	value string
}

// TODO: optimize with FSM for scanning possible next token first
func evaluate(chars []byte, tokens []Token) (Token, bool) {
	segment := string(chars)
	// new lines
	new_line_match, _ := regexp.MatchString("(\r\n|\r|\n)", segment)
	if (new_line_match) {
		return Token{ value: segment, category: "separator"}, true
	}
	// keywords
	if (contains(segment, keywords)) {
		return Token{ value: segment, category: "keyword"}, true
	}
	// separators
	if (contains(segment, separators)) {
		return Token{ value: segment, category: "separator"}, true
	}
	// operators
	if (contains(segment, operators)) {
		return Token{ value: segment, category: "operator"}, true
	}
	// strings
	strings_match, _ := regexp.MatchString("\"(.*?)\"", segment)
	if (strings_match) {
		return Token{ value: strings.Replace(segment, "\"", "", -1), category: "string"}, true
	}
	// variables
	variable, _ := regexp.MatchString("[a-zA-Z]+", segment)
	if (variable && tokens[:-1]) {
		return Token{ value: segment, category: "variable"}, true
	}


	return Token{}, false
}


func lpeek(l *Lexer) Token {
	for !seof(&l.stream) {
		char := speek(&l.stream)
		l.buffer = append(l.buffer, char)
		fmt.Println(string(l.buffer))

		_, found := evaluate(l.buffer, l.tokens)
		if (found) {
			spop(&l.stream)
			break
		}


		spop(&l.stream)
	}

	token, _ := evaluate(l.buffer)
	lpop(l)
	return token
}


func lpop(l *Lexer) {
	l.buffer = []byte{}
}

func leof(l *Lexer) bool {
	return seof(&l.stream)
}


func lexer(file string) Lexer {
	streamer := stream(file)
	var tokens []Token

	return Lexer{
		index: 0,
		buffer: []byte{},
		stream: streamer,
		tokens: tokens,
	}
}