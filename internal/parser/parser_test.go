package parser

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/keezeden/lance/internal/lexer"
	"github.com/stretchr/testify/assert"
)


func TestParserMath(t *testing.T) {
	assert := assert.New(t)
	file := "../../snippets/math.ll"

	lexerer := lexer.BuildLexer(file)
	parser := BuildParser(lexerer)

	ast := parser.Parse()

	actualBytes, _ := json.Marshal(ast)
	expectedBytes, _ := os.ReadFile("./trees/math.json")

	// this comparison is ordering sensitive
	// TODO: make this not sensitive lol
	assert.JSONEqf(string(actualBytes), string(expectedBytes), "Parser parses 'math' correctly")
}

func TestParserOutput(t *testing.T) {
	assert := assert.New(t)
	file := "../../snippets/output.ll"

	lexerer := lexer.BuildLexer(file)
	parser := BuildParser(lexerer)

	ast := parser.Parse()

	actualBytes, _ := json.Marshal(ast)
	expectedBytes, _ := os.ReadFile("./trees/output.json")

	assert.JSONEqf(string(actualBytes), string(expectedBytes), "Parser parses 'output' correctly")
}

func TestParserInput(t *testing.T) {
	assert := assert.New(t)
	file := "../../snippets/input.ll"

	lexerer := lexer.BuildLexer(file)
	parser := BuildParser(lexerer)

	ast := parser.Parse()

	actualBytes, _ := json.Marshal(ast)
	expectedBytes, _ := os.ReadFile("./trees/input.json")

	assert.JSONEqf(string(actualBytes), string(expectedBytes), "Parser parses 'input' correctly")
}

func TestConditionalInput(t *testing.T) {
	assert := assert.New(t)
	file := "../../snippets/conditional.ll"

	lexerer := lexer.BuildLexer(file)
	parser := BuildParser(lexerer)

	ast := parser.Parse()

	actualBytes, _ := json.Marshal(ast)
	expectedBytes, _ := os.ReadFile("./trees/conditional.json")

	assert.JSONEqf(string(actualBytes), string(expectedBytes), "Parser parses 'input' correctly")
}

func TestLoopInput(t *testing.T) {
	assert := assert.New(t)
	file := "../../snippets/loop.ll"

	lexerer := lexer.BuildLexer(file)
	parser := BuildParser(lexerer)

	ast := parser.Parse()

	actualBytes, _ := json.Marshal(ast)
	expectedBytes, _ := os.ReadFile("./trees/loop.json")

	assert.JSONEqf(string(actualBytes), string(expectedBytes), "Parser parses 'loop' correctly")
}

