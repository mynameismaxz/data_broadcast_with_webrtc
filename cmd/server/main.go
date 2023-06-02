package main

import (
	"mynameismaxz/data_broadcast_with_webrtc/pkg/signaling"
)

func main() {
	ss := signaling.NewWsServer()
	ss.Start()
}
