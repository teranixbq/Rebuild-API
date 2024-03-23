package route

import (
	"recything/features/community/handler"
	"recything/features/community/repository"
	"recything/features/community/service"
	userhand "recything/features/user/handler"
	userrep "recything/features/user/repository"
	userserv "recything/features/user/service"
	"recything/utils/jwt"
	achievement"recything/features/achievement/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteCommunity(e *echo.Group, db *gorm.DB) {
	achievementRepository := achievement.NewAchievementRepository(db)
	userRepository :=  userrep.NewUserRepository(db,achievementRepository)
	userService := userserv.NewUserService(userRepository)
	userHandler := userhand.NewUserHandlers(userService)

	communityRepo := repository.NewCommunityRepository(db)
	communityService := service.NewCommunityService(communityRepo)
	communityHandler := handler.NewCommunityHandler(communityService)

	admin := e.Group("/admins/manage/communities", jwt.JWTMiddleware())
	admin.POST("", communityHandler.CreateCommunity)
	admin.GET("", communityHandler.GetAllCommunity)
	admin.GET("/:id", communityHandler.GetCommunityById)
	admin.PUT("/:id", communityHandler.UpdateCommunityById)
	admin.DELETE("/:id", communityHandler.DeleteCommunityById)

	user := e.Group("/communities", jwt.JWTMiddleware())
	user.GET("", communityHandler.GetAllCommunity)
	user.GET("/:id", communityHandler.GetCommunityById)
	user.POST("/:idKomunitas", userHandler.JoinCommunity)

	event := e.Group("/admins/manage/event", jwt.JWTMiddleware())
	event.POST("/:idkomunitas", communityHandler.CreateEvent)
	event.GET("/:idkomunitas", communityHandler.ReadAllEvent)
	event.GET("/:idkomunitas/:idevent", communityHandler.ReadEvent)
	event.PUT("/:idkomunitas/:idevent", communityHandler.UpdateEvent)
	event.DELETE("/:idkomunitas/:idevent", communityHandler.DeleteEvent)

	userEvent := e.Group("/users/event", jwt.JWTMiddleware())
	userEvent.GET("/:idkomunitas", communityHandler.ReadAllEvent)
	userEvent.GET("/:idkomunitas/:idevent", communityHandler.ReadEvent)
}
