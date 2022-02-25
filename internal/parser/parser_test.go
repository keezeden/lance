package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var math_ast = map[string]interface{}{
	"type": "progam",
	"body": map[string]interface{}{
		"type": "literal",
		"expression": map[string]interface{}{
			"operator": "+",
			"body": []map[string]interface{}{
				{ "type": "literal", "value": "1"  },
				{
					"expression": map[string]interface{}{
						"operator": "+",
						"body": []map[string]interface{}{
							{ "type": "literal", "value": "2"  },
							{ "type": "literal", "value": "3"  },
						},
					},
				},
			},
		},
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