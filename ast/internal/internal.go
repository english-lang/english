package internal

import (
	"github.com/english-lang/english/ast/internal/token"
	"github.com/english-lang/english/ast/tag"
)

// AstNode type
type AstNode map[string]interface{}

// NodeType string
func (n *AstNode) NodeType() string {
	return ((*n)["Type"]).(string)
}

// NewWord type
func NewWord(w, t interface{}) interface{} {
	rt := AstNode{"Type": "Word"}
	rt["Word"] = string((w.(*token.Token)).Lit)
	rt["WordType"] = tag.WordType(t.(string))
	return &rt
}

// NewSentence type
func NewSentence(t interface{}, words ...interface{}) interface{} {
	rt := AstNode{"Type": t}
	rt["Words"] = words
	return rt
}
