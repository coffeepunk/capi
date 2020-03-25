package capi

import "testing"

func TestNewCDA(t *testing.T) {
	cda := NewCDA(CONFIG)

	if cda.Environment != CONFIG.Environment {
		t.Errorf("Environment does not match config, got %s, wanted %s", cda.Environment, CONFIG.Environment)
	}

	if cda.Api != "CDA" {
		t.Errorf("API does not match config, got %s, wanted %s", cda.Api, "CDA")
	}
}
