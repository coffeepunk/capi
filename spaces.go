package capi

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Space struct {
	client cmaClient
	Name   string `json:"name"`
	Sys    struct {
		Type      string `json:"type"`
		ID        string `json:"id"`
		Version   int    `json:"version"`
		CreatedBy *Sys `json:"createdBy"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedBy *Sys `json:"updatedBy"`
		UpdatedAt time.Time `json:"updatedAt"`
	} `json:"sys"`
}

func NewSpace(client cmaClient) Space {
	var s Space
	s.client = client
	return s
}

func (s *Space) List() Collection {
	resp := s.client.call("GET", "/spaces", nil)

	var collection Collection
	body := readRequestBody(resp.Body)

	if err := json.Unmarshal(body, &collection); err != nil {
		log.Panic("Could not unmarshal fields in buildPost", err)
	}

	return collection
}

func (s *Space) Create(name, locale string) Space {
	postData := fmt.Sprintf(`{"name": "%s", "defaultLocale": "%s"}`, name, locale)
	resp := s.client.call("POST", "/spaces", []byte(postData))
	body := readRequestBody(resp.Body)

	var space Space
	if err := json.Unmarshal(body, &space); err != nil {
		log.Panic("Could not unmarshal fields in buildPost", err)
	}

	return space
}

func (s *Space) Delete(ID string) int {
	endpoint := fmt.Sprintf("/spaces/%s", ID)
	resp := s.client.call("DELETE", endpoint, nil)

	return resp.StatusCode
}
