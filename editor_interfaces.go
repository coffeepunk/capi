package capi

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type EditorInterface struct {
	client cmaClient
	Sys struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Space struct {
			Sys struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
			} `json:"sys"`
		} `json:"space"`
		Version   int       `json:"version"`
		CreatedAt time.Time `json:"createdAt"`
		CreatedBy struct {
			Sys struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
			} `json:"sys"`
		} `json:"createdBy"`
		UpdatedAt time.Time `json:"updatedAt"`
		UpdatedBy struct {
			Sys struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
			} `json:"sys"`
		} `json:"updatedBy"`
		ContentType struct {
			Sys struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
			} `json:"sys"`
		} `json:"contentType"`
		Environment struct {
			Sys struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
			} `json:"sys"`
		} `json:"environment"`
	} `json:"sys"`
	Controls []EditorInterfaceControls `json:"controls"`
}

type EditorInterfaceControls struct {
	FieldID         string `json:"fieldId"`
	WidgetNamespace string `json:"widgetNamespace"`
	WidgetID        string `json:"widgetId"`
}

func NewEditorInterfaceControls(client cmaClient) EditorInterface {
	var ei EditorInterface
	ei.client = client
	return ei
}

func (ei *EditorInterface) Get(ContentTypeID string) EditorInterface {
	ep := fmt.Sprintf("/spaces/%s/environments/%s/content_types/%s/editor_interface", ei.client.SpaceID, ei.client.Environment, ContentTypeID)
	resp := ei.client.call("GET", ep, nil)

	var editorInterface EditorInterface
	body := readRequestBody(resp.Body)

	if err := json.Unmarshal(body, &editorInterface); err != nil {
		log.Panic("Could not unmarshal fields in get editor interface", err)
	}

	return editorInterface
}

func (ei *EditorInterface) Update(ContentTypeID string) {
	data := `{
       "controls": [
         {
           "fieldId": "url",
           "widgetNamespace": "builtin",
           "widgetId": "slugEditor"
         }
       ]
     }`

	ep := fmt.Sprintf("/spaces/%s/environments/%s/content_types/%s/editor_interface", ei.client.SpaceID, ei.client.Environment, ContentTypeID)
	ei.client.addHeader("X-Contentful-Version", "2")
	resp := ei.client.call("PUT", ep, []byte(data))
	body := readRequestBody(resp.Body)

	fmt.Println(string(body))
}
