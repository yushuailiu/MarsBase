package logging

import (
	"github.com/go-ini/ini"
	config2 "github.com/yushuailiu/MarsBase/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"time"
	"os"
)

var Log *logrus.Logger

var moduleLogMap map[string]*logrus.Logger

func Bootstrap() {
	Log = logrus.New()
	loggingConfig := config2.GetConfig().Section("logging")

	initLogging(loggingConfig)
}

func initLogging(config *ini.Section) {
	switch config.Key("channel").String() {
	case "file":
		initFileLogging()
	default:
		panic("log does not support channel of " + config.Key("channel").String())
	}
}

func initFileLogging() {
	config := config2.GetConfig().Section("logging.file")

	appName := config2.GetConfig().Section("").Key("name").String()
	basePath := config.Key("basePath").String()
	defaultLog := basePath + "/" + appName + ".log"

	switch config.Key("mode").String() {
	case "single":
		println(basePath + "/error.log")
		pathMap := lfshook.PathMap{
			logrus.ErrorLevel: basePath + "/error.log",
		}
		hook := lfshook.NewHook(
			pathMap,
			&logrus.TextFormatter{},
		)
		hook.SetDefaultPath(defaultLog)
		Log.Hooks.Add(hook)
	case "daily":
		writer := rotateWriter(appName)
		fd, err := os.OpenFile(defaultLog, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			if os.IsNotExist(err) {
				panic("the log path is not exist")
			}
			panic(err)
		}
		hook := lfshook.NewHook(
			lfshook.WriterMap{
				logrus.ErrorLevel: fd,
			},
			&logrus.TextFormatter{},
		)

		hook.SetDefaultWriter(writer)
		Log.AddHook(hook)
	default:
		panic("log mode do not support mode of " + config.Key("mode").String())
	}
}

func GetLogger(moduleName string) *logrus.Logger {

	if log, ok := moduleLogMap[moduleName]; ok {
		return log
	}

	log := logrus.New()

	writer := rotateWriter(moduleName)

	hook := lfshook.NewHook(
		lfshook.WriterMap{},
		&logrus.TextFormatter{},
	)

	hook.SetDefaultWriter(writer)
	log.AddHook(hook)
	return log
}

func rotateWriter(moduleName string) *rotatelogs.RotateLogs {

	config := config2.GetConfig().Section("logging.file")

	basePath := config.Key("basePath").String()

	defaultLog := basePath + "/" + moduleName + ".log"

	days := config.Key("days").MustInt()
	writer, _ := rotatelogs.New(
		basePath+"/"+moduleName+"-%Y-%m-%d.log",
		rotatelogs.WithLinkName(defaultLog),
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(days)),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)
	return writer
}
