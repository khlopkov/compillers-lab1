package parser

type PrecedenceRelation byte

const (
	Less PrecedenceRelation = iota + 1
	Equal
	Greater
)

func getIndexIdentifier(s Symbol) string {
	identifier := ""

	if s.SymbolType() == NonTerminal {
		identifier += "non-terminal"
	} else {
		identifier += "terminal"
	}

	identifier += s.Identifier()

	return identifier
}

type PrecedenceMatrix interface {
	GetRelationOf(a, b Symbol) PrecedenceRelation
}

func NewPrecedenceMatrix(relations ...Relation) PrecedenceMatrix {
	matrix := precedenceMatrix{
		relations: make(map[string]map[string]PrecedenceRelation),
	}

	for _, r := range relations {
		aIdx := getIndexIdentifier(r.a)
		bIdx := getIndexIdentifier(r.b)

		if matrix.relations[aIdx] == nil {
			matrix.relations[aIdx] = make(map[string]PrecedenceRelation)
		}

		matrix.relations[aIdx][bIdx] = r.relation
	}

	return &matrix
}

type precedenceMatrix struct {
	relations map[string]map[string]PrecedenceRelation
}

func (m *precedenceMatrix) GetRelationOf(a, b Symbol) PrecedenceRelation {
	aIdx := getIndexIdentifier(a)
	bIdx := getIndexIdentifier(b)

	return m.relations[aIdx][bIdx]
}

type Relation struct {
	a        Symbol
	b        Symbol
	relation PrecedenceRelation
}

func NewRelation(a Symbol, relation PrecedenceRelation, b Symbol) Relation {
	return Relation{
		a:        a,
		b:        b,
		relation: relation,
	}
}
