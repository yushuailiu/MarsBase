package adminUserHandler

import (
	"github.com/kataras/iris"
	"github.com/yushuailiu/MarsBase/pkg/myhttp"
	"github.com/yushuailiu/MarsBase/utils/functions"
	"github.com/yushuailiu/MarsBase/utils/userInfo"
	"github.com/yushuailiu/MarsBase/pkg/database/mysql"
	"github.com/yushuailiu/MarsBase/apps/user/repositories/adminRoleRepository"
	"github.com/yushuailiu/MarsBase/apps/user/models/adminUserModel"
	"github.com/yushuailiu/MarsBase/apps/user/repositories/adminUserRepository"
	"github.com/kataras/iris/sessions"
	"encoding/json"
	"io/ioutil"
	"github.com/asaskevich/govalidator"
)

func UserList(ctx iris.Context)  {
	page := ctx.URLParamIntDefault("currentPage", 1)
	pageSize := ctx.URLParamIntDefault("pageSize", 10)
	username := ctx.URLParamTrim("username")
	status := ctx.URLParamIntDefault("status", -1)
	sex := ctx.URLParamIntDefault("sex", -1)
	field := ctx.URLParamTrim("field")
	order := ctx.URLParamTrim("order")

	pageSize = functions.MinInt(pageSize, 50)


	users := adminUserRepository.UserList(page,
		pageSize, status, sex, username, field, order)

	count := adminUserRepository.Count(status, sex, username)

	myhttp.Success(ctx, "", iris.Map{
		"list": users,
		"pagination": iris.Map{
			"total": count,
			"current": page,
			"pageSize": pageSize,
		},
	})
}

func UserInfo(ctx iris.Context)  {

	id, _ := ctx.Params().GetInt("id")

	user := adminUserRepository.UserInfo(id)

	if user == nil {
		myhttp.DefaultNotFound(ctx)
		return
	}

	myhttp.DefaultSuccess(ctx, user)
}

func DeleteUser(ctx iris.Context)  {
	id, _ := ctx.Params().GetInt("id")

	adminUserRepository.DeleteUser(id)
	myhttp.DefaultSuccess(ctx, nil)
}

func UpdateUser(ctx iris.Context)  {
	id, _ := ctx.Params().GetInt("id")

	oldUserInfo := adminUserRepository.UserExists(id)
	if oldUserInfo == nil {
		myhttp.DefaultNotFound(ctx)
		return
	}

	user := new(adminUserModel.AdminUser)

	err := ctx.ReadForm(user)

	if err != nil {
		myhttp.DefaultNotFound(ctx)
		return
	}

	if user.RoleId == 0 || adminRoleRepository.RoleExists(user.RoleId) == nil{
		myhttp.ParamError(ctx, "角色不存在", nil)
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

	mysql.DB.Save(oldUserInfo)

	myhttp.DefaultSuccess(ctx, nil)
}

func AddUser(ctx iris.Context)  {
	user := new(adminUserModel.AdminUser)

	err := ctx.ReadJSON(user)

	if err != nil {
		myhttp.DefaultParamError(ctx)
		return
	}

	user.ID = 0
	user.LastIp = 0

	if adminUserRepository.UserExistsByUsername(user.Username) {
		myhttp.ParamError(ctx,"用户名已存在", nil)
	}

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

	if !userInfo.IsValidUsername(user.Username) {
		myhttp.ParamError(ctx, "非法用户名", nil)
		return
	}

	if !userInfo.IsValidPassword(user.Password) {
		myhttp.ParamError(ctx, "密码不合法", nil)
		return
	}

	user.Password,_ = userInfo.HashPassword(user.Password)

	mysql.DB.Create(user)
	if user.ID == 0 {
		myhttp.DefaultSystemError(ctx)
		return
	}

	myhttp.DefaultSuccess(ctx, nil)
	return
}

func CurUserDetailInfo(ctx iris.Context)  {

	sess := ctx.Values().Get("session").(*sessions.Session)

	user := sess.Get("user").(map[string]interface{})

	userId := int(user["id"].(float64))


	info := adminUserRepository.UserInfo(userId)
	if info.ID == 0 {
		myhttp.DefaultParamError(ctx)
		return
	}

	myhttp.DefaultSuccess(ctx, info)
}

func UpdateCurUserInfo(ctx iris.Context)  {

	sess := ctx.Values().Get("session").(*sessions.Session)

	sessionInfo := sess.Get("user").(map[string]interface{})

	userId := int(sessionInfo["id"].(float64))

	oldUserInfo := adminUserRepository.UserExists(userId)


	if oldUserInfo == nil {
		myhttp.DefaultNotFound(ctx)
		return
	}

	type Params struct {
		Avatar string `json:"avatar" valid:"required~缺少头像字段,length(4|512)~非法头像地址"`
		Email string  `json:"email" valid:"required~缺少邮箱字段,email~非法邮箱"`
		Nickname string `json:"nickname" valid:"required~缺少昵称字段,length(2|15)~非法昵称"`
		Phone	string	`json:"phone" valid:"required~缺少手机字段,matches(^1[\\d]{10}$)~非法手机"`
		Sex 	int 	`json:"sex" valid:"range(1|2)~非法性别"`
		Intro	string `json:"intro" valid:"runelength(0|100)"`
	}

	var user Params


	body, err := ioutil.ReadAll(ctx.Request().Body)
	err = json.Unmarshal(body, &user)

	if err != nil {
		myhttp.DefaultParamError(ctx)
		return
	}

	_, err = govalidator.ValidateStruct(&user)

	if err != nil {
		myhttp.ParamError(ctx, err.Error(), nil)
		return
	}


	oldUserInfo.Sex = user.Sex
	oldUserInfo.Email = user.Email
	oldUserInfo.Nickname = user.Nickname
	oldUserInfo.Phone = user.Phone
	oldUserInfo.Intro = user.Intro

	oldUserInfo.Avatar = user.Avatar

	err = mysql.DB.Save(oldUserInfo).Error
	if err != nil {
		myhttp.DefaultParamError(ctx)
		return
	}

	myhttp.DefaultSuccess(ctx, nil)
}

func UpdateUserStatus(ctx iris.Context)  {
	type Params struct {
		Status int `json:"status"`
		Ids []int `json:"ids"`
	}

	var params Params

	ctx.ReadJSON(&params)
	if adminUserRepository.UpdateUserStatus(params.Ids, params.Status) {
		myhttp.DefaultSuccess(ctx, nil)
		return
	}
	myhttp.ParamError(ctx, "", nil)
}

