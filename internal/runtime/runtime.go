package runtime

import (
	"fmt"
	"my-language/internal/evaluator"
	"my-language/internal/parser"
)

// Run выполняет AST (абстрактное синтаксическое дерево)
func Run(ast parser.ASTNode) {
	fmt.Println("Running program...")
	evaluator.Evaluate(ast)
}

// HandleError обрабатывает ошибки времени выполнения
func HandleError(err error) {
	fmt.Printf("Runtime error: %s\n", err)
}
