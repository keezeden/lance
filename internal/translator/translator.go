package translator

import "github.com/keezeden/lance/internal/parser"

type Translator struct {
	ast parser.Node
	output string
}

func (t* Translator) Translate() []byte {
	outputFormats := make(map[string]func(parser.Node) []byte)

	outputFormats["javascript"] = TranslateToJavaScript

	return outputFormats[t.output](t.ast)
}

func BuildTranslator(ast parser.Node, output string) Translator {
	return Translator{
		ast: ast,
		output: output,
	}
}