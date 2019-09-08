package capi

import (
	"fmt"
	"testing"
)

func TestContentType_List(t *testing.T) {
	contentType := NewContentType(cma)
	result := contentType.List()

	for _, ct := range result.Items {
		fmt.Println(ct.Name)
		fmt.Println("----------------")
	}
}

func TestContentType_Create(t *testing.T) {
	contentType := NewContentType(cma)
	fmt.Println("-----------------------------------")
	fmt.Println(cma.AccessToken)
	postData := `{
  "name": "Page",
  "description": "Page assembler for web pages",
  "displayField": "title",
  "fields": [
    {
      "id": "title",
      "name": "Title",
      "type": "Symbol",
      "localized": false,
      "required": false,
      "validations": [],
      "disabled": false,
      "omitted": false
    },
    {
      "id": "url",
      "name": "url",
      "type": "Symbol",
      "localized": true,
      "required": true,
      "validations": [
        {
          "unique": true
        }
      ],
      "disabled": false,
      "omitted": false
    },
    {
      "id": "layout",
      "name": "layout",
      "type": "Symbol",
      "localized": true,
      "required": true,
      "validations": [
        {
          "in": [
            "single",
            "resume",
            "blog"
          ]
        }
      ],
      "disabled": false,
      "omitted": false
    },
    {
      "id": "pageTitle",
      "name": "page title",
      "type": "Symbol",
      "localized": true,
      "required": false,
      "validations": [],
      "disabled": false,
      "omitted": false
    },
    {
      "id": "metaTitle",
      "name": "Meta Title",
      "type": "Symbol",
      "localized": true,
      "required": true,
      "validations": [
        {
          "size": {
            "max": 152
          }
        }
      ],
      "disabled": false,
      "omitted": false
    },
    {
      "id": "metaDescription",
      "name": "Meta Description",
      "type": "Symbol",
      "localized": true,
      "required": false,
      "validations": [],
      "disabled": false,
      "omitted": false
    },
    {
      "id": "content",
      "name": "content",
      "type": "Array",
      "localized": false,
      "required": false,
      "validations": [],
      "disabled": false,
      "omitted": false,
      "items": {
        "type": "Link",
        "validations": [
          {
            "linkContentType": [
              "photo",
              "textBlock"
            ]
          }
        ],
        "linkType": "Entry"
      }
    }
  ]
}`
	result := contentType.Create(postData)

	fmt.Println(result)
}

func TestContentType_Activate(t *testing.T) {
	contentType := NewContentType(cma)
	result := contentType.Activate("6YD3RiRDtrBpUey7D77LCB", 1)

	fmt.Println(result)
}
