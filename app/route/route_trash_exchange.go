package route

import (
	"recything/features/trash_exchange/handler"
	"recything/features/trash_exchange/repository"
	"recything/features/trash_exchange/service"

	userRepo "recything/features/user/repository"
	dropPointRepo "recything/features/drop-point/repository"
	trashCategoryRepo "recything/features/trash_category/repository"
	achievement"recything/features/achievement/repository"

	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteTrashExchange(e *echo.Group, db *gorm.DB) {
	achievementRepository := achievement.NewAchievementRepository(db)
	userRepository := userRepo.NewUserRepository(db,achievementRepository)
	dropPointRepository := dropPointRepo.NewDropPointRepository(db)
	trashCategoryRepository := trashCategoryRepo.NewTrashCategoryRepository(db)

	trashExchangeRepository := repository.NewTrashExchangeRepository(db)
	trashExchangeService := service.NewTrashExchangeService(trashExchangeRepository, dropPointRepository, userRepository, trashCategoryRepository)
	trashExchangeHandler := handler.NewTrashExchangeHandler(trashExchangeService)

	trashExchange := e.Group("/manage/recycles", jwt.JWTMiddleware())
	trashExchange.POST("", trashExchangeHandler.CreateTrashExchange)
	trashExchange.GET("", trashExchangeHandler.GetAllTrashExchange)
	trashExchange.GET("/:id", trashExchangeHandler.GetTrashExchangeById)
	trashExchange.DELETE("/:id", trashExchangeHandler.DeleteTrashExchange)
}
