package userModel

import (
	"github.com/yushuailiu/MarsBase/pkg/database/mysql"
)

type UserInfo struct {
	mysql.Model
	RoleId		int		`sql:"type:int;not null" json:"role_id" validate:"required,gte=0"`
	Nickname	string	`gorm:"column:nickname" json:"nickname"`
	Email		string	`gorm:"email type:varchar(100)" json:"email" validate:"required,email,omitempty"`
	Phone		string	`gorm:"type:varchar(32)" json:"phone"`
	Avatar		string	`gorm:"type:varchar(1024)" json:"avatar" validate:"uri,omitempty"`
	Status		int	`json:"status"`
	LastIp		int	`json:"last_ip"`
	Intro		string	`json:"intro"`
	Sex		int	`json:"sex"`
	City		string	`json:"city"`
}


func (UserInfo) TableName() string {
	return "user_info"
}