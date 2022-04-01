package capi

import (
	"fmt"
)

type CDA struct {
	Api             string
	EntriesEndPoint string
	Client
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
	cda.EntriesEndPoint = cda.endPoint("entries")

	cda.Client.AddHeader("Authorization", bearer)
	return cda
}

func (cda *CDA) endPoint(name string) string {
	return fmt.Sprintf("/spaces/%s/environments/%s/%s", cda.SpaceID, cda.Environment, name)
}
