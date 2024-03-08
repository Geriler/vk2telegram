package http

import (
	"github.com/labstack/echo"
	"net/http"
)

type Handler struct{}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello world"})
	})

	auth := e.Group("/auth")
	{
		auth.GET("/login", h.GetLogin)
		auth.POST("/login", h.PostLogin)
		auth.GET("/register", h.GetRegister)
		auth.POST("/register", h.PostRegister)
	}

	api := e.Group("/api")
	{
		vk := api.Group("/vk")
		{
			vk.GET("/group.search", h.GroupSearch)
			vk.GET("/wall.get", h.WallGet)
		}

		telegram := api.Group("/telegram")
		{
			telegram.POST("/send", h.Send)
		}
	}

	return e
}
