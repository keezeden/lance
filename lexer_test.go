package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var output_tokens = []Token {
	{
		category: "keyword",
		value: "out",
	},
	{
		category: "separator",
		value: "(",
	},
	{
		category: "string",
		value: "Hello World",
	},
	{
		category: "separator",
		value: ")",
	},
}

func TestLexerOutput(t *testing.T) {
	assert := assert.New(t)
	file := "./snippets/output.ll"
	
	var tokens []Token

	lexerer := lexer(file)
	for !leof(&lexerer) {
		token := lpeek(&lexerer)

		tokens = append(tokens, token)

		lpop(&lexerer)
	}

	assert.Equal(output_tokens, tokens, "Lexer tokenizes 'output' correctly")
}

var input_tokens = []Token {
	{
		category: "keyword",
		value: "const",
	},
	{
		category: "variable",
		value: "name",
	},
	{
		category: "operator",
		value: "=",
	},
	{
		category: "keyword",
		value: "in",
	},
	{
		category: "separator",
		value: "(",
	},
	{
		category: "string",
		value: "Whats your name?",
	},
	{
		category: "separator",
		value: ")",
	},
	{
		category: "keyword",
		value: "out",
	},
	{
		category: "separator",
		value: "(",
	},
	{
		category: "string",
		value: "Hello ",
	},
	{
		category: "operator",
		value: "+",
	},
	{
		category: "variable",
		value: "name",
	},
	{
		category: "separator",
		value: ")",
	},
}

func TestLexerInput(t *testing.T) {
	assert := assert.New(t)
	file := "./snippets/input.ll"
	
	var tokens []Token
	
	lexerer := lexer(file)
	for !leof(&lexerer) {
		token := lpeek(&lexerer)

		tokens = append(tokens, token)

		lpop(&lexerer)
	}

	assert.Equal(input_tokens, tokens, "Lexer tokenizes 'input' correctly")
}



var function_tokens = []Token {
	{
		category: "keyword",
		value: "const",
	},
	{
		category: "variable",
		value: "greet",
	},
	{
		category: "operator",
		value: "=",
	},
	{
		category: "separator",
		value: "(",
	},
	{
		category: "variable",
		value: "name",
	},
	{
		category: "separator",
		value: ")",
	},
	{
		category: "operator",
		value: "=",
	},
	{
		category: "operator",
		value: ">",
	},
	{
		category: "separator",
		value: "{",
	},
	{
		category: "keyword",
		value: "out",
	},
	{
		category: "separator",
		value: "(",
	},
	{
		category: "string",
		value: "Hello ",
	},
	{
		category: "operator",
		value: "+",
	},
	{
		category: "variable",
		value: "name",
	},
	{
		category: "separator",
		value: ")",
	},
	{
		category: "separator",
		value: "}",
	},
	{
		category: "keyword",
		value: "const",
	},
	{
		category: "variable",
		value: "name",
	},
	{
		category: "operator",
		value: "=",
	},
	{
		category: "keyword",
		value: "in",
	},
	{
		category: "separator",
		value: "(",
	},
	{
		category: "string",
		value: "Whats your name?",
	},
	{
		category: "separator",
		value: ")",
	},
	{
		category: "variable",
		value: "greet",
	},
	{
		category: "separator",
		value: "(",
	},
	{
		category: "variable",
		value: "name",
	},
	{
		category: "separator",
		value: ")",
	},
}

func TestLexerFunction(t *testing.T) {
	assert := assert.New(t)
	file := "./snippets/function.ll"
	
	var tokens []Token
	
	lexerer := lexer(file)
	for !leof(&lexerer) {
		token := lpeek(&lexerer)

		tokens = append(tokens, token)

		lpop(&lexerer)
	}

	assert.Equal(function_tokens, tokens, "Lexer tokenizes 'function' correctly")
}


var loop_tokens = []Token {
	{
		category: "keyword",
		value: "const",
	},
	{
		category: "variable",
		value: "data",
	},
	{
		category: "operator",
		value: "=",
	},
	{
		category: "separator",
		value: "[",
	},
	{
		category: "int",
		value: 1,
	},
	{
		category: "separator",
		value: ",",
	},
	{
		category: "int",
		value: 2,
	},
	{
		category: "separator",
		value: ",",
	},
	{
		category: "int",
		value: 3,
	},
	{
		category: "separator",
		value: "]",
	},
	{
		category: "keyword",
		value: "for",
	},
	{
		category: "separator",
		value: "(",
	},
	{
		category: "variable",
		value: "element",
	},
	{
		category: "keyword",
		value: "of",
	},
	{
		category: "variable",
		value: "data",
	},
	{
		category: "separator",
		value: ")",
	},
	{
		category: "separator",
		value: "{",
	},
	{
		category: "keyword",
		value: "out",
	},
	{
		category: "separator",
		value: "(",
	},
	{
		category: "variable",
		value: "element",
	},
	{
		category: "operator",
		value: "+",
	},
	{
		category: "int",
		value: 10,
	},
	{
		category: "separator",
		value: ")",
	},
	{
		category: "separator",
		value: "}",
	},
}

func TestLexerLoop(t *testing.T) {
	assert := assert.New(t)
	file := "./snippets/loop.ll"
	
	var tokens []Token

	lexerer := lexer(file)
	for !leof(&lexerer) {
		token := lpeek(&lexerer)

		tokens = append(tokens, token)

		lpop(&lexerer)
	}

	assert.Equal(loop_tokens, tokens, "Lexer tokenizes 'loop' correctly")
}

var conditional_tokens = []Token{
	{
		category: "keyword",
		value: "const",
	},
	{
		category: "variable",
		value: "animal",
	},
	{
		category: "operator",
		value: "=",
	},
	{
		category: "keyword",
		value: "in",
	},
	{
		category: "separator",
		value: "(",
	},
	{
		category: "string",
		value: "Whats your favourite animal?",
	},
	{
		category: "separator",
		value: ")",
	},
	{
		category: "keyword",
		value: "if",
	},
	{
		category: "separator",
		value: "(",
	},
	{
		category: "variable",
		value: "animal",
	},
	{
		category: "operator",
		value: "==",
	},
	{
		category: "string",
		value: "frog",
	},
	{
		category: "separator",
		value: ")",
	},
	{
		category: "separator",
		value: "{",
	},
	{
		category: "keyword",
		value: "out",
	},
	{
		category: "separator",
		value: "(",
	},
	{
		category: "string",
		value: "A man who enjoys culture I see",
	},
	{
		category: "separator",
		value: ")",
	},
	{
		category: "separator",
		value: "}",
	},
	{
		category: "keyword",
		value: "else",
	},
	{
		category: "separator",
		value: "{",
	},
	{
		category: "keyword",
		value: "out",
	},
	{
		category: "separator",
		value: "(",
	},
	{
		category: "string",
		value: "I prefer frogs",
	},
	{
		category: "separator",
		value: ")",
	},
	{
		category: "separator",
		value: "}",
	},
}

func TestLexerConditional(t *testing.T) {
	assert := assert.New(t)
	file := "./snippets/conditional.ll"
	
	var tokens []Token

	lexerer := lexer(file)
	for !leof(&lexerer) {
		token := lpeek(&lexerer)

		tokens = append(tokens, token)

		lpop(&lexerer)
	}

	assert.Equal(conditional_tokens, tokens, "Lexer tokenizes 'conditional' correctly")
}