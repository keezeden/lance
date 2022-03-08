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

func (p* Parser) ParseEquals() bool {
	if (p.lexer.Eof()) {
		return false
	}
	token := p.lexer.Peek()
	if (token.Type == "op" && token.Value == "=") {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p* Parser) ParseOpenBracket() bool {
	if (p.lexer.Eof()) {
		return false
	}
	token := p.lexer.Peek()
	if (token.Type == "punc" && token.Value == "{") {
		p.lexer.Pop()
		return true
	}

	return false
}
func (p* Parser) ParseClosedBracket() bool {
	if (p.lexer.Eof()) {
		return false
	}
	token := p.lexer.Peek()
	if (token.Type == "punc" && token.Value == "}") {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p* Parser) ParseIf() bool {
	if (p.lexer.Eof()) {
		return false
	}
	token := p.lexer.Peek()
	if (token.Type == "kw" && token.Value == "if") {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p* Parser) ParseElse() bool {
	if (p.lexer.Eof()) {
		return false
	}
	token := p.lexer.Peek()
	if (token.Type == "kw" && token.Value == "else") {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p* Parser) ParseConst() bool {
	if (p.lexer.Eof()) {
		return false
	}
	token := p.lexer.Peek()
	if (token.Type == "kw" && token.Value == "const") {
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

func (p* Parser) BuildAssignment(identifier Node, value Node) Node {
	body := []Node{value}
	return Node{
		"type": "assignment",
		"identifier": identifier["name"],
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

func (p* Parser) BuildConditional(condition Node, ifStatement Node, elseStatement Node) Node {
	return Node{
		"type": "conditional",
		"body": condition,
		"if": ifStatement,
		"else": elseStatement,
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

	expressionNode := p.ParseExpression()
	if (expressionNode == nil) {
		return nil
	}

	isClosedParenthesis := p.ParseClosedParenthesis()
	if (!isClosedParenthesis) {
		return nil
	}

	return p.BuildCall(varNode, expressionNode)
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
		return termsNode
	}

	return p.BuildExpression(termsNode, operatorNode, extraTermsNode)
}

func (p* Parser) ParseAssignment() Node {
	isConst := p.ParseConst()
	if (!isConst) {
		return nil
	}

	varNode := p.ParseVar()
	if (varNode == nil) {
		return nil
	}

	isEquals := p.ParseEquals()	
	if (!isEquals) {
		return nil
	}

	callNode := p.ParseCall()
	if (callNode != nil) {
		return p.BuildAssignment(varNode, callNode)
	}

	expressionNode := p.ParseExpression()
	if (expressionNode != nil) {
		return p.BuildAssignment(varNode, expressionNode)
	}

	termsNode := p.ParseTerms()
	if (termsNode != nil) {
		return p.BuildAssignment(varNode, termsNode)
	}

	return nil
}

func (p* Parser) ParseConditional() Node {
	isIf := p.ParseIf()
	if (!isIf) {
		return nil
	}

	isIfOpenParenthesis := p.ParseOpenParenthesis()
	if (!isIfOpenParenthesis) {
		return nil
	}

	ifExpressionNode := p.ParseExpression()
	if (ifExpressionNode == nil) {
		return nil
	}

	isIfClosedParenthesis := p.ParseClosedParenthesis()
	if (!isIfClosedParenthesis) {
		return nil
	}

	isIfOpenBracket := p.ParseOpenBracket()
	if (!isIfOpenBracket) {
		return nil
	}

	ifStatementNode := p.ParseStatement()
	if (ifStatementNode == nil) {
		return nil
	}
	
	isClosedBracket := p.ParseClosedBracket()
	if (!isClosedBracket) {
		return nil
	}

	isElse := p.ParseElse()
	if (!isElse) {
		return nil
	}

	isElseOpenBracket := p.ParseOpenBracket()
	if (!isElseOpenBracket) {
		return nil
	}

	elseStatementNode := p.ParseStatement()
	if (elseStatementNode == nil) {
		return nil
	}
	
	isElseClosedBracket := p.ParseClosedBracket()
	if (!isElseClosedBracket) {
		return nil
	}

	return p.BuildConditional(ifExpressionNode, ifStatementNode, elseStatementNode)
}

func (p* Parser) ParseStatement() Node {
	assignmentNode := p.ParseAssignment()
	if (assignmentNode != nil) {
		return assignmentNode
	}

	conditionalNode := p.ParseConditional()
	if (conditionalNode != nil) {
		return conditionalNode
	}

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