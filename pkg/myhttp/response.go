package myhttp

import (
	"net/http"
	"github.com/kataras/iris"
	"github.com/yushuailiu/MarsBase/pkg/err"
)

func DefaultUnauthorized(c iris.Context) {
	Unauthorized(c, "", nil)
}

func Unauthorized(c iris.Context, msg string, data interface{}) {
	if msg == "" {
		msg = err.GetCodeMsg(http.StatusUnauthorized)
	}

	c.StatusCode(http.StatusUnauthorized)
	c.JSON(iris.Map{
		"msg":	msg,
		"data":	data,
	})
}

func DefaultNotFound(c iris.Context) {
	NotFound(c, "", nil)
}
func NotFound(c iris.Context, msg string, data interface{}) {
	if msg == "" {
		msg = err.GetCodeMsg(http.StatusNotFound)
	}

	c.StatusCode(http.StatusNotFound)
	c.JSON(iris.Map{
		"msg":	msg,
		"data":	data,
	})
}

func DefaultParamError(c iris.Context) {
	ParamError(c, "", nil)
}

func ParamError(c iris.Context, msg string, data interface{}) {

	if msg == "" {
		msg = err.GetCodeMsg(http.StatusBadRequest)
	}

	c.StatusCode(http.StatusBadRequest)
	c.JSON(iris.Map{
		"msg":	msg,
		"data":	data,
	})
}

func DefaultSuccess(c iris.Context, data interface{}) {
	Success(c, "", data)
}

func Success(c iris.Context, msg string, data interface{}) {

	if msg == "" {
		msg = err.GetCodeMsg(http.StatusOK)
	}

	c.StatusCode(http.StatusOK)
	c.JSON(iris.Map{
		"msg":	msg,
		"data":	data,
	})
}

func DefaultSystemError(c iris.Context) {
	SystemError(c, "", nil)
}

func SystemError(c iris.Context, msg string, data interface{}) {

	if msg == "" {
		msg = err.GetCodeMsg(http.StatusInternalServerError)
	}

	c.StatusCode(http.StatusInternalServerError)
	c.JSON(iris.Map{
		"msg":	msg,
		"data":	data,
	})
}
