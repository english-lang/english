package internal

import (
	"github.com/english-lang/english/ast/internal/token"
	"github.com/english-lang/english/ast/tag"
)

// Node type
type Node map[string]interface{}

// NewWord type
func NewWord(w, t interface{}) interface{} {
	rt := Node{"Type": "Word"}
	rt["Word"] = string((w.(*token.Token)).Lit)
	rt["WordType"] = tag.WordType((t.(*token.Token)).Lit)
	return &rt
}
