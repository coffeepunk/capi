package capi

import (
	"net/url"
	"strings"
)

type SearchParameters struct {
	ContentType string `url:"content_type"`
	Include     string `url:"include"`
	Limit       string `url:"limit"`
	Fields      string `url:"fields"`
	Locale      string `url:"locale"`
	Skip        string `url:"skip"`
}

func buildQueryString(params SearchParameters) string {
	qs := url.Values{}
	qs.Add("include", "2")

	if params.ContentType != "" {
		qs.Add("content_type", params.ContentType)
	}

	if params.Include != "" {
		qs.Set("include", params.Include)
	}

	if params.Limit != "" {
		qs.Add("limit", params.Limit)
	}

	if params.Fields != "" {
		qs.Add("fields", params.Fields)
	}

	if params.Skip != "" {
		qs.Add("skip", params.Skip)
	}

	if params.Locale != "" {
		qs.Add("locale", params.Locale)
	}

	queryString := qs.Encode()
	if params.Fields != "" {
		queryString = strings.ReplaceAll(queryString,"fields=", "fields.")
		queryString = strings.ReplaceAll(queryString,"%3D", "=")
	}

	return queryString
}