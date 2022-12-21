package server

import (
	"chat-jobsity/internal/handler"
	"chat-jobsity/internal/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/olahol/melody"
)

type API struct {
	Server         *echo.Echo
	Melody         *melody.Melody
	messageHandler handler.MessageHandler
}

func NewAPI(messageHandler handler.MessageHandler) *API {
	api := &API{
		Server:         CreateServer(),
		Melody:         melody.New(),
		messageHandler: messageHandler,
	}

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

	return e
}

func (a *API) setRoutes() {
	a.Server.GET("/ws", func(c echo.Context) error {
		a.Melody.HandleRequest(c.Response().Writer, c.Request())
		return nil
	})
	var response models.MessageResponse

	a.Melody.HandleMessage(func(s *melody.Session, msg []byte) {
		message, err := a.messageHandler.HandleMessage(msg)
		if err != nil {
			msgErr := err.Error()
			response = models.MessageResponse{
				Message: nil,
				Error:   &msgErr,
			}
		} else {
			response = models.MessageResponse{
				Message: &message,
				Error:   nil,
			}
		}

		bytes, err := json.Marshal(response)

		a.Melody.Broadcast(bytes)
	})
}
