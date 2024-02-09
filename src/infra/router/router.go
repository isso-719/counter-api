package router

import (
	"context"
	"github.com/isso-719/counter-api/src/adapter/handler"
	"github.com/isso-719/counter-api/src/infra/datastore"
	"github.com/isso-719/counter-api/src/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter(ctx context.Context) *echo.Echo {
	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORS(),
	)

	// init Datastore
	db, err := infra.CreateDatastoreClient(ctx)
	if err != nil {
		panic(err)
	}

	// routes
	counterRepository := infra.NewCounterRepository(db)
	counterService := usecase.NewCounterService(counterRepository)
	counterHandler := handler.NewCounterHandler(counterService)

	e.GET("/", counterHandler.IncrementCounter(ctx))

	return e
}
