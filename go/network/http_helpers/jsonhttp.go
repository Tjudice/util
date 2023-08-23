package http_helpers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func GetJSON[T any](ctx context.Context, cl *http.Client, url string, requestBody io.Reader) (resp T, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, requestBody)
	if err != nil {
		return resp, err
	}
	req.Header.Add("content-type", "application/json")
	return getJSON[T](ctx, cl, req)
}

func GetJSONFn[T any](ctx context.Context, cl *http.Client, url string, requestBody io.Reader, middleware func(r *http.Request)) (resp T, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, requestBody)
	if err != nil {
		return resp, err
	}
	req.Header.Add("content-type", "application/json")
	// call middleware after all request modifications are made
	middleware(req)
	return getJSON[T](ctx, cl, req)
}

func getJSON[T any](ctx context.Context, cl *http.Client, req *http.Request) (resp T, err error) {
	res, err := cl.Do(req)
	if err != nil {
		return resp, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
