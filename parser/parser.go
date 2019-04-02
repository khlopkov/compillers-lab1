package parser

type Parser interface {
	Parse(input string) bool
}

func NewParser(grammar *Grammar) Parser {
	return &parser{
		grammar: grammar,
	}
}

type symbolStack struct {
	stack []Symbol
}

func newSymbolStack() *symbolStack {
	stack := make([]Symbol, 0, 1)
	stack = append(stack, NewSymbol("^", StartSymbol))

	return &symbolStack{
		stack: stack,
	}
}

func (s *symbolStack) top() Symbol {
	return s.stack[len(s.stack)-1]
}

func (s *symbolStack) pop() Symbol {
	var symbol Symbol
	symbol, s.stack = s.stack[len(s.stack)-1], s.stack[:len(s.stack)-1]
	return symbol
}

func (s *symbolStack) append(symbol Symbol) {
	s.stack = append(s.stack, symbol)
}

func (s *symbolStack) reduce(grammar *Grammar) bool {
	current := s.pop()
	base := make([]Symbol, 0, 8)
	base = append(base, current)

	for len(s.stack) != 0 {
		next := s.top()

		relation := grammar.matrix.GetRelationOf(next, current)
		if relation == Equal {
			current = s.pop()
			base = append([]Symbol{current}, base...)
		} else {
			break
		}
	}

	baseRule := NewRule(nil, base...)
	for _, r := range grammar.rules {
		if r.IsEqual(baseRule) {
			s.append(r.Identifier())
			return true
		}
	}

	return false
}

type parser struct {
	grammar *Grammar
}

func (p *parser) Parse(input string) bool {
	stack := newSymbolStack()

	i := 0
	for {
		if len(input) == i {
			if len(stack.stack) == 2 &&
				stack.stack[0].SymbolType() == StartSymbol &&
				p.grammar.rules[0].Identifier().IsEqual(stack.stack[1]) {
				return true
			}

			reduceSuccessful := stack.reduce(p.grammar)
			if !reduceSuccessful {
				return false
			}
			continue
		}

		symbol := input[i]

		symbolFromInput := NewSymbol(string(symbol), Terminal)
		symbolFromStack := stack.top()

		if symbolFromStack.SymbolType() == StartSymbol {
			stack.append(symbolFromInput)
			i++
			continue
		}

		relation := p.grammar.matrix.GetRelationOf(symbolFromStack, symbolFromInput)
		if relation == Greater {
			reduceSuccessful := stack.reduce(p.grammar)
			if !reduceSuccessful {
				return false
			}
		} else {
			stack.append(symbolFromInput)
			i++
		}
	}
}
