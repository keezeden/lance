package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/keezeden/lance/internal/lexer"
	"github.com/stretchr/testify/assert"
)

var MathAST map[string]interface{}
var content, _ = os.ReadFile("./trees/math.json")

var _ = json.Unmarshal(content, &MathAST)


func TestParserMath(t *testing.T) {
	assert := assert.New(t)
	file := "../../snippets/math.ll"

	lexerer := lexer.BuildLexer(file)
	parser := BuildParser(lexerer)

	ast := parser.Parse()

	js, err := json.Marshal(ast)
if err != nil {
    log.Fatal("EROROROROROROR", err)
}
fmt.Println(string(js))

	assert.Equal(ast, MathAST, "Parser parses 'math' correctly")
}