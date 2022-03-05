package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name 		   string
    Description    string    
    Input          string 
    ExpectedOutput []Token
}

var tests = []TestCase{
	{
		Name: "Output Snippet",
		Description: "Lexer tokenizes 'output' correctly",
		Input: "../../snippets/output.ll",
		ExpectedOutput: OutputTokens,
	},
	{
		Name: "Input Snippet",
		Description: "Lexer tokenizes 'input' correctly",
		Input: "../../snippets/input.ll",
		ExpectedOutput: InputTokens, 
	},
	{
		Name: "Loop Snippet",
		Description: "Lexer tokenizes 'loop' correctly",
		Input: "../../snippets/loop.ll",
		ExpectedOutput: LoopTokens, 
	},
	{
		Name: "Function Snippet",
		Description: "Lexer tokenizes 'function' correctly",
		Input: "../../snippets/function.ll",
		ExpectedOutput: FunctionTokens, 
	},
	{
		Name: "Conditional Snippet",
		Description: "Lexer tokenizes 'conditional' correctly",
		Input: "../../snippets/conditional.ll",
		ExpectedOutput: ConditionalTokens, 
	},
	{
		Name: "Math Snippet",
		Description: "Lexer tokenizes 'math' correctly",
		Input: "../../snippets/math.ll",
		ExpectedOutput: MathTokens, 
	},
}


func TestLexer(t *testing.T) {
	for _, testCase := range tests {
		t.Run(testCase.Name, func(t *testing.T) {
			assert := assert.New(t)
			
			var tokens []Token
			lexerer := BuildLexer(testCase.Input)
			for !lexerer.Eof() {
				token := lexerer.Peek()
				
				tokens = append(tokens, token)
		
				lexerer.Pop()
			}
		
			assert.Equal(testCase.ExpectedOutput, tokens, testCase.Description)
		})
	}
}
