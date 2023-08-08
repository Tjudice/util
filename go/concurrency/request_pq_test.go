package concurrency_test

import (
	"context"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tjudice/util/go/concurrency"
)

func TestPrioritizedRequestQueue(t *testing.T) {
	results := []int{}
	m := sync.Mutex{}
	getter := func(_ context.Context, i int) (int, error) {
		// fill up pq requests initially
		if i == 0 {
			time.Sleep(time.Second)
			return 0, nil
		}
		time.Sleep(100 * time.Millisecond)
		return i, nil
	}
	less := func(i int, j int) bool {
		return i < j
	}
	ctx := context.TODO()
	rpq := concurrency.NewPrioritizedRequestQueue(1, less, getter)
	go rpq.Do(ctx, 0)
	time.Sleep(50 * time.Millisecond)
	for i := 1; i < 50; i = i + 1 {
		i := rand.Intn(50)
		go func() {
			res, _ := rpq.Do(ctx, i)
			m.Lock()
			defer m.Unlock()
			results = append(results, res)
		}()
	}
	time.Sleep(10 * time.Second)
	lastValue := -1
	for _, v := range results {
		assert.LessOrEqual(t, lastValue, v, "results not properly ordered")
	}
}
