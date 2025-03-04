package parser

import "my-language/internal/lexer"

type ASTNode struct {
	Type     string
	Value    string
	Children []ASTNode
}

func Parse(tokens []lexer.Token) ASTNode {
	// Пример простого парсера
	root := ASTNode{Type: "Program", Value: ""}
	for _, token := range tokens {
		root.Children = append(root.Children, ASTNode{
			Type:  string(token.Type),
			Value: token.Value,
		})
	}
	return root
}
