package routes

import (
	"foodmarketplace/handlers"
	"github.com/labstack/echo/v4"
)

func SetRoutes(e *echo.Echo) {
	userGroup := e.Group("/api/v1/user/")
	userGroup.GET("list", handlers.List)
	userGroup.POST("create", handlers.Create)
	userGroup.PUT("update/:id", handlers.Update)
	userGroup.DELETE("delete/:id", handlers.Delete)
}
