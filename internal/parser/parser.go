package parser

import (
	"fmt"

	"github.com/keezeden/lance/internal/lexer"
)

type Parser struct {
	tree map[string]interface{}
	lexer lexer.Lexer
}

type Node = map[string]interface{}


func (p* Parser) ParseOperator() Node {
	token := p.lexer.Peek()
	if (token.Type == "op") {
		p.lexer.Pop()
		return p.BuildOperator(token.Value)
	} 

	return nil
}

func (p* Parser) ParseVar() Node {
	token := p.lexer.Peek()
	if (token.Type == "var") {
		p.lexer.Pop()
		return p.BuildVar(token.Value)
	}

	return nil
}

func (p* Parser) ParseLiteral() Node {
	token := p.lexer.Peek()
	if (token.Type == "string" || token.Type == "num") {
		p.lexer.Pop()
		return p.BuildLiteral(token.Value)
	}

	return nil
}

func (p* Parser) BuildOperator(operator interface{}) Node {
	return Node{
		"op": operator,
	}
}

func (p* Parser) BuildVar(variable interface{}) Node {
	return Node{}
}

func (p* Parser) BuildLiteral(literal interface{}) Node {
	return Node{
		"literal": literal,
	}
}


func (p* Parser) BuildExpression(rhs Node, operator Node, lhs Node) Node {
	body := []Node{lhs, rhs}
	return Node{
		"type": "expression",
		"operator": operator["op"],
		"body": body,
	  }
}

func (p* Parser) ParseSide() Node {
	sideNode := p.ParseExpression()
	if (sideNode != nil) {
		return sideNode
	}

	sideNode = p.ParseVar()
	if (sideNode != nil) {
		return sideNode
	}	

	sideNode = p.ParseLiteral()
	if (sideNode != nil) {
		return sideNode
	}
	
	return nil
}

func (p* Parser) ParseExpression() Node {
	rhsNode := p.ParseSide()
	if (rhsNode == nil) {
		return nil
	}

	operatorNode := p.ParseOperator()
	fmt.Println("OP NODE", operatorNode)
	if (operatorNode == nil) {
		return operatorNode
	}

	lhsNode := p.ParseSide()
	fmt.Println("LHS NODE", lhsNode)
	if (lhsNode == nil) {
		return nil
	}

	return p.BuildExpression(rhsNode, operatorNode, lhsNode)
}

// expression = literal | expression
func (p* Parser) Parse() Node {	
	treeNode := p.ParseExpression()

	if (treeNode == nil) {
		return nil
	}

	bodyNode := []Node{treeNode}
	return Node{
		"type": "program",
		"body": bodyNode,
	}
}

func BuildParser(lexer lexer.Lexer) Parser {
	var tree = map[string]interface{}{}
	return Parser{
		tree: tree,
		lexer: lexer,
	}
}