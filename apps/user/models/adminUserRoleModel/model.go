package adminUserRoleModel

import "github.com/yushuailiu/MarsBase/pkg/database/mysql"

type AdminUserRole struct {
	mysql.Model
	Name	string	`gorm:"name" json:"name"`
	Status	int	`gorm:"status" json:"status"`
}



func (AdminUserRole) TableName() string {
	return "admin_user_role"
}