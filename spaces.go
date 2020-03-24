package capi

import "time"

type Space struct {
	Name   string `json:"name"`
	Sys    struct {
		Type      string    `json:"type"`
		ID        string    `json:"id"`
		Version   int       `json:"version"`
		CreatedBy *Sys      `json:"createdBy"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedBy *Sys      `json:"updatedBy"`
		UpdatedAt time.Time `json:"updatedAt"`
	} `json:"sys"`
}
