package main

import (
	"context"
	"embed"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/taigrr/whalefin/xorg"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

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
		signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
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

	width, height := getScreenResolution()
	loginHandler := NewLoginHandler()
	fullscreen := GetFullScreen()

	err := wails.Run(&options.App{
		Title:  "whalefin",
		Width:  int(width),
		Height: int(height),
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 5, G: 71, B: 178, A: 255},
		OnStartup: func(ctx context.Context) {
			fullscreen.SetContext(ctx)
			loginHandler.SetContext(ctx)
		},
		Bind: []interface{}{
			fullscreen,
			loginHandler,
		},
	})
	if err != nil {
		panic(err)
	}
	<-restartWails
}
