package signaling

import (
	"github.com/labstack/echo/v4"
	"github.com/olahol/melody"
)

type WsServer struct {
	mSession   *melody.Melody
	echoClient *echo.Echo
}

func NewWsServer() *WsServer {
	m := melody.New()
	e := echo.New()

	s := WsServer{
		mSession:   m,
		echoClient: e,
	}
	s.initiateRoute()

	return &s
}

func (s *WsServer) initiateRoute() {
	// initial route for websocket
	s.echoClient.GET("/ws", func(c echo.Context) error {
		s.mSession.HandleRequest(c.Response().Writer, c.Request())
		return nil
	})

	s.mSession.HandleMessage(func(ms *melody.Session, b []byte) {
		s.mSession.Broadcast(b)
	})
}

func (s *WsServer) Start() {
	s.echoClient.Logger.Fatal(s.echoClient.Start(":10020"))
}
