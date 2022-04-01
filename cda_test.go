package capi

import (
	"fmt"
	"testing"
)

/*
func TestNewCDA(t *testing.T) {
	cda := NewCDA(CONFIG)

	if cda.Environment != CONFIG.Environment {
		t.Errorf("Environment does not match config, got %s, wanted %s", cda.Environment, CONFIG.Environment)
	}

	if cda.Api != "CDA" {
		t.Errorf("API does not match config, got %s, wanted %s", cda.Api, "CDA")
	}
}
*/

func TestEntriesEndpoint(t *testing.T) {
	ep := fmt.Sprintf("/spaces/%s/environments/%s/entries", TESTCONFIG.SpaceID, TESTCONFIG.Environment)

	cda := NewCDA(TESTCONFIG)
	if cda.EntriesEndPoint != ep {
		t.Errorf("Endpoint does not match, got %s, wanted %s", cda.EntriesEndPoint, ep)
	}
}
