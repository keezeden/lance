package main

import (
	"fmt"
	"regexp"
	"strings"
)

var separators = regexp.MustCompile("[(){}]")
var keywords = regexp.MustCompile("const|if|else|out|in")
var operators = regexp.MustCompile("\\=|\\==|\\!=|\\+|\\-|\\/|\\*")
var strs = regexp.MustCompile("\"(.*?)\"")
var indentifiers = regexp.MustCompile("[a-zA-Z]+")

var strop = regexp.MustCompile("@(.*?)@")

func replace_separators(source string) string {
	return separators.ReplaceAllString(source, "@$0@")
} 
func replace_keywords(source string) string {
	return keywords.ReplaceAllString(source, "@$0@")
} 
func replace_operators(source string) string {
	return operators.ReplaceAllString(source, "@$0@")
} 
func replace_strings(source string) string {
	return strs.ReplaceAllString(source, "@$0@")
} 
func replace_indentifiers(source string) string {
	var buffer []string
	blank := strop.ReplaceAllString(source, "")

	buffer = indentifiers.FindAllString(blank, -1)
	
	var clean = removeDuplicateStr(buffer)
	for _, id := range(clean) {
		source = strings.ReplaceAll(source, id, "@" + id + "@")
	}

	return source
} 



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
	if (matches(segment, *keywords)) {
		return Token{ value: segment, category: "keyword"}, true
	}
	// separators
	if (matches(segment, *separators)) {
		return Token{ value: segment, category: "separator"}, true
	}
	// operators
	if (matches(segment, *operators)) {
		return Token{ value: segment, category: "operator"}, true
	}
	// strings
	if (matches(segment, *strs)) {
		return Token{ value: strings.Replace(segment, "\"", "", -1), category: "string"}, true
	}
	// variables
	if (matches(segment, *indentifiers)) {
		return Token{ value: segment, category: "variable"}, true
	}

	return Token{}, false
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

	var stropped  = string(lines)

	var patterns = []func(string) string{
		replace_separators,
		replace_keywords,
		replace_operators,
		replace_strings,
		replace_indentifiers,
	}
	

	for _, pattern := range(patterns) {
		stropped = pattern(stropped)
	}

	fmt.Println(stropped)

	re := regexp.MustCompile(`@(.*?)@`)
    buffer = re.FindAllString(stropped, -1)
	
	fmt.Println(strings.ReplaceAll(strings.Join(buffer, " "), "@", ""))

	return Lexer{
		index: 0,
		buffer: buffer,
		tokens: tokens,
	}
}