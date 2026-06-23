package main

import (
	"os"
)

func getHostname() string {
	host, err := os.Hostname()
	if err != nil {
		host = "localhost"
	}
	return host
}
