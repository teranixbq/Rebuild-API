package route

import (
	userRepo "recything/features/user/repository"
	achievement"recything/features/achievement/repository"
	"recything/features/voucher/handler"
	"recything/features/voucher/repository"
	"recything/features/voucher/service"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteVoucher(e *echo.Group, db *gorm.DB) {
	achievementRepository := achievement.NewAchievementRepository(db)
	userRepository := userRepo.NewUserRepository(db,achievementRepository)
	voucherRepository := repository.NewVoucherRepository(db)
	voucherService := service.NewVoucherService(voucherRepository, userRepository)
	voucherHandler := handler.NewVoucherHandler(voucherService)

	admin := e.Group("/admins/manage/vouchers", jwt.JWTMiddleware())
	admin.POST("", voucherHandler.CreateVoucher)
	admin.GET("", voucherHandler.GetAllVoucher)
	admin.GET("/:id", voucherHandler.GetVoucherById)
	admin.PUT("/:id", voucherHandler.UpdateVoucher)
	admin.DELETE("/:id", voucherHandler.DeleteVoucherById)

	user := e.Group("/vouchers", jwt.JWTMiddleware())
	user.GET("", voucherHandler.GetAllVoucher)
	user.GET("/:id", voucherHandler.GetVoucherById)
	user.POST("", voucherHandler.CreateExchangeVoucher)
	user.POST("", voucherHandler.CreateExchangeVoucher)

	adminExchange := e.Group("/admins/manage/exchange-point", jwt.JWTMiddleware())
	adminExchange.GET("", voucherHandler.GetAllExchange)
	adminExchange.GET("/:id", voucherHandler.GetByIdExchange)
	adminExchange.PATCH("/:id", voucherHandler.UpdateStatusExchange)
}
