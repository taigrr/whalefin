package main

import (
	"embed"
	_ "embed"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/taigrr/whalefin/xorg"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

func init() {
	runtime.LockOSThread()
}

//go:embed frontend/build
var assets embed.FS

func main() {
	var xPID int
	display := os.Getenv("DISPLAY")
	os.Setenv("XDG_SESSION_TYPE", "x11")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	if display == "" {
		os.Setenv("DISPLAY", ":0")
		xPID = xorg.StartX()
		defer xorg.StopX(xPID)
	}
	go func() {
		for {
			<-sig
			if xPID != 0 {
				xorg.StopX(xPID)
			}
			f.r.Window.Close()
		}
	}()

	// run blocking wails here
	width, height := getScreenResolution()

	app := NewApp()
	err := wails.Run(&options.App{
		Assets:     assets,
		Title:      "My App",
		Width:      800,
		Height:     600,
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
			GetFullScreen(),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	//	err := wails.Run(&options.App{
	//		Width:  int(width),
	//		Height: int(height),
	//		Title:  "whalefin",
	//		JS:     js,
	//		CSS:    css,
	//		Colour: "#0547b2",
	//	})
	//	app.Bind(GetFullScreen())
	//	app.Bind(login)
	//	app.Run()
}
