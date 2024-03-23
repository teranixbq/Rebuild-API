package route

import (
	"recything/features/drop-point/handler"
	"recything/features/drop-point/repository"
	"recything/features/drop-point/service"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteDropPoint(e *echo.Group, db *gorm.DB) {

	dropPointRepository := repository.NewDropPointRepository(db)
	dropPointService := service.NewDropPointService(dropPointRepository)
	dropPointHandler := handler.NewDropPointHandler(dropPointService)

	admin := e.Group("/admins/manage/drop-points", jwt.JWTMiddleware())
	admin.POST("", dropPointHandler.CreateDropPoint)
	admin.GET("", dropPointHandler.GetAllDropPoint)
	admin.GET("/:id", dropPointHandler.GetDropPointById)
	admin.PUT("/:id", dropPointHandler.UpdateDropPoint)
	admin.DELETE("/:id", dropPointHandler.DeleteDropPoint)

	user := e.Group("/drop-points", jwt.JWTMiddleware())
	user.GET("", dropPointHandler.GetAllDropPoint)
	user.GET("/:id", dropPointHandler.GetDropPointById)
}