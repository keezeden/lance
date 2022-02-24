package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var math_ast = map[string]interface{}{
	"type": "prog",
	"body": map[string]interface{}{
		"": "",
	},
}

func TestParserMath(t *testing.T) {
	assert := assert.New(t)
	file := "./snippets/math.ll"

	lexerer := lexer(file)
	parser := parser(lexerer)

	ast := parser.parse()

	assert.Equal(ast, math_ast, "Parser parses 'math' correctly")
}