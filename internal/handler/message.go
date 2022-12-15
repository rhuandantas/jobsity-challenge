package handler

import (
	"chat-jobsity/internal/command"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{}
)

type MessageHandler struct {
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}

func (h *MessageHandler) Hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}

func (h *MessageHandler) HandleCommand(c echo.Context) error {
	msg := new(command.Message)
	if err := c.Bind(msg); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	cmdRunner, err := command.GetCommand(msg.Command)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	message, err := cmdRunner.Run(msg.Value)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, message)
}
