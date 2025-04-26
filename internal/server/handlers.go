package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	eventHttp "github.com/raflidev/golang-clean-architecture/internal/event/delivery/http"
	"github.com/raflidev/golang-clean-architecture/internal/event/repository"
	"github.com/raflidev/golang-clean-architecture/internal/event/usecase"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	v1 := e.Group("/api/v1")
	product := v1.Group("/product")
	events := v1.Group("/events")

	eventRepo := repository.NewEventRepository(s.db)

	eventUC := usecase.NewEventUseCase(eventRepo)

	eventHandler := eventHttp.NewEventHandler(eventUC)

	eventHttp.MapRoutes(events, eventHandler)

	product.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Message string `json:"message"`
			Status  int    `json:"status"`
		}{
			Message: "product",
			Status:  http.StatusOK,
		})
	})

	return nil
}
