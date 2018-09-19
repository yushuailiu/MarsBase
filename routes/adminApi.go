package routes

import (
	"github.com/kataras/iris"
	"github.com/yushuailiu/MarsBase/apps/article/handlers/articleHandler"
	"github.com/yushuailiu/MarsBase/middlewares"
	"github.com/yushuailiu/MarsBase/apps/user/handlers/adminUserHandler"
	"github.com/yushuailiu/MarsBase/apps/user/handlers/statisticsHandler"
	"github.com/yushuailiu/MarsBase/apps/site/handler/imageHandler"
	"github.com/yushuailiu/MarsBase/apps/site/handler/siteSetHandler"
	"github.com/yushuailiu/MarsBase/apps/site/handler/friendLinkHandler"
)

func InitAdminApiRoutes(app *iris.Application) {
	api := app.Party("/admin/api")
	api.Use(middlewares.IsAdminMiddleware)
	initAdminArticleRoutes(api)
}

func initAdminArticleRoutes(party iris.Party) {
	party.Get("/article", articleHandler.List)

	party.Get("/login", adminUserHandler.Login).Name = "adminLogin"
	party.Get("/logout", adminUserHandler.Logout).Name = "logout"

	// 用户管理
	party.Get("/user", adminUserHandler.UserList)
	party.Get("/user/{id:int min(1)}", adminUserHandler.UserInfo)
	party.Post("/user", adminUserHandler.AddUser)
	party.Put("/user/{id:int min(1)}", adminUserHandler.UpdateUser)
	party.Delete("/user/{id:int min(1)", adminUserHandler.DeleteUser)

	party.Get("/role", adminUserHandler.RoleList)

	party.Get("/curuser", adminUserHandler.CurUserInfo).Name = "curUser"

	party.Put("/user/status", adminUserHandler.UpdateUserStatus)

	// 当前登陆用户
	party.Get("/curuserinfo", adminUserHandler.CurUserDetailInfo)

	party.Put("/curuserinfo", adminUserHandler.UpdateCurUserInfo)

	party.Post("/image/upload", imageHandler.UploadImage)

	// 站点设置
	party.Get("/site/set", siteSetHandler.GetSiteSet)
	party.Post("/site/set", siteSetHandler.UpdateSiteSet)

	// 友情链接
	party.Post("/site/friendlink", friendLinkHandler.Add)
	party.Delete("/site/friendlink/{id:int min(1)}", friendLinkHandler.Delete)
	party.Put("/site/friendlink/{id:int min(1)}", friendLinkHandler.Update)
	party.Get("/site/friendlink", friendLinkHandler.List)
	party.Post("/site/friendlink/status", friendLinkHandler.UpdateStatus)


	party.Get("/site/statistics", statisticsHandler.SiteStatistics)
}




