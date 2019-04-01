package parser

type Rule interface {
	Identifier() Symbol
	Symbols() []Symbol
	IsEqual(rule Rule) bool
}

func NewRule(identifier Symbol, symbols ...Symbol) Rule {
	return &rule{
		symbols:    symbols,
		identifier: identifier,
	}
}

type rule struct {
	symbols    []Symbol
	identifier Symbol
}

func (r *rule) Symbols() []Symbol {
	return r.symbols
}

func (r *rule) IsEqual(rule Rule) bool {
	if len(rule.Symbols()) != len(r.Symbols()) {
		return false
	}

	comparedSymbols := rule.Symbols()
	for i, s := range r.Symbols() {
		if !comparedSymbols[i].IsEqual(s) {
			return false
		}
	}

	return true
}

func (r *rule) Identifier() Symbol {
	return r.identifier
}
