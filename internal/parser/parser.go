package parser

import "github.com/keezeden/lance/internal/lexer"

type Parser struct {
	tree interface{}
	lexer lexer.Lexer
}

func (p* Parser) Parse() map[string]interface{} {
	// for !leof(&p.lexer) {

	// 	token := lpeek(&p.lexer)

	// 	// TODO: move these to an enum
	// 	switch token.category {
	// 	case "":

	// 	}
	// }

	return map[string]interface{}{}
}

func BuildParser(lexer lexer.Lexer) Parser {
	tree := []string{}
	return Parser{
		tree: tree,
		lexer: lexer,
	}
}