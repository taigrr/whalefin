package xorg

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const xServerReadyDelay = time.Second

var ErrInvalidPID = errors.New("invalid X server pid")

func StartX() (int, error) {
	exe := exec.Command("/usr/bin/X", ":0", "vt01")
	err := exe.Start()
	if err != nil {
		return 0, fmt.Errorf("start X server: %w", err)
	}
	// wait for process to spawn
	time.Sleep(xServerReadyDelay)
	return exe.Process.Pid, nil
}

func StopX(pid int) error {
	if pid <= 0 {
		return fmt.Errorf("%w: %d", ErrInvalidPID, pid)
	}

	process, err := findProcess(pid)
	if err != nil {
		return fmt.Errorf("find X server process %d: %w", pid, err)
	}
	if err := process.Kill(); err != nil {
		return fmt.Errorf("stop X server process %d: %w", pid, err)
	}
	return nil
}

var findProcess = func(pid int) (*os.Process, error) {
	return os.FindProcess(pid)
}
