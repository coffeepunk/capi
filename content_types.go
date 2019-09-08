package capi

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type ContentTypeCollection struct {
	Sys struct {
		Type string `json:"type"`
	} `json:"sys"`
	Total int `json:"total"`
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
	Items []ContentType `json:"items"`
}

type ContentType struct {
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
	DisplayField string `json:"displayField"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Fields       []struct {
		ID          string        `json:"id"`
		Name        string        `json:"name"`
		Type        string        `json:"type"`
		Localized   bool          `json:"localized"`
		Required    bool          `json:"required"`
		Validations []interface{} `json:"validations"`
		Disabled    bool          `json:"disabled"`
		Omitted     bool          `json:"omitted"`
		LinkType    string        `json:"linkType,omitempty"`
	} `json:"fields"`
}

type ContentTypeModel struct {
	Name   string `json:"name"`
	Fields []struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Required  bool   `json:"required"`
		Localized bool   `json:"localized"`
		Type      string `json:"type"`
	} `json:"fields"`
}

func NewContentType(client cmaClient) ContentType {
	var ct ContentType
	ct.client = client
	return ct
}

func (ct *ContentType) List() ContentTypeCollection {
	ep := fmt.Sprintf("/spaces/%s/environments/%s/content_types", ct.client.SpaceID, ct.client.Environment)
	resp := ct.client.call("GET", ep, nil)

	var collection ContentTypeCollection
	body := readRequestBody(resp.Body)

	if err := json.Unmarshal(body, &collection); err != nil {
		log.Panic("Could not unmarshal fields in list content types", err)
	}

	return collection
}

func (ct *ContentType) Create(contentData string) ContentType {
	ep := fmt.Sprintf("/spaces/%s/environments/%s/content_types", ct.client.SpaceID, ct.client.Environment)

	resp := ct.client.call("POST", ep, []byte(contentData))
	body := readRequestBody(resp.Body)

	var contentType ContentType
	if err := json.Unmarshal(body, &contentType); err != nil {
		log.Panic("Could not unmarshal fields for creating a Content Type", err)
	}

	return contentType
}

func (ct *ContentType) Activate(contentID string, version int) ContentType {
	v := strconv.Itoa(version)
	ep := fmt.Sprintf("/spaces/%s/environments/%s/content_types/%s/published", ct.client.SpaceID, ct.client.Environment, contentID)

	ct.client.addHeader("X-Contentful-Version", v)
	resp := ct.client.call("PUT", ep, nil)
	body := readRequestBody(resp.Body)

	var contentType ContentType
	if err := json.Unmarshal(body, &contentType); err != nil {
		log.Panic("Could not unmarshal fields for creating a Content Type", err)
	}

	return contentType
}
