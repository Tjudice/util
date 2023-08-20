package clients

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func Get[T any](ctx context.Context, cl *http.Client, url string, requestBody io.Reader) (resp T, err error) {
	req, err := http.NewRequest(http.MethodGet, url, requestBody)
	if err != nil {
		return resp, err
	}
	req.Header.Add("content-type", "application/json")
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
