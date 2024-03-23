package route

import (
	"recything/features/faq/handler"
	"recything/features/faq/repository"
	"recything/features/faq/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteFaqs(e *echo.Group, db *gorm.DB) {
	// User
	faqRepository := repository.NewFaqRepository(db)
	faqService := service.NewFaqService(faqRepository)
	faqHandler := handler.NewFaqHandlers(faqService)

	user := e.Group("")
	user.GET("", faqHandler.GetAllFaqs)
	user.GET("/:id", faqHandler.GetFaqsById)
	
}
