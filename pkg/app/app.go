package app

import (
	"github.com/go-ini/ini"
	"github.com/gin-gonic/gin"
	"io"
	"github.com/yushuailiu/MarsBase/pkg/config"
	"github.com/yushuailiu/MarsBase/pkg/logging"
	"os"
	"github.com/kataras/iris"
)

type App struct {
	Server	*iris.Application
	Env	string
	Mode	string
	Config	*ini.File
	Name	string
	Port	string
}

var app = &App{}

func GetApp() *App {
	return app
}

func Bootstrap(env string) *App {

	app.Server = iris.New()
	app.Env = env

	app.Config = config.DefaultConfig().Bootstrap(env)
	app.Name = app.Config.Section("").Key("name").String()
	app.Port = app.Config.Section("").Key("port").String()

	if IsDevelopment() {
		gin.DefaultWriter = io.MultiWriter(logging.GetRequestLogger(), os.Stdout)
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = logging.GetRequestLogger()
	}
	return app
}

func GetConfig() *ini.File {
	return app.Config
}

func Env() string {
	return app.Env
}

func IsDevelopment() bool {
	return app.Env == "development"
}

func IsDebug() bool {
	return app.Mode == "debug"
}
