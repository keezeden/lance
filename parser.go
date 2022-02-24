package main

type Parser struct {
	tree interface{}
	lexer Lexer
}

func (p* Parser) parse() {
	for !leof(&p.lexer) {

		token := lpeek(&p.lexer)

		// TODO: move these to an enum
		switch token.category {
		case "":
			
		}
	}
}

func parser(lexer Lexer) Parser {
	tree := []string{}
	return Parser{
		tree: tree,
		lexer: lexer,
	}
}