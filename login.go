package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/taigrr/whalefin/pam"
)

// LoginHandler manages user authentication via PAM.
type LoginHandler struct {
	ctx context.Context
}

// NewLoginHandler creates a new LoginHandler instance.
func NewLoginHandler() *LoginHandler {
	return &LoginHandler{}
}

// SetContext stores the Wails runtime context.
func (l *LoginHandler) SetContext(ctx context.Context) {
	l.ctx = ctx
}

// Login authenticates a user with the given credentials and starts an X session.
func (l *LoginHandler) Login(username, password, xsession string) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}
	if password == "" {
		return errors.New("password cannot be empty")
	}
	xsession = getSession(xsession)
	if xsession == "" {
		return errors.New("invalid xsession provided")
	}
	go func() {
		pid, err := pam.Login(username, password, xsession)
		if err != nil {
			fmt.Printf("Error logging in %s: %v\n", username, err)
			return
		}
		proc, err := os.FindProcess(pid)
		if err != nil {
			fmt.Printf("Error launching login process: %v\n", err)
			return
		}

		// close the wails window; the xsession takes over
		GetFullScreen().Quit()
		// wait for the xsession to exit
		proc.Wait()

		// TODO: check to make sure we don't need to reset the setUID command
		// to be able to log in as a different user later
		pam.Logout()
		// signal main to exit (systemd will restart the display manager)
		restartWails <- true
	}()
	return nil
}
