package capi

import (
	"fmt"
	"testing"
)

func TestEntry_Create(t *testing.T) {
	entry := NewEntry(cma)
	postData := `{
	"fields": {
		"title": {
			"en": "About page for the great marketing company"
		},
		"url": {
			"en": "/about"
		},
		"layout": {
			"en": "resume"
		},
		"pageTitle": {
			"en": "The great marketing company"
		},
		"metaTitle": {
			"en": "The great marketing company"
		},
		"metaDescription": {
			"en": "A longer text that describes the marketing company, now this part is even longer"
		}
	}
}`
	result := entry.Create(postData, "1JqE691oom6ceTy2hKXTDR")

	fmt.Println(result.Sys.ID)
}

func TestEntry_Publish(t *testing.T) {
	entry := NewEntry(cma)
	result := entry.Publish("7KTVmrCc5ASlUabFs978S7", 1)

	fmt.Println(result)
}
