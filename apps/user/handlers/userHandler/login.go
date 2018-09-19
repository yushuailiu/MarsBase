package userHandler

import (
	"github.com/kataras/iris"
	"github.com/yushuailiu/MarsBase/utils/userInfo"
	"github.com/yushuailiu/MarsBase/apps/user/repositories/userRepository"
	"github.com/yushuailiu/MarsBase/pkg/myhttp"
	"github.com/kataras/iris/sessions"
	"golang.org/x/crypto/scrypt"
	"github.com/yushuailiu/MarsBase/pkg/app"
)

func Login(ctx iris.Context) {
	identifier := ctx.URLParamTrim("identifier")
	credential := ctx.URLParamTrim("credential")

	if len(identifier) == 0 || len(credential) == 0 {
		myhttp.ParamError(ctx,"密码账号不能为空", nil)
		return
	}

	bytePassword,err := scrypt.Key([]byte(credential),
		[]byte(app.GetConfig().Section("").Key("encryptKey").String()),
		16384, 8, 1, 64)
	if err != nil {
		myhttp.ParamError(ctx, "用户名或密码错误", nil)
		return
	}

	userAuth := userRepository.FindUserByLoginInfo("username", identifier, credential)

	if userAuth == nil || userAuth.Credential != string(bytePassword) {
		myhttp.ParamError(ctx, "用户名或密码错误", nil)
		return
	}

	user := userRepository.UserInfo(userAuth.UserId)

	sess := ctx.Values().Get("session").(*sessions.Session)
	sess.Set("identifier", identifier)
	sess.Set("user", user)
	sess.Set("role", "manager")
	myhttp.DefaultSuccess(ctx, nil)
}

func Logout(ctx iris.Context)  {
	sess := ctx.Values().Get("session").(*sessions.Session)
	sess.Destroy()
	myhttp.DefaultSuccess(ctx, nil)
}

func Register(ctx iris.Context) {
	email := ctx.PostValueTrim("email")
	password := ctx.PostValueTrim("password")
	username := ctx.PostValueTrim("username")

	if !userInfo.IsValidEmail(email) {
		myhttp.ParamError(ctx, "错误的邮箱", nil)
		return
	}
	if !userInfo.IsValidPassword(password) {
		myhttp.ParamError(ctx, "密码不合法", nil)
		return
	}
	if !userInfo.IsValidUsername(username) {
		myhttp.ParamError(ctx, "用户名不合法", nil)
		return
	}

	canRegister := userRepository.CanRegisterByEmailAndUsername(email, username)

	if !canRegister {
		myhttp.ParamError(ctx, "用户名或密码", nil)
		return
	}

}

func CurUserInfo(ctx iris.Context)  {
	//sess := ctx.Values().Get("session").(*sessions.Session)

	//user := sess.Get("user").(*userModel.UserInfo)

	//menu := menuRepository.GetMenuByRole(user.RoleId)

	myhttp.Success(ctx, "", iris.Map{
		"menus": []string{
			"User",
			"Site",
		},
	})
}