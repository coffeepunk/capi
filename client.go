package capi

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	Config
	headers []Header
	baseUrl string
}

type Header struct {
	key   string
	value string
}

func (c *Client) AddHeader(key, value string) {
	h := Header{
		key:   key,
		value: value,
	}
	c.headers = append(c.headers, h)
}

func readRequestBody(requestBody io.ReadCloser) []byte {
	body, err := ioutil.ReadAll(io.LimitReader(requestBody, 1048576))

	if err != nil {
		panic(err)
	}

	defer requestBody.Close()

	return body
}

func (c *Client) call(method, endpoint string, body []byte) (*http.Response, error) {
	verb := strings.ToUpper(method)
	url := fmt.Sprintf("%s%s", c.baseUrl, endpoint)
	req, err := http.NewRequest(verb, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	for _, h := range c.headers {
		req.Header.Set(h.key, h.value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
