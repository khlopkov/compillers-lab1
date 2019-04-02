package main

import "testing"

var testPairs = []struct {
	input    string
	expected bool
}{
	{"a", true},
	{"a+a", true},
	{"a-a", true},
	{"a*a", true},
	{"a/a", true},
	{"(a)", true},
	{"a+(a)", true},
	{"a-(a)", true},
	{"a*(a)", true},
	{"a/(a)", true},
	{"a+((a))", true},
	{"a+(a/a)", true},
	{"a+(a*a)", true},
	{"a+(a-a)", true},
	{"a+(a+a)", true},
	{"a++", false},
	{"a+/", false},
	{"a/+", false},
	{"b", false},
	{"a+(a(", false},
	{"a+(a/", false},
	{"a+)a)", false},
}

func GrammarTest(t *testing.T) {
	grammarParser := initParser()

	for _, testPair := range testPairs {
		if grammarParser.Parse(testPair.input) != testPair.expected {
			t.Errorf("Expected %v be %v", testPair.input, testPair.expected)
		}
	}
}
