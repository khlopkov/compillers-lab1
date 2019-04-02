package parser

type Grammar struct {
	matrix PrecedenceMatrix
	rules  []Rule
}

func NewGrammar(matrix PrecedenceMatrix, rules ...Rule) *Grammar {
	return &Grammar{
		matrix: matrix,
		rules:  rules,
	}
}
