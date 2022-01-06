package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/mitchellh/mapstructure"
	"github.com/taigrr/whalefin/pam"
)

// Counter is what we use for counting
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	XSession string `json:"xsession"`
}

func login(data map[string]interface{}) error {
	var login Login
	err := mapstructure.Decode(data, &login)
	if err != nil {
		return err
	}
	if login.Username == "" {
		return errors.New("Username cannot be empty!")
	}
	if login.Password == "" {
		return errors.New("Password cannot be empty!")
	}
	login.XSession = getSession(login.XSession)
	if login.XSession == "" {
		return errors.New("Invalid xsession provided!")
	}
	go func() {
		pid, err := pam.Login(login.Username, login.Password, login.XSession)
		if err != nil {
			fmt.Printf("Error logging in %s: %v\n", login.Username, err)
			return
		}
		proc, err := os.FindProcess(pid)
		if err != nil {
			fmt.Printf("Error launching login process: %v\n", err)
			// make sure wails is visible
			return
		}

		// hide wails
		GetFullScreen().r.Window.Close()
		// xsession has begun, wait for it to exit before taking over
		proc.Wait()

		// TODO: check to make sure we don't need to reset the setUID command
		// to be able to log in as a different user later
		pam.Logout()
		// show wails
		restartWails <- true
	}()
	return nil
}
