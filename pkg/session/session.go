package session

import (
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
	"time"
	"github.com/kataras/iris/sessions"
	"github.com/yushuailiu/MarsBase/pkg/config"
	"github.com/yushuailiu/MarsBase/pkg/app"
)

var Session *sessions.Sessions

func Bootstrap() {
	sessionConfig := config.GetConfig().Section("session")

	switch sessionConfig.Key("driver").String() {
	case "redis":
		initRedisSession()
	case "standalone":
		initStandaloneSession()
	}
}

func initStandaloneSession() {
	println("init standalone session")
	Session = sessions.New(sessions.Config{
		Cookie:				app.GetApp().Name + "_session",
		Expires:			60 * time.Minute,
		DisableSubdomainPersistence:	true,
		AllowReclaim:			true,
	})
}

func initRedisSession() {
	println("init redis session")

	redisConfig := config.GetConfig().Section("redis.session")
	println(redisConfig.Key("host").String() + ":" + redisConfig.Key("port").String())

	db := redis.New(service.Config{
		Addr:		redisConfig.Key("host").String() + ":" + redisConfig.Key("port").String(),
		Password:	redisConfig.Key("password").String(),
		Database:	redisConfig.Key("db").String(),
	})
	Session = sessions.New(sessions.Config{
		Cookie:				app.GetApp().Name + "_session_id",
		Expires:			60 * time.Minute,
		DisableSubdomainPersistence:	true,
		AllowReclaim:			true,
	},
	)
	Session.UseDatabase(db)
}
