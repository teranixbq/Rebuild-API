package route

import (
	"recything/features/achievement/handler"
	"recything/features/achievement/repository"
	"recything/features/achievement/service"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteAchievement(e *echo.Group, db *gorm.DB) {
	achievementRepository := repository.NewAchievementRepository(db)
	achievementService := service.NewAchievementService(achievementRepository)
	achievementHandler := handler.NewAchievementHandler(achievementService)

	admin := e.Group("/admins/manage/achievements", jwt.JWTMiddleware())
	admin.GET("", achievementHandler.GetAllAchievement)
	admin.PATCH("/:id", achievementHandler.UpdateById)

	user := e.Group("/achievements", jwt.JWTMiddleware())
	user.GET("", achievementHandler.GetAllAchievement)
}
