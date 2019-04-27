package internal_test

import (
	"encoding/json"
	"testing"

	"github.com/english-lang/english/ast/internal/lexer"
	"github.com/english-lang/english/ast/internal/parser"
)

func TestParseSimpleDeclarati(t *testing.T) {
	simpleSentence := `Apple<NNP> is<VBZ> an<DT> object<NN> .<.>`
	s := lexer.NewLexer([]byte(simpleSentence))
	p := parser.NewParser()
	r, err := p.Parse(s)
	v, _ := json.Marshal(r)
	t.Log(string(v))
	if err != nil {
		t.Error(err)
	}
}
