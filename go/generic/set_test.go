package generic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjudice/util/go/generic"
)

func TestSet(t *testing.T) {
	x := []int{1, 2, 3, 4, 5, 5, 5, 5}
	s := generic.NewSet(x)
	assert.EqualValues(t, s.Has(1), true, "set should contain 1")
	assert.EqualValues(t, s.Has(5), true, "set should contain 5")
	assert.EqualValues(t, s.Has(10), false, "set should not have 10")
	assert.EqualValues(t, s.Len(), 5, "set length should be 5")
	s.Delete(1)
	assert.EqualValues(t, s.Has(1), false, "set should not contain 1")
	assert.EqualValues(t, s.Len(), 4, "set length should be 4")
}
