package concurrency_test

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/tjudice/util/go/concurrency"
)

func TestPrioritizedRequestQueue(t *testing.T) {
	v := []int{}
	m := sync.Mutex{}
	getter := func(_ context.Context, i int) (int, error) {
		// fill up pq requests initially
		if i == 0 {
			time.Sleep(time.Second)
			return 0, nil
		}
		time.Sleep(time.Duration(i) * time.Second)
		return i, nil
	}
	less := func(i int, j int) bool {
		return i < j
	}
	ctx := context.TODO()
	rpq := concurrency.NewPrioritizedRequestQueue(1, less, getter)
	go rpq.Do(ctx, 0)
	for i := 1; i < 4; i = i + 1 {
		i := i
		go func() {
			res, _ := rpq.Do(ctx, i)
			m.Lock()
			defer m.Unlock()
			v = append(v, res)
		}()
	}
	time.Sleep(10 * time.Second)
	log.Println(v)
}
