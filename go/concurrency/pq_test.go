package concurrency_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trevorjudice/util/go/concurrency"
)

func TestBasic(t *testing.T) {
	cmpFn := func(i, j int64) bool {
		return i > j
	}
	pq := concurrency.NewPriorityQueue[int64](100000, cmpFn)
	go func() {
		for i := int64(0); i < 1000000; i = i + 1 {
			pq.Push(i)
		}
	}()
	cnt := 1000000
	last := int64(-1)
	for {
		next, ok := pq.TryPop()
		if !ok {
			continue
		}
		assert.LessOrEqual(t, last, next, "previous value is >= current value")
		cnt -= 1
		if cnt <= 0 {
			break
		}
	}
}
