package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayInt_Iterator(t *testing.T) {
	data := ArrayInt{1, 3, 5, 7, 8}
	iterator := data.Iterator()

	// i 用于测试
	i := 0
	for iterator.HasNext() {
		assert.Equal(t, data[i], iterator.CurrentItem())
		iterator.Next()
		i++
	}
}
