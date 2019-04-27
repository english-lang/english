package internal_test

import (
	"testing"

	"github.com/english-lang/english/ast/internal/lexer"
	"github.com/english-lang/english/ast/internal/parser"
)

func TestParseSimpleDeclarati(t *testing.T) {
	simpleSentence := `Apple<NNP> is<VBZ> an<DT> object<NN> .<.>`
	s := lexer.NewLexer([]byte(simpleSentence))
	p := parser.NewParser()
	_, err := p.Parse(s)
	if err != nil {
		t.Error(err)
	}

}
