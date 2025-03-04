package evaluator

import "my-language/internal/parser"

func Evaluate(ast parser.ASTNode) {
	// Пример простого интерпретатора
	for _, child := range ast.Children {
		println(child.Type, ":", child.Value)
	}
}
