package parser

import "github.com/keezeden/lance/internal/lexer"

func (p *Parser) BuildOperator(token lexer.Token) Node {
	return Node{
		"op": token.Value,
	}
}

func (p *Parser) BuildVar(token lexer.Token) Node {
	return Node{
		"type": "variable",
		"name": token.Value,
	}
}

func (p *Parser) BuildBracketAccessor(parentNode Node, childNode Node) Node {
	return Node{
		"type":   "variable",
		"name":   childNode,
		"parent": parentNode,
	}
}

func (p *Parser) BuildLiteral(token lexer.Token) Node {
	return Node{
		"type":  "literal",
		"value": token.Value,
	}
}

func (p *Parser) BuildExpression(rhs Node, operator Node, lhs Node) Node {
	body := []Node{lhs, rhs}
	return Node{
		"type":     "expression",
		"operator": operator["op"],
		"body":     body,
	}
}

func (p *Parser) BuildAssignment(identifier Node, value Node) Node {
	body := []Node{value}
	return Node{
		"type":       "assignment",
		"identifier": identifier["name"],
		"body":       body,
	}
}

func (p *Parser) BuildCall(identifier Node, terms Node) Node {
	arguments := []Node{terms}
	return Node{
		"type":       "call",
		"identifier": identifier["name"],
		"arguments":  arguments,
	}
}

func (p *Parser) BuildConditional(condition Node, ifStatement []Node, elseStatement []Node) Node {
	return Node{
		"type":      "conditional",
		"condition": condition,
		"if":        ifStatement,
		"else":      elseStatement,
	}
}

func (p *Parser) BuildLoop(condition Node, statements []Node) Node {
	return Node{
		"type":      "loop",
		"condition": condition,
		"body":      statements,
	}
}

func (p *Parser) BuildArray(terms []Node) Node {
	return Node{
		"type":  "array",
		"value": terms,
	}
}

func (p *Parser) BuildFunction(arguments []Node, statements []Node) Node {
	return Node{
		"type":  "function",
		"arguments": arguments,
		"body": statements,
	}
}