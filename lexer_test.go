package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var content = []Token {
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

func TestLexer(t *testing.T) {
	assert := assert.New(t)
	file := "./snippets/hello_world.ll"
	
	var tokens []Token
	
	lexerer := lexer(file)
	for !leof(&lexerer) {
		token := lpeek(&lexerer)

		tokens = append(tokens, token)

		lpop(&lexerer)
	}

	assert.Equal(content, tokens, "Lexer tokenizes correctly")
}

