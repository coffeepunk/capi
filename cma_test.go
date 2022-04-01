package capi

import "testing"

func TestNewCMA(t *testing.T) {
	cma := NewCMA(TESTCONFIG)

	if cma.Environment != TESTCONFIG.Environment {
		t.Errorf("Environment does not match config, got %s, wanted %s", cma.Environment, TESTCONFIG.Environment)
	}

	if cma.Api != "CMA" {
		t.Errorf("API does not match config, got %s, wanted %s", cma.Api, "CMA")
	}
}
