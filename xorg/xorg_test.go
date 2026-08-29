package xorg

import (
	"errors"
	"os"
	"testing"
)

func TestStopXRejectsInvalidPID(t *testing.T) {
	err := StopX(0)
	if !errors.Is(err, ErrInvalidPID) {
		t.Fatalf("got %v, want ErrInvalidPID", err)
	}
}

func TestStopXReturnsFindProcessError(t *testing.T) {
	wantErr := errors.New("lookup failed")
	originalFindProcess := findProcess
	findProcess = func(pid int) (*os.Process, error) {
		if pid != 1234 {
			t.Fatalf("got pid %d, want 1234", pid)
		}
		return nil, wantErr
	}
	defer func() {
		findProcess = originalFindProcess
	}()

	err := StopX(1234)
	if !errors.Is(err, wantErr) {
		t.Fatalf("got %v, want %v", err, wantErr)
	}
}
