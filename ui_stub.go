//go:build !desktop

package main

import "context"

// FullScreen is a stub for non-desktop builds.
type FullScreen struct {
	ctx context.Context
}

var f = FullScreen{}

var restartWails = make(chan bool, 1)

func GetFullScreen() *FullScreen {
	return &f
}

func (f *FullScreen) SetContext(ctx context.Context) {
	f.ctx = ctx
}

func (f *FullScreen) Quit() {}
