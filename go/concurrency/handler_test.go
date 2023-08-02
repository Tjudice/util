package concurrency_test

import (
	"concurrency"
	"context"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDoNoError(t *testing.T) {
	t.Parallel()
	counter := atomic.Int64{}
	testFn := func(i int64) error {
		switch i {
		case 4:
			counter.Add(i)
		default:
			time.Sleep(time.Second * time.Duration(i))
			counter.Add(i)
		}
		return nil
	}
	assert.NoError(t, concurrency.Do(5, testFn, 3, 5, 1, 1, 4), "Expected no error but got error")
	assert.EqualValues(t, 14, counter.Load(), "Expected 14")
}

func TestDoError(t *testing.T) {
	t.Parallel()
	counter := atomic.Int64{}
	testFn := func(i int64) error {
		switch i {
		case 2:
			counter.Add(i)
			return fmt.Errorf("2 test")
		default:
			time.Sleep(time.Second * time.Duration(i))
			counter.Add(i)
		}
		return nil
	}
	assert.Error(t, concurrency.Do(5, testFn, 2, 3, 5, 3, 1, 2), "Expected an error to occur but got no error")
	assert.EqualValues(t, 16, counter.Load(), "Error when calling Do() should not stop execution")
}

func TestDoContextNoError(t *testing.T) {
	t.Parallel()
	counter := atomic.Int64{}
	testFn := func(_ context.Context, i int64) error {
		switch i {
		default:
			time.Sleep(time.Second * time.Duration(i))
			counter.Add(i)
			return nil
		}
	}
	ctx := context.TODO()
	assert.NoError(t, concurrency.DoContext(ctx, 5, testFn, 1, 2, 3, 1, 1), "Expected an error to occur but go no error")
	assert.EqualValues(t, 8, counter.Load(), "Expected function to return before adding other counters")
}

func TestDoContextError(t *testing.T) {
	t.Parallel()
	counter := atomic.Int64{}
	testFn := func(_ context.Context, i int64) error {
		switch i {
		case 2:
			counter.Add(i)
			return fmt.Errorf("2 test")
		default:
			time.Sleep(time.Hour)
			counter.Add(i)
			return nil
		}
	}
	ctx := context.TODO()
	assert.Error(t, concurrency.DoContext(ctx, 5, testFn, 10, 10, 10, 10, 2), "Expected an error to occur but go no error")
	assert.EqualValues(t, 2, counter.Load(), "Expected function to return before adding other counters")
}

func TestDoChan(t *testing.T) {
	t.Parallel()
	counter := atomic.Int64{}
	testFn := func(i int64) (int64, error) {
		switch i {
		case 2:
			return 2, fmt.Errorf("2")
		default:
			return i, nil
		}
	}
	ch := concurrency.DoChan(5, 5, testFn, 4, 3, 1, 1, 2)
	for {
		next, ok := <-ch
		if !ok {
			break
		}
		id, val, err := next.Value()
		counter.Add(val)
		if err != nil {
			assert.EqualValues(t, id, 2, "response with id != 2 should not have caused an error")
			continue
		}
		assert.NotEqualValues(t, id, 2, "response with id 2 should have caused an error")
	}
	assert.EqualValues(t, 11, counter.Load(), "counter should add up to 11")
}
