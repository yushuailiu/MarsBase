package adminUserHandler

import (
	"github.com/kataras/iris"
	"github.com/yushuailiu/MarsBase/utils/functions"
	"github.com/yushuailiu/MarsBase/pkg/myhttp"
	"github.com/yushuailiu/MarsBase/apps/user/repositories/adminUserRepository"
)

func RoleList(ctx iris.Context)  {
	pageSize := ctx.URLParamIntDefault("pageSize", 10)

	pageSize = functions.MaxInt(pageSize, 50)


	roles := adminUserRepository.RoleList()

	myhttp.Success(ctx, "", roles)
}