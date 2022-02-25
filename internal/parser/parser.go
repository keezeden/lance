package parser

type Parser struct {
	tree interface{}
	lexer Lexer
}

func (p* Parser) parse() map[string]interface{} {
	// for !leof(&p.lexer) {

	// 	token := lpeek(&p.lexer)

	// 	// TODO: move these to an enum
	// 	switch token.category {
	// 	case "":

	// 	}
	// }

	return map[string]interface{}{}
}

func parser(lexer Lexer) Parser {
	tree := []string{}
	return Parser{
		tree: tree,
		lexer: lexer,
	}
}