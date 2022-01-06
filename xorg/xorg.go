package xorg

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func StartX() int {
	cmd := "/usr/bin/X :0 vt01"
	exe := exec.Command("/bin/bash", "-c", cmd)
	err := exe.Start()
	if err != nil {
		log.Printf("Could not start X server: %v\n", err)
		os.Exit(1)
	}
	// wait for process to spawn
	time.Sleep(time.Second)
	return exe.Process.Pid
}

func StopX(pid int) {
	p, err := os.FindProcess(pid)
	if err != nil {
		log.Printf("Could not find X server pid: %v\n", err)
	}
	p.Kill()
}
