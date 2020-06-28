package capi

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type EntriesCollection struct {
	Sys struct {
		Type string `json:"type"`
	} `json:"sys"`
	Total    int     `json:"total"`
	Skip     int     `json:"skip"`
	Limit    int     `json:"limit"`
	Items    []Entry `json:"items"`
	Includes struct {
		Entry []Entry
		Asset []Asset
	}
}

type Entry struct {
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
		ContentType struct {
			Sys struct {
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
				ID       string `json:"id"`
			} `json:"sys"`
		} `json:"contentType"`
		Locale string `json:"locale"`
	} `json:"sys"`
	Fields map[string]interface{} `json:"fields"`
}
func entriesEndpoint(spaceID, environment string) string {
	return fmt.Sprintf("/spaces/%s/environments/%s/entries", spaceID, environment)
}

func (cda *CDA) GetEntry(entryID string) Entry {
	ep := entriesEndpoint(cda.SpaceID, cda.Environment)
	ep = fmt.Sprintf("%s/%s", ep, entryID)
	resp, err := cda.client.call("GET", ep, nil)
	if err != nil {
		log.Println(err)
	}

	var entry Entry
	body := readRequestBody(resp.Body)

	if err := json.Unmarshal(body, &entry); err != nil {
		log.Panic("Could not unmarshal entry", err)
	}

	return entry
}

func (cda *CDA) GetEntries(params SearchParameters) (EntriesCollection, error) {
	var collection EntriesCollection
	ep := entriesEndpoint(cda.SpaceID, cda.Environment)
	qs := buildQueryString(params)

	ep = fmt.Sprintf("%s?%s", ep, qs)
	resp, err := cda.client.call("GET", ep, nil)
	if err != nil {
		return collection, err
	}

	body := readRequestBody(resp.Body)
	if err := json.Unmarshal(body, &collection); err != nil {
		log.Panic("Could not unmarshal entries in GetEntries", err)
	}

	return collection, nil
}
