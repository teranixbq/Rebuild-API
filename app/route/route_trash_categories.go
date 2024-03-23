package route

import (
	"recything/features/trash_category/handler"
	"recything/features/trash_category/repository"
	"recything/features/trash_category/service"

	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteTrash(e *echo.Group, db *gorm.DB) {
	//manage trash category
	trashCategoryRepository := repository.NewTrashCategoryRepository(db)
	trashCategoryService := service.NewTrashCategoryService(trashCategoryRepository)
	trashCategoryHandler := handler.NewTrashCategoryHandler(trashCategoryService)

	//Manage trash category
	admin := e.Group("admins/manage/trashes", jwt.JWTMiddleware())
	admin.POST("", trashCategoryHandler.CreateCategory)
	admin.GET("", trashCategoryHandler.GetAllCategory)
	admin.GET("/categories", trashCategoryHandler.GetAllCategoriesFetch)
	admin.GET("/:id", trashCategoryHandler.GetById)
	admin.PUT("/:id", trashCategoryHandler.UpdateCategory)
	admin.DELETE("/:id", trashCategoryHandler.DeleteById)

	user := e.Group("/trashes", jwt.JWTMiddleware())
	user.GET("/:id", trashCategoryHandler.GetById)
	user.GET("", trashCategoryHandler.GetAllCategoriesFetch)



}
