package dbl_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
)

type DBL struct {
	baseURL   *url.URL
	userAgent string

	token      string
	httpClient *http.Client
}

// NewDBL creates a new top.gg client with the given token and HTTP client
func NewDBL(token string, client *http.Client) *DBL {
	if client == nil {
		client = &http.Client{}
	}

	base, _ := url.Parse(BaseURL)

	return &DBL{
		baseURL:    base,
		userAgent:  fmt.Sprintf("dbl-go/%s (%s) Golang/%s", Version, runtime.GOOS, runtime.Version()),
		token:      token,
		httpClient: client,
	}
}

func (c *DBL) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: "/api" + path}
	u := c.baseURL.ResolveReference(rel)

	fmt.Println("Calling ", rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Authorization", c.token)
	return req, nil
}

func (c *DBL) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
