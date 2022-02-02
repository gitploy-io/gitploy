package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gitploy-io/gitploy/pkg/e"
)

type (
	Client struct {
		// Reuse a single struct instead of allocating one for each service on the heap.
		common *client

		// Services used for talking to different parts of the Gitploy API.
		Repo *RepoService
	}

	client struct {
		// HTTP client used to communicate with the API.
		httpClient *http.Client

		// Base URL for API requests. BaseURL should
		// always be specified with a trailing slash.
		BaseURL *url.URL
	}

	service struct {
		*client
	}

	ErrorResponse struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
)

func NewClient(host string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseURL, _ := url.Parse(host)

	c := &Client{
		common: &client{httpClient: httpClient, BaseURL: baseURL},
	}

	c.Repo = &RepoService{client: c.common}

	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
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

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.
func (c *client) Do(ctx context.Context, req *http.Request, v interface{}) error {
	if ctx == nil {
		return fmt.Errorf("There is no context")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to request: %w", err)
	}

	defer res.Body.Close()

	// Return internal errors
	if res.StatusCode > 299 {
		errRes := &ErrorResponse{}

		out, _ := ioutil.ReadAll(res.Body)
		if err := json.Unmarshal(out, errRes); err != nil {
			return fmt.Errorf("Failed to parse an error response: %w", err)
		}

		return e.NewErrorWithMessage(e.ErrorCode(errRes.Code), errRes.Message, nil)
	}

	if v != nil {
		return json.NewDecoder(res.Body).Decode(v)
	}

	return nil
}
