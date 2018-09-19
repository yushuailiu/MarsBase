package adminUserModel

import "github.com/yushuailiu/MarsBase/pkg/database/mysql"

type AdminUser struct {
	mysql.Model
	RoleId		int	`gorm:"role_id" json:"role_id"`
	Username	string	`gorm:"username" json:"username"`
	Nickname	string `gorm:"column:nickname" json:"nickname"`
	Password	string	`gorm:"column:password" json:"password"`
	Email		string	`gorm:"column:email" json:"email"`
	Phone		string	`gorm:"column:phone" json:"phone"`
	Avatar		string	`gorm:"column:avatar" json:"avatar"`
	Status		int	`gorm:"column:status" json:"status"`
	LastIp		int	`gorm:"column:last_ip" json:"last_ip"`
	Sex		int	`gorm:"column:sex" json:"sex"`
	Intro		string	`json:"intro"`
}



func (AdminUser) TableName() string {
	return "admin_user_info"
}