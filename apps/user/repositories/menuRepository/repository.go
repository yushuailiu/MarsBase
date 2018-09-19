package menuRepository

import (
	"github.com/yushuailiu/MarsBase/pkg/database/mysql"
	"github.com/yushuailiu/MarsBase/apps/user/models/adminMenuModel"
)

func GetMenuByRole(roleId int) (menus []*adminMenuModel.AdminMenu) {
	mysql.DB.Select("admin_menu.*").
		Joins("admin_menu on admin_menu_map.role_id = ?", roleId).
		Find(&menus)
	return
}