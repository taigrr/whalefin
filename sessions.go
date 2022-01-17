package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type xsession struct {
	Name string `json:"name"`
	Exec string `json:"exec"`
}

var xinitrc = xsession{
	Name: ".xinitrc",
	Exec: "/bin/bash --login .xinitrc",
}

func getSession(name string) string {
	sessions := loadSessions()
	for _, session := range sessions {
		if session.Name == name {
			return session.Exec
		}
	}
	return ""
}

// get a slice of all available xsessions
func loadSessions() (list []xsession) {
	list = []xsession{xinitrc}
	for _, dir := range getXDGDirs() {
		sessionDir := filepath.Join(dir, "xsessions")
		// check to see if the directory exists first
		files, err := ioutil.ReadDir(sessionDir)
		if err != nil {
			continue
		}
		for _, file := range files {
			desktopFile := filepath.Join(sessionDir, file.Name())
			session, err := parseXSession(desktopFile)
			if err != nil {
				fmt.Printf("xsession file %s is invalid: %v\n", file.Name(), err)
				continue
			}
			list = append(list, session)
		}
	}
	return list
}

// extracts the name and exec keys from a .desktop file
//
// returns an error if name or exec are missing, empty, or
// not inside of the [Desktop Entry] group header.
// Visit https://specifications.freedesktop.org/desktop-entry-spec/desktop-entry-spec-latest.html#example
// for more information about the Desktop Entry Specification
func parseXSession(path string) (xsession, error) {
	data := xsession{}
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Could not open xsession file: %v\n", err)
		return data, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var currentGroup string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "[") {
			currentGroup = line
		}
		if currentGroup != "[Desktop Entry]" {
			continue
		}
		// TODO: Load the proper locale-aware Name
		if strings.HasPrefix(line, "Name=") {
			name := strings.Split(line, "=")
			if len(name) == 1 {
				return data, errors.New("Desktop Entry Name field is empty")
			}
			data.Name = name[1]
		} else if strings.HasPrefix(line, "Exec=") {
			exec := strings.Split(line, "=")
			if len(exec) == 1 {
				return data, errors.New("Desktop Entry Exec field is empty")
			}
			data.Exec = exec[1]
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Could not read file: %v\n", err)
		return data, err
	}
	if data.Exec == "" {
		return data, errors.New("Desktop Entry Exec field is missing")
	}
	if data.Name == "" {
		return data, errors.New("Desktop Entry Name field is missing")
	}
	return data, nil
}

// Get an array of XDG data directory paths
//
// Assumes the uid has been set to the correct user already.
// Specification sourced from the official source at
// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
func getXDGDirs() (dirs []string) {
	// according to the spec, $XDG_DATA_HOME takes precedence
	// but if it's not set, it should be assumed to be $HOME/.local/share
	// Additionally, $XDG_DATA_DIRS contains system-wide data which should
	// be taken into account at a lower precedence. If empty, assume it is
	// "/usr/local/share/:/usr/share/"
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		home, err := os.UserHomeDir()
		if err == nil {
			dataHome = filepath.Join(home, ".local/share")
		}
	}
	dataDirs := os.Getenv("XDG_DATA_DIRS")
	if dataDirs == "" {
		dataDirs = "/usr/local/share/:/usr/share/"
	}
	dirs = []string{dataHome}
	dirs = append(dirs, strings.Split(dataDirs, ":")...)
	return
}
