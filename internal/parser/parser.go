package parser

import (
	"github.com/keezeden/lance/internal/lexer"
)

type Parser struct {
	tree map[string]interface{}
	lexer lexer.Lexer
}

type Node = map[string]interface{}


func (p* Parser) ParseOperator() Node {
	if (p.lexer.Eof()) {
		return nil
	}
	token := p.lexer.Peek()
	if (token.Type == "op") {
		p.lexer.Pop()
		return p.BuildOperator(token)
	} 

	return nil
}

func (p* Parser) ParseVar() Node {
	if (p.lexer.Eof()) {
		return nil
	}
	token := p.lexer.Peek()
	if (token.Type == "var" || token.Type == "kw") {
		p.lexer.Pop()
		return p.BuildVar(token)
	}

	return nil
}

func (p* Parser) ParseOpenParenthesis() bool {
	if (p.lexer.Eof()) {
		return false
	}
	token := p.lexer.Peek()
	if (token.Type == "punc" && token.Value == "(") {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p* Parser) ParseClosedParenthesis() bool {
	if (p.lexer.Eof()) {
		return false
	}
	token := p.lexer.Peek()
	if (token.Type == "punc" && token.Value == ")") {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p* Parser) ParseLiteral() Node {
	if (p.lexer.Eof()) {
		return nil
	}
	token := p.lexer.Peek()
	if (token.Type == "str" || token.Type == "num") {
		p.lexer.Pop()
		return p.BuildLiteral(token)
	}

	return nil
}

func (p* Parser) BuildOperator(token lexer.Token) Node {
	return Node{
		"op": token.Value,
	}
}

func (p* Parser) BuildVar(token lexer.Token) Node {
	return Node{
		"type": "variable",
		"name": token.Value,
	}
}

func (p* Parser) BuildLiteral(token lexer.Token) Node {
	return Node{
		"type": "literal",
		"value": token.Value,
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

func (p* Parser) BuildCall(identifier Node, terms Node) Node {
	arguments := []Node{terms}
	return Node{
		"type": "call",
		"identifier": identifier["name"],
		"arguments": arguments,
	  }
}

func (p* Parser) ParseTerms() Node {
	sideNode := p.ParseVar()
	if (sideNode != nil) {
		return sideNode
	}	
	sideNode = p.ParseLiteral()
	if (sideNode != nil) {
		return sideNode
	}

	return nil
}

func (p* Parser) ParseCall() Node {
	varNode := p.ParseVar()
	if (varNode == nil) {
		return nil
	}
	
	isOpenParenthesis := p.ParseOpenParenthesis()
	if (!isOpenParenthesis) {
		return nil
	}

	termNode := p.ParseTerms()
	if (termNode == nil) {
		return nil
	}

	isClosedParenthesis := p.ParseClosedParenthesis()
	if (!isClosedParenthesis) {
		return nil
	}

	return p.BuildCall(varNode, termNode)
}

func (p* Parser) ParseAppendedTerms() (Node, Node) {
	operatorNode := p.ParseOperator()
	if (operatorNode == nil) {
		return nil, nil
	}

	termsNode := p.ParseTerms()
	if (termsNode == nil) {
		return operatorNode, nil
	}

	appendedOperatorTerms, appendedTermTerms := p.ParseAppendedTerms()
	if (appendedOperatorTerms != nil && appendedTermTerms != nil) {
		return operatorNode, p.BuildExpression(termsNode, appendedOperatorTerms, appendedTermTerms)
	}

	return operatorNode, termsNode
}


func (p* Parser) ParseExpression() Node {
	termsNode := p.ParseTerms()
	if (termsNode == nil) {
		return nil
	}

	operatorNode, extraTermsNode := p.ParseAppendedTerms()
	if (operatorNode == nil || extraTermsNode == nil) {
		return nil
	}

	return p.BuildExpression(termsNode, operatorNode, extraTermsNode)
}

func (p* Parser) ParseStatement() Node {
	callNode := p.ParseCall()
	if (callNode != nil) {
		return callNode
	}

	expressionNode := p.ParseExpression()
	if (expressionNode != nil) {
		return expressionNode
	}

	return nil
}

// expression = literal | expression
func (p* Parser) Parse() Node {	
	treeNode := p.ParseStatement()

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