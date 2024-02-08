package handler

import (
	"context"
	"github.com/isso-719/counter-api/src/domain"
	"github.com/isso-719/counter-api/src/usecase"
	"github.com/labstack/echo/v4"
)

type IFCounterHandler interface {
	IncrementCounter(ctx context.Context) echo.HandlerFunc
}

type CounterHandler struct {
	counterService usecase.IFCounterService
}

func NewCounterHandler(counterService usecase.IFCounterService) IFCounterHandler {
	return &CounterHandler{
		counterService: counterService,
	}
}

type IncrementCounterRequest struct {
	URL string `json:"url"`
}

type IncrementCounterResponse struct {
	Counter domain.Counter `json:"counter"`
}

type IncrementCounterErrorResponse struct {
	Error string `json:"error"`
}

func (c *CounterHandler) IncrementCounter(ctx context.Context) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var req IncrementCounterRequest
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(400, IncrementCounterErrorResponse{
				Error: err.Error(),
			})
		}

		url := ctx.QueryParam("url")
		count, err := c.counterService.Increment(ctx.Request().Context(), url)
		if err != nil {
			return ctx.JSON(400, IncrementCounterErrorResponse{
				Error: err.Error(),
			})
		}

		return ctx.JSON(200, IncrementCounterResponse{
			Counter: *count,
		})
	}
}
