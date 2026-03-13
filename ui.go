package main

import (
	"github.com/wailsapp/wails"
)

var f FullScreen

func init() {
	f = FullScreen{}
}

// FullScreen manages the full-screen window state.
type FullScreen struct {
	r *wails.Runtime
}

func GetFullScreen() *FullScreen {
	return &f
}

// WailsInit is called when the component is being initialised
func (f *FullScreen) WailsInit(runtime *wails.Runtime) error {
	f.r = runtime
	f.r.Window.Fullscreen()
	return nil
}
