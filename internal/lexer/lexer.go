package lexer

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/keezeden/lance/internal/stream"
	"github.com/keezeden/lance/pkg/utils"
)

var separators = regexp.MustCompile("[(){}\\[\\]\\,\\.]")
var keywords = regexp.MustCompile("\\b(const|let|if|else|out|in|while)\\b")
var operators = regexp.MustCompile("(=){1,2}|\\!=|\\+|\\-|\\/|\\*|\\>|\\<")
var strs = regexp.MustCompile("\"(.*?)\"")
var ints = regexp.MustCompile("\\d+")
var indentifiers = regexp.MustCompile("[a-zA-Z]+")

var strop = regexp.MustCompile("@(.*?)@")

func replace_separators(source string) string {
	return separators.ReplaceAllString(source, "@$0@")
} 
func replace_keywords(source string) string {
	strings_cleaned, string_literals := RemoveStrings(source)
	updated := keywords.ReplaceAllString(strings_cleaned, "@$0@")
	return AddStrings(updated, string_literals)
} 
func replace_operators(source string) string {
	return operators.ReplaceAllString(source, "@$0@")
} 
func replace_strings(source string) string {
	return strs.ReplaceAllString(source, "@$0@")
} 
func replace_ints(source string) string {
	return ints.ReplaceAllString(source, "@$0@")
} 
func replace_indentifiers(source string) string {
	strings_cleaned, string_literals := RemoveStrings(source)

	blank := strop.ReplaceAllString(strings_cleaned, "     ")	

	var clean = utils.RemoveDuplicateStrings(indentifiers.FindAllString(blank, -1))
	
	for _, id := range(clean) {
		re := regexp.MustCompile("\\b" + id + "\\b")
		strings_cleaned = re.ReplaceAllString(strings_cleaned, "@" + id + "@")
	}
	
	return AddStrings(strings_cleaned, string_literals)
} 

func RemoveStrings(source string) (string, [][]string) {
	string_literals := strs.FindAllStringSubmatch(source, -1)
	strings_cleaned := strs.ReplaceAllString(source, "$")
	return strings_cleaned, string_literals
}

func AddStrings(source string, literals [][]string) string {
	var buffer string = source

	var final string
	var counter = 0

	for _, char := range(buffer) {
		if (string(char) == "$") {
			final = final + literals[counter][0]
			counter++
		} else {
			final = final + string(char)
		}
	}
	
	return final
}


type Lexer struct {
	index int
	buffer []string
	tokens []Token
}

type Token struct {
	Type string
	Value interface{}
}

// TODO: optimize with FSM for scanning possible next token first
func evaluate(segment string) (Token, bool) {
	// keywords
	if (utils.Matches(segment, *keywords)) {
		return Token{ Value: segment, Type: "kw"}, true
	}
	// separators
	if (utils.Matches(segment, *separators)) {
		return Token{ Value: segment, Type: "punc"}, true
	}
	// operators
	if (utils.Matches(segment, *operators)) {
		return Token{ Value: segment, Type: "op"}, true
	}
	// strings
	if (utils.Matches(segment, *strs)) {
		return Token{ Value: strings.Replace(segment, "\"", "", -1), Type: "str"}, true
	}
	//ints
	if (utils.Matches(segment, *ints)) {
		parsed, _ := strconv.Atoi(segment)
		return Token{ Value: parsed, Type: "num"}, true
	}
	// variables
	if (utils.Matches(segment, *indentifiers)) {
		return Token{ Value: segment, Type: "var"}, true
	}

	return Token{}, false
}

func (l *Lexer) Peek() Token {
	token, _ := evaluate(l.buffer[l.index])

	return token
}


func (l *Lexer) Pop() {
	l.index++
}

func (l *Lexer) Eof() bool {
	return l.index == len(l.buffer) 
}


func BuildLexer(file string) Lexer {
	streamer := stream.BuildStream(file)
	var tokens []Token
	var lines []byte
	var buffer []string

	for !streamer.Eof() {
		char := streamer.Peek()
		lines = append(lines, char)

		streamer.Pop()
	}

	var stropped  = string(lines)

	var patterns = []func(string) string{
		replace_separators,
		replace_keywords,
		replace_operators,
		replace_strings,
		replace_ints,
		replace_indentifiers,
	}

	for _, pattern := range(patterns) {
		stropped = pattern(stropped)
	}
	

	

	for _, word := range(strop.FindAllString(stropped, -1)) {
		buffer = append(buffer, strings.ReplaceAll(word, "@", ""))
	}


	return Lexer{
		index: 0,
		buffer: buffer,
		tokens: tokens,
	}
}