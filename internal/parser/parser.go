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

func (p* Parser) ParseWhile() bool {
	if (p.lexer.Eof()) {
		return false
	}
	token := p.lexer.Peek()
	if (token.Type == "kw" && token.Value == "while") {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p* Parser) ParseOpenSquareBracket() bool {
	if (p.lexer.Eof()) {
		return false
	}
	token := p.lexer.Peek()
	if (token.Type == "punc" && token.Value == "[") {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p* Parser) ParseClosedSquareBracket() bool {
	if (p.lexer.Eof()) {
		return false
	}
	token := p.lexer.Peek()
	if (token.Type == "punc" && token.Value == "]") {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p* Parser) ParseComma() bool {
	if (p.lexer.Eof()) {
		return false
	}
	token := p.lexer.Peek()
	if (token.Type == "punc" && token.Value == ",") {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p* Parser) ParseDot() bool {
	if (p.lexer.Eof()) {
		return false
	}
	token := p.lexer.Peek()
	if (token.Type == "punc" && token.Value == ".") {
		p.lexer.Pop()
		return true
	}

	return false
}
// TODO: Remove `let` and solve loop iterators instead
func (p* Parser) ParseLet() bool {
	if (p.lexer.Eof()) {
		return false
	}
	token := p.lexer.Peek()
	if (token.Type == "kw" && token.Value == "let") {
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


func (p* Parser) BuildBracketAccessor(parentNode Node, childNode Node) Node {
	return Node{
		"type": "variable",
		"name": childNode,
		"parent": parentNode,
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
		"condition": condition,
		"if": ifStatement,
		"else": elseStatement,
	  }
}

func (p* Parser) BuildLoop(condition Node, statements []Node) Node {
	return Node{
		"type": "loop",
		"condition": condition,
		"body": statements,		
	  }
}

func (p* Parser) BuildArray(terms []Node) Node {
	return Node{
		"type": "array",
		"value": terms,
	  }
}

func (p* Parser) ParseTerms() Node {
	sideNode := p.ParseAccessor()
	if (sideNode != nil) {
		return sideNode
	}
	
	sideNode = p.ParseLiteral()
	if (sideNode != nil) {
		return sideNode
	}

	sideNode = p.ParseVar()
	if (sideNode != nil) {
		return sideNode
	}

	sideNode = p.ParseArray()
	if (sideNode != nil) {
		return sideNode
	}

	return nil
}

func (p* Parser) ParseAccessor() Node {
	node := p.ParseBracketAccessor()
	if (node != nil) {
		return node
	}

	return nil
}


func (p* Parser) ParseBracketAccessor() Node {
	var parentNode Node
	parentNode = p.ParseVar()
	
	if (parentNode == nil)	{
		return nil
	}

	isOpenSquareBracket := p.ParseOpenSquareBracket()
	if (!isOpenSquareBracket) {
		return parentNode
	}

	childNode := p.ParseVar()
	if (childNode == nil)	{
		return nil
	}

	isClosedSquareBracket := p.ParseClosedSquareBracket()
	if (!isClosedSquareBracket) {
		return nil
	}

	return p.BuildBracketAccessor(parentNode, childNode)
}

func (p* Parser) ParseArrayTerms(terms []Node) []Node {
	termNode := p.ParseTerms()
	if (termNode == nil) {
		return nil
	}

	terms = append(terms, termNode)

	isComma := p.ParseComma()
	if (!isComma) {
		return terms
	}

	return p.ParseArrayTerms(terms)
}

func (p* Parser) ParseArray() Node {
	isOpenSquareBracket := p.ParseOpenSquareBracket()
	if (!isOpenSquareBracket) {
		return nil
	}

	var termNodes []Node
	termNodes = p.ParseArrayTerms(termNodes)

	isClosedSquareBracket := p.ParseClosedSquareBracket()
	if (!isClosedSquareBracket) {
		return nil
	}

	return p.BuildArray(termNodes)
}

func (p* Parser) ParseCall() Node {
	varNode := p.ParseVar()
	if (varNode == nil) {
		return nil
	}
	
	isOpenParenthesis := p.ParseOpenParenthesis()
	if (!isOpenParenthesis) {
		return varNode
	}

	expressionNode := p.ParseExpression(nil)
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


func (p* Parser) ParseExpression(preTermNode Node) Node {
	termsNode := preTermNode
	if (preTermNode == nil) {
		termsNode = p.ParseTerms()
		if (termsNode == nil) {
			return nil
		}
}

	operatorNode, extraTermsNode := p.ParseAppendedTerms()
	if (operatorNode == nil || extraTermsNode == nil) {
		return termsNode
	}

	return p.BuildExpression(termsNode, operatorNode, extraTermsNode)
}

func (p* Parser) ParseAssignment() Node {
	isConst := p.ParseConst()
	isLet := p.ParseLet()
	
	if (!isConst && !isLet) {
		return nil
	}
	
	varNode := p.ParseVar()
	if (varNode == nil) {
		return nil
	}

	isEquals := p.ParseEquals()	
	if (!isEquals) {
		return varNode
	}

	callNode := p.ParseCall()
	if (callNode != nil) {
		return p.BuildAssignment(varNode, callNode)
	}

	expressionNode := p.ParseExpression(nil)
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

	ifExpressionNode := p.ParseExpression(nil)
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

func (p* Parser) ParseLoop() Node {
	isWhile := p.ParseWhile()
	if (!isWhile) {
		return nil
	}

	isOpenParenthesis := p.ParseOpenParenthesis()
	if (!isOpenParenthesis) {
		return nil
	}

	expressionNode := p.ParseExpression(nil)
	if (expressionNode == nil) {
		return nil
	}

	isClosedParenthesis := p.ParseClosedParenthesis()
	if (!isClosedParenthesis) {
		return nil
	}
	
	isOpenBracket := p.ParseOpenBracket()
	if (!isOpenBracket) {
		return nil
	}
	
	var statementNodes []Node

	for (!p.ParseClosedBracket()) {
		statementNode := p.ParseStatement()	
		if (statementNode != nil) {
			statementNodes = append(statementNodes, statementNode)
		}
	}



	return p.BuildLoop(expressionNode, statementNodes)
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

	loopNode := p.ParseLoop()
	if (loopNode != nil) {
		return loopNode
	}

	callNode := p.ParseCall()

	expressionNode := p.ParseExpression(callNode)
	if (expressionNode != nil) {
		return expressionNode
	}

	return callNode
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