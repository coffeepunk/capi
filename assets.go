package capi

import "time"

// Not sure this asset struct is correct.
// Fields could be a Fields map[string]interface{} `json:"fields"`
// Also not sure about envs, publish counter and so on.
type Asset struct {
	Sys    struct {
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
		PublishedVersion int       `json:"publishedVersion"`
		PublishedAt      time.Time `json:"publishedAt"`
		FirstPublishedAt time.Time `json:"firstPublishedAt"`
		CreatedBy        struct {
			Sys struct {
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
				ID       string `json:"id"`
			} `json:"sys"`
		} `json:"createdBy"`
		UpdatedBy struct {
			Sys struct {
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
				ID       string `json:"id"`
			} `json:"sys"`
		} `json:"updatedBy"`
		PublishedCounter int `json:"publishedCounter"`
		Version          int `json:"version"`
		PublishedBy      struct {
			Sys struct {
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
				ID       string `json:"id"`
			} `json:"sys"`
		} `json:"publishedBy"`
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
