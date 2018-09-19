package statisticsHandler

import (
	"github.com/kataras/iris"
	"github.com/yushuailiu/MarsBase/pkg/myhttp"
	"github.com/yushuailiu/MarsBase/apps/user/repositories/adminUserRepository"
)

func SiteStatistics(ctx iris.Context)  {

	adminCount := adminUserRepository.Count(-1, -1, "")

	myhttp.Success(ctx, "", iris.Map{
		"adminCount": adminCount,
	})
}