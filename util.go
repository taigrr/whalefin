package main

import (
	"os"

	"github.com/BurntSushi/xgb/randr"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
)

func getScreenResolution() (w uint16, h uint16) {
	// Sensible defaults
	w, h = 1024, 768
	conn, err := xgbutil.NewConn()
	if err != nil {
		return
	}
	err = randr.Init(conn.Conn())
	if err != nil {
		return
	}
	root := xproto.Setup(conn.Conn()).DefaultScreen(conn.Conn()).Root
	resources, err := randr.GetScreenResources(conn.Conn(), root).Reply()
	if err != nil {
		return
	}
	if len(resources.Outputs) < 1 {
		return
	}
	output := resources.Outputs[0]
	reply, err := randr.GetOutputInfo(conn.Conn(), output, 0).Reply()
	if err != nil {
		return
	}
	if reply.NumModes < 1 {
		return
	}
	info, err := randr.GetCrtcInfo(conn.Conn(), reply.Crtc, 0).Reply()
	if err != nil {
		return
	}
	return info.Width, info.Height
}

func getHostname() string {
	host, err := os.Hostname()
	if err != nil {
		host = "localhost"
	}
	return host
}
