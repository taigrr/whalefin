package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/taigrr/whalefin/pam"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	XSession string `json:"xsession"`
}
type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.Show()
}
func (a *App) shutdown(ctx context.Context) {
}

func (a *App) Close() {
	runtime.Quit(a.ctx)
}
func (a *App) Show() {
	runtime.WindowShow(a.ctx)
	runtime.WindowFullscreen(a.ctx)
}
func (a *App) Hide() {
	runtime.WindowHide(a.ctx)
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

func (a *App) Login(username, password, xsession string) error {
	if username == "" {
		return errors.New("Username cannot be empty!")
	}
	if password == "" {
		return errors.New("Password cannot be empty!")
	}
	xsession = getSession(xsession)
	if xsession == "" {
		return errors.New("Invalid XSession provided!")
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
			// make sure wails is visible
			a.Show()
			return
		}

		// hide wails
		a.Hide()
		// xsession has begun, wait for it to exit before taking over
		proc.Wait()

		// TODO: check to make sure we don't need to reset the setUID command
		// to be able to log in as a different user later
		pam.Logout()
		// show wails
		a.Show()
	}()
	return nil
}
