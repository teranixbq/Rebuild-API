package route

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {
	base := e.Group("")
	loginPage := e.Group("/")
	user := e.Group("/users")
	admin := e.Group("/admins")
	report := e.Group("/reports")
	faq := e.Group("/faq")
	recybot := e.Group("/recybot")


	RouteLoginPage(loginPage, db)
	RouteUser(user, db)
	RouteReport(report, db)
	RouteAdmin(admin, db)
	RouteArticle(base, db)
	RouteDropPoint(base, db)
	RouteFaqs(faq, db)
	RouteRecybot(recybot, db)
	RouteAchievement(base, db)
	RouteVoucher(base, db)
	RouteMissions(base, db)
	RouteDailyPoint(user,db)
	RouteTrash(base,db)
	RouteTrashExchange(admin, db)
	RouteCommunity(base, db)
	RouteDashboard(base, db)
}
