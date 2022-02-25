package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var output_tokens = []Token {
	{
		category: "kw",
		value: "out",
	},
	{
		category: "punc",
		value: "(",
	},
	{
		category: "str",
		value: "Hello World",
	},
	{
		category: "punc",
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
		category: "kw",
		value: "const",
	},
	{
		category: "var",
		value: "name",
	},
	{
		category: "op",
		value: "=",
	},
	{
		category: "kw",
		value: "in",
	},
	{
		category: "punc",
		value: "(",
	},
	{
		category: "str",
		value: "Whats your name?",
	},
	{
		category: "punc",
		value: ")",
	},
	{
		category: "kw",
		value: "out",
	},
	{
		category: "punc",
		value: "(",
	},
	{
		category: "str",
		value: "Hello ",
	},
	{
		category: "op",
		value: "+",
	},
	{
		category: "var",
		value: "name",
	},
	{
		category: "punc",
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
		category: "kw",
		value: "const",
	},
	{
		category: "var",
		value: "greet",
	},
	{
		category: "op",
		value: "=",
	},
	{
		category: "punc",
		value: "(",
	},
	{
		category: "var",
		value: "name",
	},
	{
		category: "punc",
		value: ")",
	},
	{
		category: "op",
		value: "=",
	},
	{
		category: "op",
		value: ">",
	},
	{
		category: "punc",
		value: "{",
	},
	{
		category: "kw",
		value: "out",
	},
	{
		category: "punc",
		value: "(",
	},
	{
		category: "str",
		value: "Hello ",
	},
	{
		category: "op",
		value: "+",
	},
	{
		category: "var",
		value: "name",
	},
	{
		category: "punc",
		value: ")",
	},
	{
		category: "punc",
		value: "}",
	},
	{
		category: "kw",
		value: "const",
	},
	{
		category: "var",
		value: "name",
	},
	{
		category: "op",
		value: "=",
	},
	{
		category: "kw",
		value: "in",
	},
	{
		category: "punc",
		value: "(",
	},
	{
		category: "str",
		value: "Whats your name?",
	},
	{
		category: "punc",
		value: ")",
	},
	{
		category: "var",
		value: "greet",
	},
	{
		category: "punc",
		value: "(",
	},
	{
		category: "var",
		value: "name",
	},
	{
		category: "punc",
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
		category: "kw",
		value: "const",
	},
	{
		category: "var",
		value: "data",
	},
	{
		category: "op",
		value: "=",
	},
	{
		category: "punc",
		value: "[",
	},
	{
		category: "num",
		value: 1,
	},
	{
		category: "punc",
		value: ",",
	},
	{
		category: "num",
		value: 2,
	},
	{
		category: "punc",
		value: ",",
	},
	{
		category: "num",
		value: 3,
	},
	{
		category: "punc",
		value: "]",
	},
	{
		category: "kw",
		value: "for",
	},
	{
		category: "punc",
		value: "(",
	},
	{
		category: "var",
		value: "element",
	},
	{
		category: "kw",
		value: "of",
	},
	{
		category: "var",
		value: "data",
	},
	{
		category: "punc",
		value: ")",
	},
	{
		category: "punc",
		value: "{",
	},
	{
		category: "kw",
		value: "out",
	},
	{
		category: "punc",
		value: "(",
	},
	{
		category: "var",
		value: "element",
	},
	{
		category: "op",
		value: "+",
	},
	{
		category: "num",
		value: 10,
	},
	{
		category: "punc",
		value: ")",
	},
	{
		category: "punc",
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
		category: "kw",
		value: "const",
	},
	{
		category: "var",
		value: "animal",
	},
	{
		category: "op",
		value: "=",
	},
	{
		category: "kw",
		value: "in",
	},
	{
		category: "punc",
		value: "(",
	},
	{
		category: "str",
		value: "Whats your favourite animal?",
	},
	{
		category: "punc",
		value: ")",
	},
	{
		category: "kw",
		value: "if",
	},
	{
		category: "punc",
		value: "(",
	},
	{
		category: "var",
		value: "animal",
	},
	{
		category: "op",
		value: "==",
	},
	{
		category: "str",
		value: "frog",
	},
	{
		category: "punc",
		value: ")",
	},
	{
		category: "punc",
		value: "{",
	},
	{
		category: "kw",
		value: "out",
	},
	{
		category: "punc",
		value: "(",
	},
	{
		category: "str",
		value: "A man who enjoys culture I see",
	},
	{
		category: "punc",
		value: ")",
	},
	{
		category: "punc",
		value: "}",
	},
	{
		category: "kw",
		value: "else",
	},
	{
		category: "punc",
		value: "{",
	},
	{
		category: "kw",
		value: "out",
	},
	{
		category: "punc",
		value: "(",
	},
	{
		category: "str",
		value: "I prefer frogs",
	},
	{
		category: "punc",
		value: ")",
	},
	{
		category: "punc",
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