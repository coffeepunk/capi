package capi

import "fmt"

type CMA struct {
	Api string
	client
}

func NewCMA(config Config) CMA {
	bearer := fmt.Sprintf("Bearer %s", config.AccessToken)
	var cma CMA
	cma.SpaceID = config.SpaceID
	cma.AccessToken = config.AccessToken
	cma.Environment = config.Environment
	cma.OrganisationID = config.OrganisationID
	cma.baseUrl = "https://api.contentful.com"
	cma.Api = "CMA"

	cma.client.addHeader("Authorization", bearer)
	cma.addHeader("X-Contentful-Organization", config.OrganisationID)
	cma.addHeader("Content-Type", "application/vnd.contentful.management.v1+json")
	return cma
}
