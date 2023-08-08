package algorithms_test

import (
	"testing"

	"github.com/tjudice/util/go/algorithms"
)

type Ohlcv struct {
	Time  int64
	Open  float64
	Close float64
	High  float64
	Low   float64
}

func (o Ohlcv) Position() int64 {
	return o.Time
}

func TestLocfOHLCVData(t *testing.T) {
	testData := []Ohlcv{
		{Time: 15, Open: 10, Close: 8, High: 13, Low: 7},
		{Time: 13, Open: 8, Close: 9, High: 12, Low: 7},
		{Time: 3, Open: 10, Close: 8, High: 13, Low: 7},
		{Time: 18, Open: 10, Close: 24, High: 28, Low: 10},
	}
	lastObserved := algorithms.LocfObservations(0, 25, 5, testData, func(t int64, last Ohlcv) Ohlcv {
		if t == last.Time {
			return last
		}
		return Ohlcv{
			Time: t,
			Open: last.Close,
		}
	})
}
