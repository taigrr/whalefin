package main

import (
	"testing"
)

func TestGetHostname(t *testing.T) {
	host := getHostname()
	if host == "" {
		t.Error("expected non-empty hostname")
	}
}
