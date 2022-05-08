package parser

func (p *Parser) ParseOperator() Node {
	if p.lexer.Eof() {
		return nil
	}
	token := p.lexer.Peek()
	if token.Type == "op" {
		p.lexer.Pop()
		return p.BuildOperator(token)
	}

	return nil
}

func (p *Parser) ParseVar() Node {
	if p.lexer.Eof() {
		return nil
	}
	token := p.lexer.Peek()
	if token.Type == "var" || token.Type == "kw" {
		p.lexer.Pop()
		return p.BuildVar(token)
	}

	return nil
}

func (p *Parser) ParseOpenParenthesis() bool {
	if p.lexer.Eof() {
		return false
	}
	token := p.lexer.Peek()
	if token.Type == "punc" && token.Value == "(" {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p *Parser) ParseClosedParenthesis() bool {
	if p.lexer.Eof() {
		return false
	}
	token := p.lexer.Peek()
	if token.Type == "punc" && token.Value == ")" {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p *Parser) ParseEquals() bool {
	if p.lexer.Eof() {
		return false
	}
	token := p.lexer.Peek()
	if token.Type == "op" && token.Value == "=" {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p *Parser) ParseGreaterThan() bool {
	if p.lexer.Eof() {
		return false
	}
	token := p.lexer.Peek()
	if token.Type == "op" && token.Value == ">" {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p *Parser) ParseOpenBracket() bool {
	if p.lexer.Eof() {
		return false
	}
	token := p.lexer.Peek()
	if token.Type == "punc" && token.Value == "{" {
		p.lexer.Pop()
		return true
	}

	return false
}
func (p *Parser) ParseClosedBracket() bool {
	if p.lexer.Eof() {
		return false
	}
	token := p.lexer.Peek()
	if token.Type == "punc" && token.Value == "}" {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p *Parser) ParseIf() bool {
	if p.lexer.Eof() {
		return false
	}
	token := p.lexer.Peek()
	if token.Type == "kw" && token.Value == "if" {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p *Parser) ParseElse() bool {
	if p.lexer.Eof() {
		return false
	}
	token := p.lexer.Peek()
	if token.Type == "kw" && token.Value == "else" {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p *Parser) ParseWhile() bool {
	if p.lexer.Eof() {
		return false
	}
	token := p.lexer.Peek()
	if token.Type == "kw" && token.Value == "while" {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p *Parser) ParseOpenSquareBracket() bool {
	if p.lexer.Eof() {
		return false
	}
	token := p.lexer.Peek()
	if token.Type == "punc" && token.Value == "[" {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p *Parser) ParseClosedSquareBracket() bool {
	if p.lexer.Eof() {
		return false
	}
	token := p.lexer.Peek()
	if token.Type == "punc" && token.Value == "]" {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p *Parser) ParseComma() bool {
	if p.lexer.Eof() {
		return false
	}
	token := p.lexer.Peek()
	if token.Type == "punc" && token.Value == "," {
		p.lexer.Pop()
		return true
	}

	return false
}

// TODO: Remove `let` and solve loop iterators instead
func (p *Parser) ParseLet() bool {
	if p.lexer.Eof() {
		return false
	}
	token := p.lexer.Peek()
	if token.Type == "kw" && token.Value == "let" {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p *Parser) ParseConst() bool {
	if p.lexer.Eof() {
		return false
	}
	token := p.lexer.Peek()
	if token.Type == "kw" && token.Value == "const" {
		p.lexer.Pop()
		return true
	}

	return false
}

func (p *Parser) ParseLiteral() Node {
	if p.lexer.Eof() {
		return nil
	}
	token := p.lexer.Peek()
	if token.Type == "str" || token.Type == "num" {
		p.lexer.Pop()
		return p.BuildLiteral(token)
	}

	return nil
}