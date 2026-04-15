package main

import (
	"testing"
)

func TestLoginEmptyUsername(t *testing.T) {
	handler := NewLoginHandler()
	err := handler.Login("", "pass", ".xinitrc")
	if err == nil {
		t.Error("expected error for empty username")
	}
}

func TestLoginEmptyPassword(t *testing.T) {
	handler := NewLoginHandler()
	err := handler.Login("testuser", "", ".xinitrc")
	if err == nil {
		t.Error("expected error for empty password")
	}
}

func TestLoginInvalidSession(t *testing.T) {
	handler := NewLoginHandler()
	err := handler.Login("testuser", "pass", "nonexistent-session")
	if err == nil {
		t.Error("expected error for invalid session")
	}
}
