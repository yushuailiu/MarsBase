package userHandler

import (
	"github.com/kataras/iris"
	"github.com/yushuailiu/MarsBase/pkg/myhttp"
	"github.com/yushuailiu/MarsBase/utils/functions"
	"github.com/yushuailiu/MarsBase/apps/user/repositories/userRepository"
	"github.com/yushuailiu/MarsBase/apps/user/models/userModel"
	"github.com/yushuailiu/MarsBase/utils/userInfo"
	"github.com/yushuailiu/MarsBase/pkg/database/mysql"
	"github.com/yushuailiu/MarsBase/apps/user/models/userAuthModel"
	"golang.org/x/crypto/scrypt"
	"github.com/yushuailiu/MarsBase/pkg/app"
	"github.com/yushuailiu/MarsBase/apps/user/repositories/adminRoleRepository"
)

func UserList(ctx iris.Context)  {
	keyword := ctx.URLParamTrim("keyword")
	page := ctx.URLParamIntDefault("page", 1)
	pageSize := ctx.URLParamIntDefault("pageSize", 10)

	pageSize = functions.MaxInt(pageSize, 50)


	users := userRepository.UserList(keyword, page, pageSize)

	count := userRepository.Count(keyword)

	myhttp.Success(ctx, "", iris.Map{
		"list": users,
		"keyword": keyword,
		"total": count,
		"page": page,
		"pageSize": pageSize,
	})
}

func UserInfo(ctx iris.Context)  {

	id, _ := ctx.Params().GetInt("id")

	user := userRepository.UserInfo(id)

	if user == nil {
		myhttp.DefaultNotFound(ctx)
		return
	}

	myhttp.DefaultSuccess(ctx, user)
}

func DeleteUser(ctx iris.Context)  {
	id, _ := ctx.Params().GetInt("id")

	userRepository.DeleteUser(id)
	myhttp.DefaultSuccess(ctx, nil)
}

func UpdateUser(ctx iris.Context)  {
	id, _ := ctx.Params().GetInt("id")

	oldUserInfo := userRepository.UserExists(id)
	if oldUserInfo == nil {
		myhttp.DefaultNotFound(ctx)
		return
	}

	user := new(userModel.UserInfo)

	err := ctx.ReadForm(user)

	if err != nil {
		myhttp.DefaultNotFound(ctx)
		return
	}

	oldUserInfo.RoleId = user.RoleId

	if len(user.Email) == 0 || !userInfo.IsValidEmail(user.Email) {
		myhttp.ParamError(ctx, "非法邮件", nil)
		return
	}
	oldUserInfo.Email = user.Email

	if len(user.Nickname) == 0 || !userInfo.IsValidNickname(user.Nickname) {
		myhttp.ParamError(ctx, "非法昵称", nil)
		return
	}
	oldUserInfo.Nickname = user.Nickname

	if len(user.Phone) == 0 || !userInfo.IsValidPhone(user.Phone) {
		myhttp.ParamError(ctx, "手机号格式错误", nil)
		return
	}
	oldUserInfo.Phone = user.Phone

	if len(user.Avatar) > 1000 {
		myhttp.ParamError(ctx, "头像地址过长", nil)
		return
	}
	oldUserInfo.Avatar = user.Avatar

	if len(user.Intro) > 1000 {
		myhttp.ParamError(ctx, "简介过长", nil)
		return
	}
	oldUserInfo.Intro = user.Intro

	if len(user.City) > 200 {
		myhttp.ParamError(ctx, "城市名过长", nil)
		return
	}
	oldUserInfo.City = user.City

	mysql.DB.Save(oldUserInfo)

	myhttp.DefaultSuccess(ctx, nil)
}

func AddUser(ctx iris.Context)  {
	user := new(userModel.UserInfo)

	err := ctx.ReadForm(user)

	if err != nil {
		myhttp.DefaultParamError(ctx)
		return
	}

	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	user.ID = 0
	user.LastIp = 0

	// todo 封装一个类似 laravel 的 validator
	if user.RoleId == 0 || nil == adminRoleRepository.RoleExists(user.RoleId) {
		myhttp.ParamError(ctx, "角色不存在", nil)
		return
	}

	if len(user.Email) != 0 && !userInfo.IsValidEmail(user.Email) {
		myhttp.ParamError(ctx, "非法邮件", nil)
		return
	}

	if len(user.Nickname) == 0 || !userInfo.IsValidNickname(user.Nickname) {
		myhttp.ParamError(ctx, "非法昵称", nil)
		return
	}

	if len(user.Phone) != 0 && !userInfo.IsValidPhone(user.Phone) {
		myhttp.ParamError(ctx, "手机号格式错误", nil)
		return
	}

	if len(user.Avatar) > 1000 {
		myhttp.ParamError(ctx, "头像地址过长", nil)
		return
	}

	if len(user.Intro) > 1000 {
		myhttp.ParamError(ctx, "简介过长", nil)
		return
	}

	if len(user.City) > 200 {
		myhttp.ParamError(ctx, "城市名过长", nil)
		return
	}

	if !userInfo.IsValidUsername(username) {
		myhttp.ParamError(ctx, "非法用户名", nil)
		return
	}

	if !userInfo.IsValidPassword(password) {
		myhttp.ParamError(ctx, "密码不合法", nil)
		return
	}


	tx := mysql.DB.Begin()

	if !tx.NewRecord(*user) {
		tx.Rollback()
		myhttp.DefaultParamError(ctx)
		return
	}

	auth := new(userAuthModel.UserAuth)

	bytePassword,err := scrypt.Key([]byte(password),
		[]byte(app.GetConfig().Section("").Key("encryptKey").String()),
		16384, 8, 1, 64)

	auth.LoginType = "username"
	auth.Identifier = username
	auth.Credential = string(bytePassword)
	auth.UserId = user.ID

	if !tx.NewRecord(auth) {
		tx.Rollback()
		myhttp.DefaultParamError(ctx)
		return
	}

	myhttp.DefaultSuccess(ctx, nil)
	return
}



