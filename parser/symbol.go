package parser

type SymbolType byte

const (
	Terminal SymbolType = iota + 1
	NonTerminal
)

type Symbol interface {
	Identifier() string
	SymbolType() SymbolType
	IsEqual(symbol Symbol) bool
}

func NewSymbol(
	identifier string,
	symbolType SymbolType,
) Symbol {
	return &symbol{
		identifier: identifier,
		symbolType: symbolType,
	}
}

type symbol struct {
	identifier string
	symbolType SymbolType
}

func (s *symbol) Identifier() string {
	return s.identifier
}

func (s *symbol) SymbolType() SymbolType {
	return s.symbolType
}

func (s *symbol) IsEqual(symbol Symbol) bool {
	return symbol.Identifier() == s.Identifier() &&
		symbol.SymbolType() == s.SymbolType()
}
