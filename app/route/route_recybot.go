package route

import (
	"recything/features/recybot/handler"
	"recything/features/recybot/repository"
	"recything/features/recybot/service"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteRecybot(e *echo.Group, db *gorm.DB) {
	recybotRepository := repository.NewRecybotRepository(db)
	recybotService := service.NewRecybotService(recybotRepository)
	recybotHandler := handler.NewRecybotHandler(recybotService)

	e.POST("",recybotHandler.RecyBotChat,jwt.JWTMiddleware())
}
