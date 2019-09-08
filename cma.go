package capi

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type CMAConfig struct {
	AccessToken string
	SpaceID string
	Environment string
	OrganisationID string
}

type cmaClient struct {
	Headers []header
	BaseUrl string
	Url string
	CMAConfig
}

type header struct {
	key string
	value string
}

func NewCMAClient(config CMAConfig) cmaClient {
	var cc cmaClient
	bearer := fmt.Sprintf("Bearer %s", config.AccessToken)
	cc.SpaceID = config.SpaceID
	cc.Environment = config.Environment
	cc.OrganisationID = config.OrganisationID
	cc.AccessToken = config.AccessToken
	cc.BaseUrl = "https://api.contentful.com"
	cc.addHeader("Authorization", bearer)
	cc.addHeader("X-Contentful-Organization", config.OrganisationID)
	cc.addHeader("Content-Type", "application/vnd.contentful.management.v1+json")

	return cc
}

func (c *cmaClient) call(method, endpoint string, body []byte) *http.Response {
	verb := strings.ToUpper(method)
	url := fmt.Sprintf("%s%s", c.BaseUrl, endpoint)

	req, err := http.NewRequest(verb, url, bytes.NewReader(body))

	if err != nil {
		log.Panic(err)
	}

	return makeRequest(c.Headers, req)
}

func (c *cmaClient) addHeader(key, value string) {
	h := header{
		key: key,
		value: value,
	}
	c.Headers = append(c.Headers, h)
}

func makeRequest(headers []header, req *http.Request) *http.Response {
	for _, h := range headers {
		req.Header.Set(h.key, h.value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Panic(err)
	}

	return resp
}

func readRequestBody(requestBody io.ReadCloser) []byte {
	body, err := ioutil.ReadAll(io.LimitReader(requestBody, 1048576))

	if err != nil {
		panic(err)
	}

	defer requestBody.Close()

	return body
}
