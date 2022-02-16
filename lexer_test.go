package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var hello_world_tokens = []Token {
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

func TestLexerHelloWorld(t *testing.T) {
	assert := assert.New(t)
	file := "./snippets/hello_world.ll"
	
	var tokens []Token

	lexerer := lexer(file)
	for !leof(&lexerer) {
		token := lpeek(&lexerer)

		tokens = append(tokens, token)

		lpop(&lexerer)
	}

	assert.Equal(hello_world_tokens, tokens, "Lexer tokenizes 'hello_world' correctly")
}

var hello_human_tokens = []Token {
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

func TestLexerHelloHuman(t *testing.T) {
	assert := assert.New(t)
	file := "./snippets/hello_human.ll"
	
	var tokens []Token
	
	lexerer := lexer(file)
	for !leof(&lexerer) {
		token := lpeek(&lexerer)

		tokens = append(tokens, token)

		lpop(&lexerer)
	}

	assert.Equal(hello_human_tokens, tokens, "Lexer tokenizes 'hello_human' correctly")
}



var greet_human_tokens = []Token {
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

func TestLexerGreetHuman(t *testing.T) {
	assert := assert.New(t)
	file := "./snippets/greet_human.ll"
	
	var tokens []Token
	
	lexerer := lexer(file)
	for !leof(&lexerer) {
		token := lpeek(&lexerer)

		tokens = append(tokens, token)

		lpop(&lexerer)
	}

	assert.Equal(greet_human_tokens, tokens, "Lexer tokenizes 'greet_human' correctly")
}
