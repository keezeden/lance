package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/keezeden/lance/internal/lexer"
	"github.com/keezeden/lance/internal/parser"
	"github.com/urfave/cli"
)


func main() {
	app := cli.NewApp()
	app.Name = "lance"
	app.Usage = "compile .ll files"
	app.Action = func(c *cli.Context) error {
	  filepath := c.Args().Get(0)

	  if (filepath == "") {
		fmt.Println("No filepath provided")
		return nil
	  }

	  lexerer := lexer.BuildLexer(filepath)
	  parserer := parser.BuildParser(lexerer)

	  ast := parserer.Parse()
	  bytes, _ := json.Marshal(ast)

	  err := os.WriteFile("../output.json", bytes, 0644)
	  if err != nil {
		log.Fatal(err)
	  }

	  return nil
	}
  
	err := app.Run(os.Args)
	if err != nil {
	  log.Fatal(err)
	}
}