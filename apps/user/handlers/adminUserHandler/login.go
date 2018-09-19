package adminUserHandler

import (
	"github.com/kataras/iris"
	"github.com/yushuailiu/MarsBase/utils/userInfo"
	"github.com/yushuailiu/MarsBase/apps/user/repositories/adminUserRepository"
	"github.com/yushuailiu/MarsBase/pkg/myhttp"
	"github.com/kataras/iris/sessions"
	"github.com/yushuailiu/MarsBase/apps/user/repositories/userRepository"
)

func Login(ctx iris.Context) {
	identifier := ctx.URLParamTrim("identifier")
	credential := ctx.URLParamTrim("credential")

	if len(identifier) == 0 || len(credential) == 0 {
		myhttp.ParamError(ctx,"密码账号不能为空", nil)
		return
	}

	password,err := userInfo.HashPassword(credential)

	if err != nil {
		myhttp.ParamError(ctx, "用户名或密码错误", nil)
		return
	}

	user := adminUserRepository.FindUserByUsername(identifier)

	if user.Status == 1 {
		myhttp.Unauthorized(ctx, "账号已被禁止", nil)
		return
	}

	if user.ID == 0 || user.Password != password {
		myhttp.ParamError(ctx, "用户名或密码错误", nil)
		return
	}

	sess := ctx.Values().Get("session").(*sessions.Session)
	sess.Set("identifier", identifier)
	sess.Set("user", user)
	sess.Set("role", user.RoleName)
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
	sess := ctx.Values().Get("session").(*sessions.Session)

	user := sess.Get("user")
	menu := []string{
		"Guest",
	}

	var info iris.Map

	if user != nil {
		user := user.(map[string]interface{})
		info = iris.Map{
			"nickname": user["nickname"],
			"avatar": user["avatar"],
			"role_name": user["role_name"],
			"id": user["id"],
		}
		menu = []string{
			"User",
			"Site",
			"Dashboard",
			"Exception",
		}
	}
	//menu := menuRepository.GetMenuByRole(user.RoleId)

	myhttp.Success(ctx, "", iris.Map{
		"userInfo": info,
		"menus": menu,
	})
}