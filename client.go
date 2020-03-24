package capi

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type client struct {
	Config
	headers []header
	baseUrl string
}

type header struct {
	key   string
	value string
}

/*
func Open(config Config) Connection {
	bearer := fmt.Sprintf("Bearer %s", config.AccessToken)

	var conn Connection
	conn.CDA.SpaceID = config.SpaceID
	conn.CDA.Environment = config.Environment
	conn.CDA.AccessToken = config.AccessToken
	conn.CDA.baseUrl = getBaseUrl("CDA")
	conn.CDA.addHeader("Authorization", bearer)

	conn.CMA.SpaceID = config.SpaceID
	conn.CMA.Environment = config.Environment
	conn.CMA.AccessToken = config.AccessToken
	conn.CMA.baseUrl = getBaseUrl("CMA")
	conn.CMA.addHeader("Authorization", bearer)
	conn.CMA.addHeader("X-Contentful-Organization", config.OrganisationID)
	conn.CMA.addHeader("Content-Type", "application/vnd.contentful.management.v1+json")

	return conn
}

func getBaseUrl(api string) string {
	var url string
	switch api {
	case "CDN":
		url = "https://cdn.contentful.com"
	case "CPA":
		url = "https://preview.contentful.com"
	case "CMA":
		url = "https://api.contentful.com"
	case "IA":
		url = "https://images.ctfassets.net"
	default:
		url = "https://cdn.contentful.com"
	}

	return url
}
*/

func (c *client) addHeader(key, value string) {
	h := header{
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

func (c *client) call(method, endpoint string, body []byte) (*http.Response, error) {
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
