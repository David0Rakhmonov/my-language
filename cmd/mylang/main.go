package main

import (
	"fmt"
	"mylang/internal/evaluator"
	"mylang/internal/lexer"
	"mylang/internal/parser"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mylang <filename>")
		return
	}

	filename := os.Args[1]
	code, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}

	tokens := lexer.Tokenize(string(code))
	ast := parser.Parse(tokens)
	evaluator.Evaluate(ast)
}
