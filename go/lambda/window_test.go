package lambda_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	lambda "github.com/tjudice/util/go/algorithms"
)

func TestWindowSimple(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	expected := [][]int{
		{1, 2, 3},
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{3, 4, 5},
		{4, 5},
	}
	result := lambda.Window(1, 2, items)
	assert.EqualValues(t, expected, result, "Window: Slices should be equal")
	resultFn := make([][]int, 0, len(items))
	lambda.WindowFn(1, 2, items, func(group []int) {
		resultFn = append(resultFn, group)
	})
	assert.EqualValues(t, expected, resultFn, "WindowFn: Slices should be equal")
}

func TestWindowZeroPreceding(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	expected := [][]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
		{4, 5},
		{5},
	}
	result := lambda.Window(0, 2, items)
	assert.EqualValues(t, expected, result, "Window: Slices should be equal")
	resultFn := make([][]int, 0, len(items))
	lambda.WindowFn(0, 2, items, func(group []int) {
		resultFn = append(resultFn, group)
	})
	assert.EqualValues(t, expected, resultFn, "WindowFn: Slices should be equal")
}

func TestWindowZeroFollowing(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	expected := [][]int{
		{1},
		{1, 2},
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
	}
	result := lambda.Window(2, 0, items)
	assert.EqualValues(t, expected, result, "Window: Slices should be equal")
	resultFn := make([][]int, 0, len(items))
	lambda.WindowFn(2, 0, items, func(group []int) {
		resultFn = append(resultFn, group)
	})
	assert.EqualValues(t, expected, resultFn, "WindowFn: Slices should be equal")
}

func TestWindowBothZero(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	expected := [][]int{
		{1},
		{2},
		{3},
		{4},
		{5},
	}
	result := lambda.Window(0, 0, items)
	assert.EqualValues(t, expected, result, "Window: Slices should be equal")
	resultFn := make([][]int, 0, len(items))
	lambda.WindowFn(0, 0, items, func(group []int) {
		resultFn = append(resultFn, group)
	})
	assert.EqualValues(t, expected, resultFn, "WindowFn: Slices should be equal")
}
