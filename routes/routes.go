package routes

import "github.com/kataras/iris"

func InitRoutes(app *iris.Application)  {
	InitApiRoutes(app)
	InitAdminApiRoutes(app)
}