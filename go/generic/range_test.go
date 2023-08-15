package generic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjudice/util/go/generic"
)

func TestRangeSimple(t *testing.T) {
	testInfo := generic.DivideRangeInclusive(0, 3, 1)
	expected := []generic.Range[int]{
		{
			Start: 0,
			End:   0,
		},
		{
			Start: 1,
			End:   1,
		},
		{
			Start: 2,
			End:   2,
		},
		{
			Start: 3,
			End:   3,
		},
	}
	assert.EqualValues(t, expected, testInfo, "Arrays should be equal")
}
