package adminRoleRepository

import (
	. "github.com/yushuailiu/MarsBase/pkg/database/mysql"
	"github.com/yushuailiu/MarsBase/apps/user/models/adminUserRoleModel"
	"fmt"
)

func IsManagerRole(role string) bool {
	if role == "master" || role == "manager" {
		return true
	}
	return false
}

func RoleExists(id int) (*adminUserRoleModel.AdminUserRole)  {
	var role adminUserRoleModel.AdminUserRole
	DB.Where("id = ?", id).First(&role)
	fmt.Println(role)
	return &role
}