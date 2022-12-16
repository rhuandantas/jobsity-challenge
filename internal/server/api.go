package server

import (
	"chat-jobsity/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/olahol/melody"
)

type API struct {
	Server         *echo.Echo
	Melody         *melody.Melody
	messageHandler *handler.MessageHandler
}

func NewAPI(messageHandler *handler.MessageHandler) *API {
	api := &API{
		Server:         CreateServer(),
		messageHandler: messageHandler,
	}

	m := melody.New()
	api.Melody = m

	api.setRoutes()

	return api
}

func CreateServer() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Static("/", "../public")
	e.File("/", "public/index.html")
	e.File("/channel/:channel_id", "public/chan.html")

	return e
}

func (a *API) setRoutes() {
	a.Server.GET("/ws", func(c echo.Context) error {
		a.Melody.HandleRequest(c.Response().Writer, c.Request())
		return nil
	})

	a.Melody.HandleMessage(func(s *melody.Session, msg []byte) {

		message, err := a.messageHandler.HandleMessage(msg)
		if err != nil {
			msg = []byte(err.Error())
		}

		if message != "" {
			msg = []byte(message)
		}

		a.Melody.BroadcastFilter(msg, func(q *melody.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})
}
