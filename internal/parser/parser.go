package parser

import (
	"github.com/keezeden/lance/internal/lexer"
)

type Parser struct {
	tree map[string]interface{}
	lexer lexer.Lexer
}

type Node = map[string]interface{}


func (p* Parser) Parse() Node {	
	var treeNodes []Node
	for (!p.lexer.Eof()) {
		statementNode := p.ParseStatement()
		if (statementNode != nil) {			
			treeNodes = append(treeNodes, statementNode)
		}
	}

	return Node{
		"type": "program",
		"body": treeNodes,
	}
}

func BuildParser(lexer lexer.Lexer) Parser {
	var tree = map[string]interface{}{}
	return Parser{
		tree: tree,
		lexer: lexer,
	}
}