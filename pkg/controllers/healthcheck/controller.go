package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const Route = "/healthcheck"

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

type Response struct {
	Message string `json:"message"`
}

func (c *Controller) Handler(e echo.Context) error {
	resp := &Response{
		Message: "success",
	}
	return e.JSON(http.StatusOK, resp)
}
