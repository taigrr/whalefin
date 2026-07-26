//go:build !desktop

package main

// FullScreen is a stub for non-desktop builds.
type FullScreen struct{}

var f = FullScreen{}

var restartWails = make(chan bool, 1)

func GetFullScreen() *FullScreen {
	return &f
}

func (f *FullScreen) Quit() {}
