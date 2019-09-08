package capi

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Asset struct {
	client cmaClient
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
		Title struct{
			En string `json:"en"`
		} `json:"title"`
		Description string `json:"description"`
		File struct {
			En struct {
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
			} `json:"en"`
		} `json:"file"`
	} `json:"fields"`
}

func NewAsset(client cmaClient) Asset {
	var a Asset
	a.client = client
	return a
}

func (a *Asset) Create(assetData string) Asset {
	ep := fmt.Sprintf("/spaces/%s/environments/%s/assets", a.client.SpaceID, a.client.Environment)

	resp := a.client.call("POST", ep, []byte(assetData))
	body := readRequestBody(resp.Body)

	var asset Asset
	if err := json.Unmarshal(body, &asset); err != nil {
		log.Panic("Could not unmarshal asset in Assets create", err)
	}

	return asset
}

func (a *Asset) Process(assetID, locale string, version int) int {
	v := strconv.Itoa(version)
	ep := fmt.Sprintf("/spaces/%s/environments/%s/assets/%s/files/%s/process", a.client.SpaceID, a.client.Environment, assetID, locale)

	a.client.addHeader("X-Contentful-Version", v)
	resp := a.client.call("PUT", ep, nil)

	return resp.StatusCode
}

func (a *Asset) Publish(assetID string, version int) Asset {
	v := strconv.Itoa(version)
	ep := fmt.Sprintf("/spaces/%s/environments/%s/assets/%s/published", a.client.SpaceID, a.client.Environment, assetID)

	a.client.addHeader("X-Contentful-Version", v)
	resp := a.client.call("PUT", ep, nil)
	body := readRequestBody(resp.Body)

	var asset Asset
	if err := json.Unmarshal(body, &asset); err != nil {
		log.Panic("Could not unmarshal asset in Assets publish", err)
	}

	return asset
}
