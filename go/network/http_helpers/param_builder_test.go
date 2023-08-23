package http_helpers_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjudice/util/go/network/http_helpers"
)

func TestAddString(t *testing.T) {
	q := http_helpers.NewURLEncoder(url.Values{})
	q.Add("test", "hi")
	assert.EqualValues(t, "hi", q.Get("test"))
}

func TestAddBytes(t *testing.T) {
	q := http_helpers.NewURLEncoder(url.Values{})
	q.Add("test", []byte("hi"))
	assert.EqualValues(t, "hi", q.Get("test"))
}

func TestAddNumberTypes(t *testing.T) {
	q := http_helpers.NewURLEncoder(url.Values{})
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
	q.Add("float32", float32(11.1))
	q.Add("float64", float64(12.1))
	assert.EqualValues(t, "1", q.Get("int"))
	assert.EqualValues(t, "2", q.Get("int8"))
	assert.EqualValues(t, "3", q.Get("int16"))
	assert.EqualValues(t, "4", q.Get("int32"))
	assert.EqualValues(t, "5", q.Get("int64"))
	assert.EqualValues(t, "6", q.Get("uint"))
	assert.EqualValues(t, "7", q.Get("uint8"))
	assert.EqualValues(t, "8", q.Get("uint16"))
	assert.EqualValues(t, "9", q.Get("uint32"))
	assert.EqualValues(t, "10", q.Get("uint64"))
	assert.EqualValues(t, "11.1", q.Get("float32"))
	assert.EqualValues(t, "12.1", q.Get("float64"))
}

func TestDel(t *testing.T) {
	q := http_helpers.NewURLEncoder(url.Values{})
	q.Add("test", "hi")
	assert.EqualValues(t, "hi", q.Get("test"))
	q.Del("test")
	assert.EqualValues(t, "", q.Get("test"))
}

func TestEncodeSingle(t *testing.T) {
	q := http_helpers.NewURLEncoder(url.Values{})
	q.Add("test", "val1")
	encoded, err := q.Encode()
	assert.NoError(t, err)
	assert.EqualValues(t, "test=val1", encoded)
}

func TestEncodeMultipleSameKey(t *testing.T) {
	q := http_helpers.NewURLEncoder(url.Values{})
	q.Add("test", "val1")
	q.Add("test", "val2")
	encoded, err := q.Encode()
	assert.NoError(t, err)
	assert.EqualValues(t, "test=val1&test=val2", encoded)
}
