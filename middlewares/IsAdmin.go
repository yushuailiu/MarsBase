package middlewares

import (
	"github.com/kataras/iris/context"
	"github.com/yushuailiu/MarsBase/pkg/session"
	"github.com/yushuailiu/MarsBase/pkg/myhttp"
	"reflect"
	"github.com/yushuailiu/MarsBase/apps/user/repositories/adminRoleRepository"
)

func IsAdminMiddleware(c context.Context)  {

	sess := session.Session.Start(c)

	roleInterface := sess.Get("role")

	role := ""

	if roleInterface != nil && reflect.TypeOf(role).Kind() == reflect.String {
		role = roleInterface.(string)
	}

	c.Values().Set("session", sess)

	routeName := c.GetCurrentRoute().Name()
	if !adminRoleRepository.IsManagerRole(role) && routeName != "adminLogin" &&
		routeName != "curUser" && routeName != "logout"{
		myhttp.DefaultUnauthorized(c)
		return
	}


	c.Next()
}
