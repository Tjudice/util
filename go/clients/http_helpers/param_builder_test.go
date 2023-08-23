package http_helpers_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjudice/util/go/clients/http_helpers"
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

func TestParamBuilderAddIntegerTypes(t *testing.T) {
	url := url.Values{}
	q := http_helpers.NewURLEncoder(url)
	q.Add("int", 1)
	q.Add("int8", int8(2))
	q.Add("int16", int16(3))
	assert.EqualValues(t, q.Get("test"), "hi")
}
