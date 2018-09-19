package routes

import (
	"github.com/kataras/iris"
	"github.com/yushuailiu/MarsBase/apps/user/handlers/userHandler"
	"github.com/yushuailiu/MarsBase/apps/site/handler/siteSetHandler"
)

func InitApiRoutes(app *iris.Application) {
	api := app.Party("/api")
	initArticleRoutes(api)
}

func initArticleRoutes(route iris.Party) {
	route.Get("/articles", userHandler.Login)
	route.Get("/site/set", siteSetHandler.GetSiteSet)
	route.Get("/hello", siteSetHandler.Hello)
}
