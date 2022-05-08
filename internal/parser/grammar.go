package parser

func (p *Parser) ParseTerms() Node {
	sideNode := p.ParseAccessor()
	if sideNode != nil {
		return sideNode
	}

	sideNode = p.ParseLiteral()
	if sideNode != nil {
		return sideNode
	}

	sideNode = p.ParseVar()
	if sideNode != nil {
		return sideNode
	}

	sideNode = p.ParseArray()
	if sideNode != nil {
		return sideNode
	}

	sideNode = p.ParseFunction()
	if sideNode != nil {
		return sideNode
	}

	return nil
}

func (p *Parser) ParseAccessor() Node {
	node := p.ParseBracketAccessor()
	if node != nil {
		return node
	}

	return nil
}

func (p *Parser) ParseBracketAccessor() Node {
	var parentNode Node
	parentNode = p.ParseVar()

	if parentNode == nil {
		return nil
	}

	isOpenSquareBracket := p.ParseOpenSquareBracket()
	if !isOpenSquareBracket {
		return parentNode
	}

	childNode := p.ParseVar()
	if childNode == nil {
		return nil
	}

	isClosedSquareBracket := p.ParseClosedSquareBracket()
	if !isClosedSquareBracket {
		return nil
	}

	return p.BuildBracketAccessor(parentNode, childNode)
}

func (p *Parser) ParseArrayTerms(terms []Node) []Node {
	termNode := p.ParseTerms()
	if termNode == nil {
		return nil
	}

	terms = append(terms, termNode)

	isComma := p.ParseComma()
	if !isComma {
		return terms
	}

	return p.ParseArrayTerms(terms)
}

func (p *Parser) ParseArray() Node {
	isOpenSquareBracket := p.ParseOpenSquareBracket()
	if !isOpenSquareBracket {
		return nil
	}

	var termNodes []Node
	termNodes = p.ParseArrayTerms(termNodes)

	isClosedSquareBracket := p.ParseClosedSquareBracket()
	if !isClosedSquareBracket {
		return nil
	}

	return p.BuildArray(termNodes)
}

func (p *Parser) ParseCall() Node {
	varNode := p.ParseVar()
	if varNode == nil {
		return nil
	}

	isOpenParenthesis := p.ParseOpenParenthesis()
	if !isOpenParenthesis {
		return varNode
	}

	expressionNode := p.ParseExpression(nil)
	if expressionNode == nil {
		return nil
	}

	isClosedParenthesis := p.ParseClosedParenthesis()
	if !isClosedParenthesis {
		return nil
	}

	return p.BuildCall(varNode, expressionNode)
}

func (p *Parser) ParseAppendedTerms() (Node, Node) {
	operatorNode := p.ParseOperator()
	if operatorNode == nil {
		return nil, nil
	}

	termsNode := p.ParseTerms()
	if termsNode == nil {
		return operatorNode, nil
	}

	appendedOperatorTerms, appendedTermTerms := p.ParseAppendedTerms()
	if appendedOperatorTerms != nil && appendedTermTerms != nil {
		return operatorNode, p.BuildExpression(termsNode, appendedOperatorTerms, appendedTermTerms)
	}

	return operatorNode, termsNode
}

func (p *Parser) ParseExpression(preTermNode Node) Node {
	termsNode := preTermNode
	if preTermNode == nil {
		termsNode = p.ParseTerms()
		if termsNode == nil {
			return nil
		}
	}

	operatorNode, extraTermsNode := p.ParseAppendedTerms()
	if operatorNode == nil || extraTermsNode == nil {
		return termsNode
	}

	return p.BuildExpression(termsNode, operatorNode, extraTermsNode)
}

func (p *Parser) ParseAssignment() Node {
	isConst := p.ParseConst()
	isLet := p.ParseLet()

	if !isConst && !isLet {
		return nil
	}

	varNode := p.ParseVar()
	if varNode == nil {
		return nil
	}

	isEquals := p.ParseEquals()
	if !isEquals {
		return varNode
	}

	callNode := p.ParseCall()
	if callNode != nil {
		return p.BuildAssignment(varNode, callNode)
	}

	expressionNode := p.ParseExpression(nil)
	if expressionNode != nil {
		return p.BuildAssignment(varNode, expressionNode)
	}

	termsNode := p.ParseTerms()
	if termsNode != nil {
		return p.BuildAssignment(varNode, termsNode)
	}

	return nil
}

