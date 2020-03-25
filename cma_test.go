package capi

import "testing"

func TestNewCMA(t *testing.T) {
	cma := NewCMA(CONFIG)

	if cma.Environment != CONFIG.Environment {
		t.Errorf("Environment does not match config, got %s, wanted %s", cma.Environment, CONFIG.Environment)
	}

	if cma.Api != "CMA" {
		t.Errorf("API does not match config, got %s, wanted %s", cma.Api, "CMA")
	}
}
