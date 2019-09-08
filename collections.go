package capi

type Collection struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
	Sys struct {
		Type string `json:"type"`
	} `json:"sys"`
	Items []struct {
		*Space
		*ContentType
	} `json:"items"`
}
