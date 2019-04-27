package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProseEnglish(t *testing.T) {
	d, e := ParseArticle("Go is an open-source programming language created at Google.")
	if e != nil {
		t.Error(e)
	}
	assert.Equal(t, "NNP", d.Paragraphs[0].Sentences[0].Tokens[0].Tag)
}

func TestSimpleDeclarativeSentence(t *testing.T) {

	s := "Apple is an object."
	a, _ := ParseArticle(s)
	words := a.Paragraphs[0].Sentences[0].Tokens

	assert.Equal(t, "NNP", words[0].Tag)
	assert.Equal(t, "VBZ", words[1].Tag)
	assert.Equal(t, "DT", words[2].Tag)
	assert.Equal(t, "NN", words[3].Tag)
	assert.Equal(t, ".", words[4].Tag)

}
