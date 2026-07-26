//go:build desktop

package main

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

var f FullScreen

func init() {
	f = FullScreen{}
}

// FullScreen manages the full-screen window state.
type FullScreen struct {
	app *application.App
}

func GetFullScreen() *FullScreen {
	return &f
}

// SetApp stores a reference to the running Wails application.
func (f *FullScreen) SetApp(app *application.App) {
	f.app = app
}

// Quit closes the Wails application.
func (f *FullScreen) Quit() {
	if f.app != nil {
		f.app.Quit()
	}
}
