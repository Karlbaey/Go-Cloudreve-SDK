package cloudreve

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/cookiejar"
)

// A Client is a core structure that interact
// with a Cloudreve server.
type Client struct {
	Host       string       // Cloudreve server host. e.g. https://cloudreve.exapmle.com
	HTTPClient *http.Client // Send HTTP requests
}

func NewClient(host string) *Client {
	jar, _ := cookiejar.New(nil)

	return &Client{
		Host: host,
		HTTPClient: &http.Client{
			Jar: jar,
		},
	}
}

// do is a private helper method.
//
// body: a request object to be sent, which will be serialize to JSON.
//
// res: receive a data object in a response.
func (c *Client) do(method, path string, body, res interface{}) error {
	// concatenate full URL
	url := c.Host + path

	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Parse the general response structure
	// Note: Here, Response[T] is the generic structure defined in models.go
	var apiResponse APIResponse[any]
	if res != nil {
		apiResponse.Data = res
	}
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return err
	}

	// check errors from API
	if apiResponse.Code != 0 {
		return &APIError{Code: apiResponse.Code, Msg: apiResponse.Msg}
	}

	return nil
}

type APIError struct {
	Code int
	Msg  string
}

func (a *APIError) Error() string {
	return a.Msg
}
