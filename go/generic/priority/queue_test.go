package priority_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trevorjudice/util/go/generic/priority"
)

func TestSimpleInsert(t *testing.T) {
	cmp := func(i, j int64) bool {
		return i > j
	}
	q := priority.NewPriorityQueue(10, cmp)
	q.Push(3)
	q.Push(4)
	q.Push(17)
	q.Push(-1)
	q.Push(2)
	assert.EqualValues(t, 17, q.Pop(), "first value popped should be 17")
	assert.EqualValues(t, 4, q.Pop(), "second value should be 4")
	assert.EqualValues(t, 3, q.Pop(), "third value should be 3")
	assert.EqualValues(t, 2, q.Pop(), "fourth value should be 2")
	assert.EqualValues(t, 1, q.Len(), "length should be one")
	assert.EqualValues(t, -1, q.Pop(), "fifth value should be -1")
}
