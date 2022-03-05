package lexer

var OutputTokens = []Token{
	{
		Type:  "kw",
		Value: "out",
	},
	{
		Type:  "punc",
		Value: "(",
	},
	{
		Type:  "str",
		Value: "Hello World",
	},
	{
		Type:  "punc",
		Value: ")",
	},
}

var InputTokens = []Token{
	{
		Type:  "kw",
		Value: "const",
	},
	{
		Type:  "var",
		Value: "name",
	},
	{
		Type:  "op",
		Value: "=",
	},
	{
		Type:  "kw",
		Value: "in",
	},
	{
		Type:  "punc",
		Value: "(",
	},
	{
		Type:  "str",
		Value: "Whats your name?",
	},
	{
		Type:  "punc",
		Value: ")",
	},
	{
		Type:  "kw",
		Value: "out",
	},
	{
		Type:  "punc",
		Value: "(",
	},
	{
		Type:  "str",
		Value: "Hello ",
	},
	{
		Type:  "op",
		Value: "+",
	},
	{
		Type:  "var",
		Value: "name",
	},
	{
		Type:  "punc",
		Value: ")",
	},
}

var FunctionTokens = []Token{
	{
		Type:  "kw",
		Value: "const",
	},
	{
		Type:  "var",
		Value: "greet",
	},
	{
		Type:  "op",
		Value: "=",
	},
	{
		Type:  "punc",
		Value: "(",
	},
	{
		Type:  "var",
		Value: "name",
	},
	{
		Type:  "punc",
		Value: ")",
	},
	{
		Type:  "op",
		Value: "=",
	},
	{
		Type:  "op",
		Value: ">",
	},
	{
		Type:  "punc",
		Value: "{",
	},
	{
		Type:  "kw",
		Value: "out",
	},
	{
		Type:  "punc",
		Value: "(",
	},
	{
		Type:  "str",
		Value: "Hello ",
	},
	{
		Type:  "op",
		Value: "+",
	},
	{
		Type:  "var",
		Value: "name",
	},
	{
		Type:  "punc",
		Value: ")",
	},
	{
		Type:  "punc",
		Value: "}",
	},
	{
		Type:  "kw",
		Value: "const",
	},
	{
		Type:  "var",
		Value: "name",
	},
	{
		Type:  "op",
		Value: "=",
	},
	{
		Type:  "kw",
		Value: "in",
	},
	{
		Type:  "punc",
		Value: "(",
	},
	{
		Type:  "str",
		Value: "Whats your name?",
	},
	{
		Type:  "punc",
		Value: ")",
	},
	{
		Type:  "var",
		Value: "greet",
	},
	{
		Type:  "punc",
		Value: "(",
	},
	{
		Type:  "var",
		Value: "name",
	},
	{
		Type:  "punc",
		Value: ")",
	},
}

var LoopTokens = []Token{
	{
		Type:  "kw",
		Value: "const",
	},
	{
		Type:  "var",
		Value: "data",
	},
	{
		Type:  "op",
		Value: "=",
	},
	{
		Type:  "punc",
		Value: "[",
	},
	{
		Type:  "num",
		Value: 1,
	},
	{
		Type:  "punc",
		Value: ",",
	},
	{
		Type:  "num",
		Value: 2,
	},
	{
		Type:  "punc",
		Value: ",",
	},
	{
		Type:  "num",
		Value: 3,
	},
	{
		Type:  "punc",
		Value: "]",
	},
	{
		Type:  "kw",
		Value: "for",
	},
	{
		Type:  "punc",
		Value: "(",
	},
	{
		Type:  "var",
		Value: "element",
	},
	{
		Type:  "kw",
		Value: "of",
	},
	{
		Type:  "var",
		Value: "data",
	},
	{
		Type:  "punc",
		Value: ")",
	},
	{
		Type:  "punc",
		Value: "{",
	},
	{
		Type:  "kw",
		Value: "out",
	},
	{
		Type:  "punc",
		Value: "(",
	},
	{
		Type:  "var",
		Value: "element",
	},
	{
		Type:  "op",
		Value: "+",
	},
	{
		Type:  "num",
		Value: 10,
	},
	{
		Type:  "punc",
		Value: ")",
	},
	{
		Type:  "punc",
		Value: "}",
	},
}

var ConditionalTokens = []Token{
	{
		Type:  "kw",
		Value: "const",
	},
	{
		Type:  "var",
		Value: "animal",
	},
	{
		Type:  "op",
		Value: "=",
	},
	{
		Type:  "kw",
		Value: "in",
	},
	{
		Type:  "punc",
		Value: "(",
	},
	{
		Type:  "str",
		Value: "Whats your favourite animal?",
	},
	{
		Type:  "punc",
		Value: ")",
	},
	{
		Type:  "kw",
		Value: "if",
	},
	{
		Type:  "punc",
		Value: "(",
	},
	{
		Type:  "var",
		Value: "animal",
	},
	{
		Type:  "op",
		Value: "==",
	},
	{
		Type:  "str",
		Value: "frog",
	},
	{
		Type:  "punc",
		Value: ")",
	},
	{
		Type:  "punc",
		Value: "{",
	},
	{
		Type:  "kw",
		Value: "out",
	},
	{
		Type:  "punc",
		Value: "(",
	},
	{
		Type:  "str",
		Value: "A man who enjoys culture I see",
	},
	{
		Type:  "punc",
		Value: ")",
	},
	{
		Type:  "punc",
		Value: "}",
	},
	{
		Type:  "kw",
		Value: "else",
	},
	{
		Type:  "punc",
		Value: "{",
	},
	{
		Type:  "kw",
		Value: "out",
	},
	{
		Type:  "punc",
		Value: "(",
	},
	{
		Type:  "str",
		Value: "I prefer frogs",
	},
	{
		Type:  "punc",
		Value: ")",
	},
	{
		Type:  "punc",
		Value: "}",
	},
}

var MathTokens = []Token{
	{
		Type:  "num",
		Value: 1,
	},
	{
		Type:  "op",
		Value: "+",
	},
	{
		Type:  "num",
		Value: 2,
	},
	{
		Type:  "op",
		Value: "+",
	},
	{
		Type:  "num",
		Value: 3,
	},
}