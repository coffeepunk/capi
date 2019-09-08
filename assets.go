package capi

import "time"

type Asset struct {
	Sys struct {
		Space struct {
			Sys struct {
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
				ID       string `json:"id"`
			} `json:"sys"`
		} `json:"space"`
		ID          string    `json:"id"`
		Type        string    `json:"type"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
		Environment struct {
			Sys struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
			} `json:"sys"`
		} `json:"environment"`
		Revision int    `json:"revision"`
		Locale   string `json:"locale"`
	} `json:"sys"`
	Fields struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		File        struct {
			URL     string `json:"url"`
			Details struct {
				Size  int `json:"size"`
				Image struct {
					Width  int `json:"width"`
					Height int `json:"height"`
				} `json:"image"`
			} `json:"details"`
			FileName    string `json:"fileName"`
			ContentType string `json:"contentType"`
		} `json:"file"`
	} `json:"fields"`
}

