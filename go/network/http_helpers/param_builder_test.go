package http_helpers_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjudice/util/go/network/http_helpers"
)

func TestParamBuilderAddString(t *testing.T) {
	url := url.Values{}
	q := http_helpers.NewURLEncoder(url)
	q.Add("test", "hi")
	assert.EqualValues(t, q.Get("test"), "hi")
}

func TestParamBuilderAddBytes(t *testing.T) {
	url := url.Values{}
	q := http_helpers.NewURLEncoder(url)
	q.Add("test", []byte("hi"))
	assert.EqualValues(t, q.Get("test"), "hi")
}

func TestParamBuilderAddNumberTypes(t *testing.T) {
	url := url.Values{}
	q := http_helpers.NewURLEncoder(url)
	q.Add("int", 1)
	q.Add("int8", int8(2))
	q.Add("int16", int16(3))
	q.Add("int32", int32(4))
	q.Add("int64", int64(5))
	q.Add("uint", uint(6))
	q.Add("uint8", uint8(7))
	q.Add("uint16", uint16(8))
	q.Add("uint32", uint32(9))
	q.Add("uint64", uint64(10))
	q.Add("float32", float32(12.0))
	q.Add("float64", float64(13.0))
	assert.EqualValues(t, q.Get("int"), "1")
	assert.EqualValues(t, q.Get("int8"), "2")
	assert.EqualValues(t, q.Get("int16"), "3")
	assert.EqualValues(t, q.Get("int32"), "4")
	assert.EqualValues(t, q.Get("int64"), "5")
	assert.EqualValues(t, q.Get("uint"), "6")
	assert.EqualValues(t, q.Get("uint8"), "7")
	assert.EqualValues(t, q.Get("uint16"), "8")
	assert.EqualValues(t, q.Get("uint32"), "9")
	assert.EqualValues(t, q.Get("uint64"), "10")
	assert.EqualValues(t, q.Get("float32"), "12.0")
	assert.EqualValues(t, q.Get("float64"), "13.0")
}
