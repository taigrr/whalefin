package main

import (
	_ "embed"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/taigrr/whalefin/xorg"
	"github.com/wailsapp/wails"
)

//go:embed frontend/build/main.js
var js string

//go:embed frontend/build/main.css
var css string

var restartWails chan bool

func init() {
	runtime.LockOSThread()
	restartWails = make(chan bool)
}

func main() {
	var xPID int
	display := os.Getenv("DISPLAY")
	os.Setenv("XDG_SESSION_TYPE", "x11")
	if display == "" {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGTERM)
		go func() {
			for {
				<-sig
				xorg.StopX(xPID)
			}
		}()
		xPID = xorg.StartX()
		defer xorg.StopX(xPID)
		os.Setenv("DISPLAY", ":0")
	}

	// run blocking wails here
	width, height := getScreenResolution()
	for {
		app := wails.CreateApp(&wails.AppConfig{
			Width:  int(width),
			Height: int(height),
			Title:  "whalefin",
			JS:     js,
			CSS:    css,
			Colour: "#0547b2",
		})
		app.Bind(GetFullScreen())
		app.Run()
		<-restartWails
	}
}
