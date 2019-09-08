package capi

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type EntriesCollection struct {
	Sys struct {
		Type string `json:"type"`
	} `json:"sys"`
	Total int `json:"total"`
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
	Items []Entry `json:"items"`
	Includes struct {
		Entry []Entry
		Asset []Asset
	}
}

type Entry struct {
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
		ContentType struct {
			Sys struct {
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
				ID       string `json:"id"`
			} `json:"sys"`
		} `json:"contentType"`
	} `json:"sys"`
	Fields map[string]interface{} `json:"fields"`
}

func NewEntry(client cmaClient) Entry {
	var e Entry
	e.client = client
	return e
}

func (e *Entry) List() EntriesCollection {
	ep := fmt.Sprintf("/spaces/%s/environments/%s/entries", e.client.SpaceID, e.client.Environment)
	resp := e.client.call("GET", ep, nil)

	var collection EntriesCollection
	body := readRequestBody(resp.Body)

	if err := json.Unmarshal(body, &collection); err != nil {
		log.Panic("Could not unmarshal fields in list entries", err)
	}

	return collection
}

func (e *Entry) Create(entryData, contentTypeID string) Entry {
	ep := fmt.Sprintf("/spaces/%s/environments/%s/entries", e.client.SpaceID, e.client.Environment)
	e.client.addHeader("X-Contentful-Content-Type", contentTypeID)

	resp := e.client.call("POST", ep, []byte(entryData))
	body := readRequestBody(resp.Body)

	var entry Entry
	if err := json.Unmarshal(body, &entry); err != nil {
		log.Panic("Could not unmarshal fields for creating a Content Type", err)
	}

	return entry
}

func (e *Entry) Publish(entryID string, version int) Entry {
	v := strconv.Itoa(version)
	ep := fmt.Sprintf("/spaces/%s/environments/%s/entries/%s/published", e.client.SpaceID, e.client.Environment, entryID)

	e.client.addHeader("X-Contentful-Version", v)
	resp := e.client.call("PUT", ep, nil)
	body := readRequestBody(resp.Body)

	var entry Entry
	if err := json.Unmarshal(body, &entry); err != nil {
		log.Panic("Could not unmarshal fields for publish Entry", err)
	}

	return entry
}
