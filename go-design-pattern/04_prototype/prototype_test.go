package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestKeywords_Clone(t *testing.T) {
	updateAt, _ := time.Parse("2006", "2020")

	words := Keywords{
		"testA": &Keyword{
			word:      "testA",
			visit:     1,
			UpdatedAt: &updateAt,
		},
		"testB": &Keyword{
			word:      "testB",
			visit:     2,
			UpdatedAt: &updateAt,
		},
		"testC": &Keyword{
			word:      "testC",
			visit:     3,
			UpdatedAt: &updateAt,
		},
	}

	now := time.Now()

	updatedWords := []*Keyword{
		{
			word:      "testB",
			visit:     10,
			UpdatedAt: &now,
		},
	}

	got := words.Clone(updatedWords)
	assert.Equal(t, words["testA"], got["testA"])
	assert.NotEqual(t, words["testB"], got["testB"])
	assert.NotEqual(t, updatedWords[0], got["testB"])
	assert.Equal(t, words["testC"], got["testC"])
}
