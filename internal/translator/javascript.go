package translator

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/keezeden/lance/internal/parser"
)

func TranslateLiteral(ast parser.Node) []byte {
	value := ast["value"]
	if (reflect.ValueOf(value).Kind() == reflect.String) {
		return []byte(fmt.Sprintf("\"%s\"", value))
	} else {
		return []byte(strconv.Itoa(value.(int)))
	}
	
}

func TranslateCall(ast parser.Node) []byte {
	if (ast["identifier"] == "out") {
		args := ast["arguments"].([]parser.Node)
		// this will need to recursivly solve the args (func(1 + 2, 3 + 4))
		return []byte(fmt.Sprintf("console.log(%s)", TranslateToJavaScript(args[0])))
	}
	
	return []byte("lol")
}

func TranslateToJavaScript(ast parser.Node) []byte {
	if (ast["body"] == nil) {
		switch ast["type"] {
		case "literal":
			return TranslateLiteral(ast)
		}
	}
	for _, statement := range ast["body"].([]parser.Node) {
		switch statement["type"] {
		case "call":
			return TranslateCall(statement)
		}
	}

	return nil
}