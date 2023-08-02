package priority_test

import (
	"testing"

	"github.com/trevorjudice/util/go/generic/heap"

	"github.com/stretchr/testify/assert"
	"github.com/trevorjudice/util/go/generic/priority"
)

func TestSimpleInsert(t *testing.T) {
	cmp := func(i, j int64) bool {
		return i > j
	}
	q := priority.NewQueue(10, cmp)
	heap.Push[int64](q, 3)
	heap.Push[int64](q, 4)
	heap.Push[int64](q, 17)
	heap.Push[int64](q, -1)
	assert.EqualValues(t, 17, q.Pop(), "first value popped should be 17")
	assert.EqualValues(t, 4, q.Pop(), "second value should be 4")
	assert.EqualValues(t, 3, q.Pop(), "third value should be 3")
	assert.EqualValues(t, -1, q.Pop(), "fourth value should be -1")

}
