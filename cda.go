package capi

import "fmt"

type CDA struct {
	Api string
	client
}

func NewCDA(config Config) CDA {
	bearer := fmt.Sprintf("Bearer %s", config.AccessToken)
	var cda CDA
	cda.SpaceID = config.SpaceID
	cda.AccessToken = config.AccessToken
	cda.Environment = config.Environment
	cda.OrganisationID = config.OrganisationID
	cda.baseUrl = "https://cdn.contentful.com"
	cda.Api = "CDA"

	cda.client.addHeader("Authorization", bearer)
	return cda
}
