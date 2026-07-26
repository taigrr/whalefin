//go:build desktop

package main

import (
	"embed"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/taigrr/whalefin/xorg"
	"github.com/wailsapp/wails/v3/pkg/application"
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
	fullscreen := GetFullScreen()
	loginHandler := NewLoginHandler()

	app := application.New(application.Options{
		Name: "whalefin",
		Services: []application.Service{
			application.NewService(fullscreen),
			application.NewService(loginHandler),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
	})

	window := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:            "whalefin",
		Width:            int(width),
		Height:           int(height),
		BackgroundColour: application.NewRGB(5, 71, 178),
	})
	window.Fullscreen()

	fullscreen.SetApp(app)

	err := app.Run()
	if err != nil {
		panic(err)
	}
	<-restartWails
}
