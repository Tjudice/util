package generic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjudice/util/go/generic"
)

func TestSet(t *testing.T) {
	x := []int{1, 2, 3, 4, 5, 5, 5, 5}
	s := generic.NewSet(x)
	assert.EqualValues(t, true, s.Has(1), "set should contain 1")
	assert.EqualValues(t, true, s.Has(5), "set should contain 5")
	assert.EqualValues(t, false, s.Has(10), "set should not have 10")
	assert.EqualValues(t, 5, s.Len(), "set length should be 5")
	s.Delete(1)
	assert.EqualValues(t, false, s.Has(1), "set should not contain 1")
	assert.EqualValues(t, 4, s.Len(), "set length should be 4")
}

type NonComparable struct {
	Item int
}

func TestSetFn(t *testing.T) {
	x := []*NonComparable{
		{1},
		{2},
		{3},
		{4},
		{5},
		{5},
		{5},
	}
	s := generic.NewSetFn(x, func(f *NonComparable) int {
		return f.Item
	})
	assert.EqualValues(t, true, s.Has(1), "set should contain 1")
	assert.EqualValues(t, true, s.Has(5), "set should contain 5")
	assert.EqualValues(t, false, s.Has(10), "set should not have 10")
	assert.EqualValues(t, 5, s.Len(), "set length should be 5")
	s.Delete(1)
	assert.EqualValues(t, false, s.Has(1), "set should not contain 1")
	assert.EqualValues(t, 4, s.Len(), "set length should be 4")
}
