package capi

import "time"

type ContentTypeCollection struct {
	Sys struct {
		Type string `json:"type"`
	} `json:"sys"`
	Total int           `json:"total"`
	Skip  int           `json:"skip"`
	Limit int           `json:"limit"`
	Items []ContentType `json:"items"`
}

type ContentType struct {
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
