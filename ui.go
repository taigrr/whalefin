package main

import (
	"context"

	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

var f FullScreen

func init() {
	f = FullScreen{}
}

// FullScreen manages the full-screen window state.
type FullScreen struct {
	ctx context.Context
}

func GetFullScreen() *FullScreen {
	return &f
}

// SetContext stores the Wails runtime context and applies fullscreen mode.
func (f *FullScreen) SetContext(ctx context.Context) {
	f.ctx = ctx
	wailsRuntime.WindowFullscreen(ctx)
}

// Quit closes the Wails application window.
func (f *FullScreen) Quit() {
	wailsRuntime.Quit(f.ctx)
}
