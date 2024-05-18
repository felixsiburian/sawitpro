package router

import (
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/service"
	"github.com/labstack/echo/v4"
)

func NewRouter(
	e *echo.Echo,
	service service.ServiceInterface,
) {
	h := handler.NewServer(e, service)

	r := e.Group("")
	r.POST("/estate", h.CreateEstate)
	r.POST("/estate/:id/tree", h.CreateTree)
	r.PATCH("/:id/tree", h.UpdateTree)
	r.GET("/estate/:id/stats", h.TreeStatsByEstateId)
}
