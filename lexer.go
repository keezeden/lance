package main

import (
	"fmt"
	"regexp"
	"strings"
)

var separators = []string{"(", ")", "{", "}", " "}
var keywords = []string{"const", "if", "else", "out", "in"}
var operators = []string{"=", "==", "!=", "+", "-", "/", "*"}

type Lexer struct {
	index int
	buffer []string
	tokens []Token
}

type Token struct {
	category string
	value string
}

// TODO: optimize with FSM for scanning possible next token first
func evaluate(segment string) (Token, bool) {
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
	if (variable) {
		return Token{ value: segment, category: "variable"}, true
	}

	return Token{}, false
}

func split(r rune) bool {
    return contains(string(r), separators)
}

func lpeek(l *Lexer) Token {
	 token, _ := evaluate(l.buffer[l.index])

	 return token
}


func lpop(l *Lexer) {
	l.index++
}

func leof(l *Lexer) bool {
	return l.index == len(l.buffer) 
}


func lexer(file string) Lexer {
	streamer := stream(file)
	var tokens []Token
	var lines []byte
	var buffer []string

	for !seof(&streamer) {
		char := speek(&streamer)
		lines = append(lines, char)

		spop(&streamer)
	}

	re := regexp.MustCompile(`[^\s(){}"']+|([^\s"']*"([^"]*)"[^\s"']*)+|'([^']*)`)
	buffer = re.FindAllString(string(lines), -1)
	fmt.Println(strings.Join(buffer, ", "))

	return Lexer{
		index: 0,
		buffer: buffer,
		tokens: tokens,
	}
}