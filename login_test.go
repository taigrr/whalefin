package main

import (
	"testing"
)

func TestLoginEmptyUsername(t *testing.T) {
	data := map[string]interface{}{
		"Username": "",
		"Password": "pass",
		"XSession": ".xinitrc",
	}
	err := login(data)
	if err == nil {
		t.Error("expected error for empty username")
	}
}

func TestLoginEmptyPassword(t *testing.T) {
	data := map[string]interface{}{
		"Username": "testuser",
		"Password": "",
		"XSession": ".xinitrc",
	}
	err := login(data)
	if err == nil {
		t.Error("expected error for empty password")
	}
}

func TestLoginInvalidSession(t *testing.T) {
	data := map[string]interface{}{
		"Username": "testuser",
		"Password": "pass",
		"XSession": "nonexistent-session",
	}
	err := login(data)
	if err == nil {
		t.Error("expected error for invalid session")
	}
}

func TestLoginInvalidData(t *testing.T) {
	// Pass data that can't be decoded properly
	data := map[string]interface{}{
		"Username": 12345, // wrong type
	}
	err := login(data)
	if err == nil {
		t.Error("expected error for invalid data types")
	}
}
