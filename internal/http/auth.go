package http

import (
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) GetLogin(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Login page"})
}
func (h *Handler) PostLogin(ctx echo.Context) error {
	return nil
}

func (h *Handler) GetRegister(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Register page"})
}
func (h *Handler) PostRegister(ctx echo.Context) error {
	return nil
}
