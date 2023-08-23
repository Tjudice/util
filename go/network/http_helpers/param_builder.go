package http_helpers

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type UrlEncoder interface {
	Add(key string, value any) UrlEncoder
	Del(key string) UrlEncoder
	Encode() (string, error)
	Get(key string) string
	Has(key string) bool
	Set(key string, value any) UrlEncoder
}

type URLMarshaler interface {
	MarshalURL() (string, error)
}

type QueryStringBuilder struct {
	underlying url.Values
	err        error
}

func NewURLEncoder(url url.Values) UrlEncoder {
	return &QueryStringBuilder{url, nil}
}

func (q *QueryStringBuilder) Add(key string, value any) UrlEncoder {
	s, err := anyToString(value)
	if err != nil {
		q.err = err
		return q
	}
	q.underlying.Add(key, s)
	return q
}

func (q *QueryStringBuilder) Del(key string) UrlEncoder {
	q.underlying.Del(key)
	return q
}

func (q *QueryStringBuilder) Encode() (string, error) {
	return q.underlying.Encode(), q.err
}

func (q *QueryStringBuilder) Get(key string) string {
	return q.underlying.Get(key)
}

func (q *QueryStringBuilder) Has(key string) bool {
	return q.underlying.Has(key)
}

func (q *QueryStringBuilder) Set(key string, value any) UrlEncoder {
	s, err := anyToString(value)
	if err != nil {
		q.err = err
		return q
	}
	q.underlying.Set(key, s)
	return q
}

func anyToString(value any) (string, error) {
	switch v := value.(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	case URLMarshaler:
		return v.MarshalURL()
	case json.Marshaler:
		bts, err := v.MarshalJSON()
		return string(bts), err
	default:
		return string(fmt.Sprint(v)), nil
	}
}
