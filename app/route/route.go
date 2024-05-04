package route

import (
	"github.com/labstack/echo/v4"
	supabase "github.com/supabase-community/storage-go"
	"gorm.io/gorm"
	"github.com/redis/go-redis/v9"
)

func New(e *echo.Echo, db *gorm.DB, sb *supabase.Client, rdb *redis.Client) {
	base := e.Group("")
	loginPage := e.Group("/")
	user := e.Group("/users")
	admin := e.Group("/admins")
	report := e.Group("/reports")
	faq := e.Group("/faq")
	recybot := e.Group("/recybot")

	RouteLoginPage(loginPage, db)
	RouteUser(user, db)
	RouteReport(report, db, sb)
	RouteAdmin(admin, db, sb,rdb)
	RouteArticle(base, db, sb)
	RouteDropPoint(base, db)
	RouteFaqs(faq, db)
	RouteRecybot(recybot, db,rdb)
	RouteAchievement(base, db)
	RouteVoucher(base, db, sb)
	RouteMissions(base, db, sb)
	RouteDailyPoint(user, db, sb)
	RouteTrash(base, db)
	RouteTrashExchange(admin, db)
	RouteCommunity(base, db, sb)
	RouteDashboard(base, db)
}
