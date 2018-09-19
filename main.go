//go:generate go-assets-builder config -o config/config.go -p config
package main

import (
	"github.com/kataras/iris"
	recover2 "github.com/kataras/iris/middleware/recover"
	"github.com/yushuailiu/MarsBase/routes"
	app2 "github.com/yushuailiu/MarsBase/pkg/app"
	"github.com/spf13/pflag"
	"github.com/yushuailiu/MarsBase/pkg/session"
	"github.com/yushuailiu/MarsBase/pkg/database/mysql"
	"github.com/yushuailiu/MarsBase/pkg/logging"
)

func main() {
	var env *string = pflag.String("env", "", "运行环境类型")
	pflag.Parse()

	app := app2.Bootstrap(*env)
	session.Bootstrap()
	mysql.Bootstrap()
	logging.Bootstrap()

	app.Server.Use(recover2.New())

	app.Server.RegisterView(iris.HTML("./views", ".html"))
	app.Server.StaticWeb("/", "./public")

	routes.InitRoutes(app.Server)

	app.Server.Run(iris.Addr(":" + app.Port))
}
