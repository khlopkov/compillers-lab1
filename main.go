package main

import (
	"fmt"

	"github.com/khlopkov/compillers-lab1/parser"
)

var A = parser.NewSymbol("A", parser.NonTerminal)
var B = parser.NewSymbol("B", parser.NonTerminal)
var A1 = parser.NewSymbol("A1", parser.NonTerminal)
var B1 = parser.NewSymbol("B1", parser.NonTerminal)
var C = parser.NewSymbol("C", parser.NonTerminal)

var plus = parser.NewSymbol("+", parser.Terminal)
var minus = parser.NewSymbol("-", parser.Terminal)
var mul = parser.NewSymbol("*", parser.Terminal)
var div = parser.NewSymbol("/", parser.Terminal)

var openBracket = parser.NewSymbol("(", parser.Terminal)
var closeBracket = parser.NewSymbol(")", parser.Terminal)

var a = parser.NewSymbol("a", parser.Terminal)

func initRules() []parser.Rule {

	rules := make([]parser.Rule, 0, 10)

	rules = append(
		rules,
		parser.NewRule(A, A, plus, B1),
		parser.NewRule(A, A, minus, B1),
		parser.NewRule(A, B1),
		parser.NewRule(A1, A),

		parser.NewRule(B, B, mul, C),
		parser.NewRule(B, B, div, C),
		parser.NewRule(B, C),
		parser.NewRule(B1, B),

		parser.NewRule(C, a),
		parser.NewRule(C, openBracket, A1, closeBracket),
	)

	return rules
}

func initMatrix() parser.PrecedenceMatrix {
	return parser.NewPrecedenceMatrix(
		parser.NewRelation(A, parser.Equal, plus),
		parser.NewRelation(A, parser.Equal, minus),
		parser.NewRelation(A, parser.Greater, closeBracket),

		parser.NewRelation(B1, parser.Greater, plus),
		parser.NewRelation(B1, parser.Greater, minus),
		parser.NewRelation(B1, parser.Greater, closeBracket),

		parser.NewRelation(A1, parser.Equal, closeBracket),

		parser.NewRelation(B, parser.Greater, plus),
		parser.NewRelation(B, parser.Greater, minus),
		parser.NewRelation(B, parser.Greater, closeBracket),
		parser.NewRelation(B, parser.Equal, mul),
		parser.NewRelation(B, parser.Equal, div),

		parser.NewRelation(C, parser.Greater, plus),
		parser.NewRelation(C, parser.Greater, minus),
		parser.NewRelation(C, parser.Greater, closeBracket),
		parser.NewRelation(C, parser.Greater, mul),
		parser.NewRelation(C, parser.Greater, div),

		parser.NewRelation(plus, parser.Equal, B1),
		parser.NewRelation(plus, parser.Less, B),
		parser.NewRelation(plus, parser.Less, C),
		parser.NewRelation(plus, parser.Less, a),
		parser.NewRelation(plus, parser.Less, openBracket),

		parser.NewRelation(minus, parser.Equal, B1),
		parser.NewRelation(minus, parser.Less, B),
		parser.NewRelation(minus, parser.Less, C),
		parser.NewRelation(minus, parser.Less, a),
		parser.NewRelation(minus, parser.Less, openBracket),

		parser.NewRelation(mul, parser.Equal, C),
		parser.NewRelation(mul, parser.Less, a),
		parser.NewRelation(mul, parser.Less, openBracket),

		parser.NewRelation(div, parser.Equal, C),
		parser.NewRelation(div, parser.Less, a),
		parser.NewRelation(div, parser.Less, openBracket),

		parser.NewRelation(a, parser.Greater, plus),
		parser.NewRelation(a, parser.Greater, minus),
		parser.NewRelation(a, parser.Greater, mul),
		parser.NewRelation(a, parser.Greater, div),
		parser.NewRelation(a, parser.Greater, closeBracket),

		parser.NewRelation(openBracket, parser.Less, A),
		parser.NewRelation(openBracket, parser.Less, B1),
		parser.NewRelation(openBracket, parser.Equal, A1),
		parser.NewRelation(openBracket, parser.Less, B),
		parser.NewRelation(openBracket, parser.Less, C),
		parser.NewRelation(openBracket, parser.Less, a),
		parser.NewRelation(openBracket, parser.Less, openBracket),

		parser.NewRelation(closeBracket, parser.Greater, plus),
		parser.NewRelation(closeBracket, parser.Greater, minus),
		parser.NewRelation(closeBracket, parser.Greater, mul),
		parser.NewRelation(closeBracket, parser.Greater, div),
		parser.NewRelation(closeBracket, parser.Greater, closeBracket),
	)
}

func main() {
	grammarParser := initParser()
	fmt.Println(grammarParser.Parse("a+((a))"))
}

func initParser() parser.Parser {
	return parser.NewParser(
		parser.NewGrammar(initMatrix(), initRules()...),
	)
}
