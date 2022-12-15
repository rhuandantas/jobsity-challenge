package server

import (
	"chat-jobsity/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"net/http"
	"text/template"
)

type API struct {
	Server         *echo.Echo
	messageHandler *handler.MessageHandler
}

func NewAPI(messageHandler *handler.MessageHandler) *API {
	api := &API{
		Server:         CreateServer(),
		messageHandler: messageHandler,
	}

	api.setRoutes()

	return api
}

func CreateServer() *echo.Echo {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "../public")
	e.File("/", "public/index.html")

	return e
}

func (a *API) setRoutes() {
	a.Server.POST("/command", a.messageHandler.HandleCommand)

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	a.Server.Renderer = renderer
	// Named route "foobar"
	a.Server.GET("/something", func(c echo.Context) error {
		return c.Render(http.StatusOK, "template.html", map[string]interface{}{
			"name": "Dolly!",
		})
	}).Name = "foobar"
	a.Server.GET("/ws", a.messageHandler.Hello)

}

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
