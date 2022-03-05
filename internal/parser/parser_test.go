package parser

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/keezeden/lance/internal/lexer"
	"github.com/stretchr/testify/assert"
)

var expectedBytes, _ = os.ReadFile("./trees/math.json")

func TestParserMath(t *testing.T) {
	assert := assert.New(t)
	file := "../../snippets/math.ll"

	lexerer := lexer.BuildLexer(file)
	parser := BuildParser(lexerer)

	ast := parser.Parse()

	actualBytes, _ := json.Marshal(ast)

	// this comparison is ordering sensitive
	// TODO: make this not sensitive lol
	assert.JSONEqf(string(actualBytes), string(expectedBytes), "Parser parses 'math' correctly")
}