package translator

import (
	"testing"

	"github.com/keezeden/lance/internal/lexer"
	"github.com/keezeden/lance/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestTranslator(t *testing.T) {
	assert := assert.New(t)
	file := "../../snippets/output.ll"

	lexerer := lexer.BuildLexer(file)
	parserer := parser.BuildParser(lexerer)

	ast := parserer.Parse()

	translator := BuildTranslator(ast, "javascript")

	actualOutput := translator.Translate()
	expectedOutput := "console.log(\"Hello World\")"

	assert.Equal(expectedOutput,string(actualOutput), "Translator translates 'output' correctly")
}