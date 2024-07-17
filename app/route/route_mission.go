package route

import (
	admin "recything/features/admin/repository"
	user "recything/features/user/repository"

	"recything/features/mission/handler"
	"recything/features/mission/repository"
	"recything/features/mission/service"

	achievement "recything/features/achievement/repository"

	"recything/utils/jwt"

	"recything/utils/storage"

	"github.com/labstack/echo/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gorm.io/gorm"
)

func RouteMissions(e *echo.Group, db *gorm.DB, sb *s3.Client) {
	supabaseConfig := storage.NewStorage(sb)

	adminRepository := admin.NewAdminRepository(db, supabaseConfig)
	achievementRepository := achievement.NewAchievementRepository(db)
	userRepository := user.NewUserRepository(db, achievementRepository)

	missionRepository := repository.NewMissionRepository(db, supabaseConfig)
	missionService := service.NewMissionService(missionRepository, adminRepository, userRepository, supabaseConfig)
	missionHandler := handler.NewMissionHandler(missionService)

	admin := e.Group("/admins/manage/missions", jwt.JWTMiddleware())

	admin.POST("", missionHandler.CreateMission)
	admin.GET("", missionHandler.GetAllMission)
	admin.DELETE("/:id", missionHandler.DeleteMission)
	admin.GET("/:id", missionHandler.FindById)
	admin.PUT("/:id", missionHandler.UpdateMission)
	// admin.PUT("/:id/stages", missionHandler.UpdateMissionStage)

	admin.GET("/approvals", missionHandler.GetAllMissionApproval)
	admin.GET("/approvals/:id", missionHandler.GetMissionApprovalById)
	admin.PUT("/approvals/:id", missionHandler.UpdateStatusApprovalMission)

	user := e.Group("/missions", jwt.JWTMiddleware())
	// user.GET("", missionHandler.GetAllMission)
	user.GET("/:id", missionHandler.FindById)
	user.POST("", missionHandler.ClaimMission)
	user.POST("/proof", missionHandler.CreateUploadMission)
	user.PUT("/proof/:id", missionHandler.UpdateUploadMission)
	user.GET("", missionHandler.GetAllMissionUser)
	user.GET("/history/:idTransaksi", missionHandler.FindHistoryById)

}
