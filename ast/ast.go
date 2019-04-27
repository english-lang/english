package ast

import (
	"gopkg.in/jdkato/prose.v2"
)

// Article type
type Article struct {
	Source    string
	Sentences []*Sentence
}

// Sentence type
type Sentence struct {
	Source string
	Tokens []*prose.Token
}

// ParseArticle from source
func ParseArticle(article string) (*Article, error) {
	rt := &Article{Source: article, Sentences: []*Sentence{}}
	sArticle, err := prose.NewDocument(article)
	if err != nil {
		return nil, err
	}
	for _, sentence := range sArticle.Sentences() {

		sSentence, err := prose.NewDocument(sentence.Text)

		if err != nil {
			return nil, err
		}
		oSentence := &Sentence{Source: sentence.Text, Tokens: []*prose.Token{}}

		for _, sToken := range sSentence.Tokens() {
			// create ref with new instance, avoid loop pointer issue.
			oToken := &prose.Token{Text: sToken.Text, Label: sToken.Label, Tag: sToken.Tag}
			oSentence.Tokens = append(oSentence.Tokens, oToken)
		}

		rt.Sentences = append(rt.Sentences, oSentence)
	}
	return rt, nil
}