func (p *Parser) ParseConditional() Node {
	isIf := p.ParseIf()
	if !isIf {
		return nil
	}

	isIfOpenParenthesis := p.ParseOpenParenthesis()
	if !isIfOpenParenthesis {
		return nil
	}

	ifExpressionNode := p.ParseExpression(nil)
	if ifExpressionNode == nil {
		return nil
	}

	isIfClosedParenthesis := p.ParseClosedParenthesis()
	if !isIfClosedParenthesis {
		return nil
	}

	isIfOpenBracket := p.ParseOpenBracket()
	if !isIfOpenBracket {
		return nil
	}

	var ifStatementNodes []Node
	for !p.ParseClosedBracket() {
		statementNode := p.ParseStatement()
		if statementNode != nil {
			ifStatementNodes = append(ifStatementNodes, statementNode)
		}
	}

	isElse := p.ParseElse()
	if !isElse {
		return nil
	}

	isElseOpenBracket := p.ParseOpenBracket()
	if !isElseOpenBracket {
		return nil
	}

	var elseStatementNodes []Node
	for !p.ParseClosedBracket() {
		statementNode := p.ParseStatement()
		if statementNode != nil {
			elseStatementNodes = append(elseStatementNodes, statementNode)
		}
	}

	return p.BuildConditional(ifExpressionNode, ifStatementNodes, elseStatementNodes)
}

func (p *Parser) ParseLoop() Node {
	isWhile := p.ParseWhile()
	if !isWhile {
		return nil
	}

	isOpenParenthesis := p.ParseOpenParenthesis()
	if !isOpenParenthesis {
		return nil
	}

	expressionNode := p.ParseExpression(nil)
	if expressionNode == nil {
		return nil
	}

	isClosedParenthesis := p.ParseClosedParenthesis()
	if !isClosedParenthesis {
		return nil
	}

	isOpenBracket := p.ParseOpenBracket()
	if !isOpenBracket {
		return nil
	}

	var statementNodes []Node

	for !p.ParseClosedBracket() {
		statementNode := p.ParseStatement()
		if statementNode != nil {
			statementNodes = append(statementNodes, statementNode)
		}
	}

	return p.BuildLoop(expressionNode, statementNodes)
}

func (p *Parser) ParseStatement() Node {
	assignmentNode := p.ParseAssignment()
	if assignmentNode != nil {
		return assignmentNode
	}

	conditionalNode := p.ParseConditional()
	if conditionalNode != nil {
		return conditionalNode
	}

	loopNode := p.ParseLoop()
	if loopNode != nil {
		return loopNode
	}

	callNode := p.ParseCall()

	expressionNode := p.ParseExpression(callNode)
	if expressionNode != nil {
		return expressionNode
	}

	return callNode
}

func (p *Parser) ParseFunction() Node {
	isOpenParenthesis := p.ParseOpenParenthesis()
	if !isOpenParenthesis {
		return nil
	}

	var argumentNodes []Node
	argumentNodes = p.ParseArrayTerms(argumentNodes)

	isClosedParenthesis := p.ParseClosedParenthesis()
	if !isClosedParenthesis {
		return nil
	}

	isEquals := p.ParseEquals()
	if !isEquals {
		return nil
	}

	isGreaterThan := p.ParseGreaterThan()
	if !isGreaterThan {
		return nil
	}

	isOpenBracket := p.ParseOpenBracket()
	if !isOpenBracket {
		return nil
	}

	var statementNodes []Node
	for !p.ParseClosedBracket() {
		statementNode := p.ParseStatement()
		if statementNode != nil {
			statementNodes = append(statementNodes, statementNode)
		}
	}

	return p.BuildFunction(argumentNodes, statementNodes)
}