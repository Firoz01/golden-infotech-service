package v1

import (
	"github.com/golden-infotech/config"
	"github.com/golden-infotech/lib/logger"
	"github.com/golden-infotech/repository"
	"github.com/golden-infotech/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
)

// Setup all routers
func SetupRouters(c *echo.Echo, conf *config.Config, db *bun.DB, jwtConfig middleware.JWTConfig, logger logger.Logger) {

	booksRepository := repository.NewBooksRepository(db)
	booksService := service.NewBooksService(booksRepository)
	booksHandler := NewBooksHandler(booksService, logger)

	authenticated := middleware.JWTWithConfig(jwtConfig)

	v1 := c.Group("/api/v1")

	booksGroup := v1.Group("/books")

	booksHandler.MapBooksRoutes(booksGroup, authenticated)
}
