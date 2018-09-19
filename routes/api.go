package routes

import (
	"github.com/kataras/iris"
	"github.com/yushuailiu/MarsBase/apps/user/handlers/userHandler"
	"github.com/yushuailiu/MarsBase/apps/payment/handlers/teegonHandler"
)

func InitApiRoutes(app *iris.Application) {
	api := app.Party("/api")
	initArticleRoutes(api)
}

func initArticleRoutes(route iris.Party) {
	route.Get("/articles", userHandler.Login)
	route.Get("/create/order", teegonHandler.CreateOrder)
}
