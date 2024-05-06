package route

import (
	"recything/app/database"
	"recything/features/recybot/handler"
	"recything/features/recybot/repository"
	"recything/features/recybot/service"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RouteRecybot(e *echo.Group, db *gorm.DB,rdb *redis.Client) {
	redisDB := database.NewRedis(rdb)
	recybotRepository := repository.NewRecybotRepository(db,redisDB)
	recybotService := service.NewRecybotService(recybotRepository)
	recybotHandler := handler.NewRecybotHandler(recybotService)

	e.POST("",recybotHandler.RecyBotChat,jwt.JWTMiddleware())
}
