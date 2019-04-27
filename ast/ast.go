package ast

import (
	"strings"

	"github.com/english-lang/english/ast/internal"
	"gopkg.in/jdkato/prose.v2"
)

// Article type
type Article struct {
	Source     string
	Paragraphs []*Paragraph
}

// Paragraph type
type Paragraph struct {
	Source    string
	Sentences []*Sentence
}

// Sentence type
type Sentence struct {
	Source string
	Tokens []*prose.Token
	Ast    *internal.AstNode
}

// ParseArticle from source
func ParseArticle(article string) (*Article, error) {
	rt := &Article{Source: article, Paragraphs: []*Paragraph{}}
	for _, para := range strings.Split(article, "\n") {
		oPara := &Paragraph{Source: para}
		sPara, err := prose.NewDocument(para)
		if err != nil {
			return nil, err
		}
		for _, sentence := range sPara.Sentences() {
			sSentence, err := prose.NewDocument(sentence.Text)
			oSentence := &Sentence{Source: sentence.Text, Tokens: []*prose.Token{}}
			if err != nil {
				return nil, err
			}
			for _, sToken := range sSentence.Tokens() {
				// create ref with new instance, avoid loop pointer issue.
				oToken := &prose.Token{Text: sToken.Text, Label: sToken.Label, Tag: sToken.Tag}
				oSentence.Tokens = append(oSentence.Tokens, oToken)
			}
			oPara.Sentences = append(oPara.Sentences, oSentence)
		}
		rt.Paragraphs = append(rt.Paragraphs, oPara)
	}
	return rt, nil
}
