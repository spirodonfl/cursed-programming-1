package main

import (
	"fmt"
	"os"
	"time"

	"github.com/shreyassanthu77/cisp/interpreter"
	"github.com/shreyassanthu77/cisp/lexer"
	"github.com/shreyassanthu77/cisp/parser"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: crap <input>")
		return
	}

	for _, arg := range args {
		file, err := os.ReadFile(arg)
		if err != nil {
			fmt.Printf("Error loading file: %s\n", err)
			continue
		}
		input := string(file)
		fmt.Println(">> Executing:", arg)
		fmt.Println("-------------------------")

		lex := lexer.New(input)
		par := parser.New(lex)

		ast, err := par.Parse()
		if err != nil {
			fmt.Println(err)
			return
		}

		t := time.Now()
		res, err := interpreter.Eval(ast)
		done := time.Since(t)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("-------------------------")
		fmt.Printf("Main Returned %v in: %v\n\n", res, done)
	}
}
