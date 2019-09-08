package capi

import (
	"fmt"
	"testing"
)

func TestAsset_Create(t *testing.T) {
	asset := NewAsset(cma)
	postData := `{
	"fields": {
		"title": {
			"en": "Same tree different name"
		},
		"file": {
			"en": {
				"contentType": "image/jpeg",
				"fileName": "tree-1234.jpeg",
				"upload": ""
			}
		}
	}
}`
	result := asset.Create(postData)

	fmt.Println(result.Sys.ID)
}

func TestAsset_Process(t *testing.T) {
	asset := NewAsset(cma)
	result := asset.Process("", "en", 1)

	fmt.Println(result)
}

func TestAsset_Publish(t *testing.T) {
	asset := NewAsset(cma)
	result := asset.Publish("", 2)

	fmt.Println(result)
}
