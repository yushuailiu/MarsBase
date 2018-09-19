package userAuthModel

import "github.com/yushuailiu/MarsBase/pkg/database/mysql"

type UserAuth struct {
	mysql.Model
	UserId		int	`gorm:"user_id" json:"user_id"`
	LoginType	string	`gorm:"login_type" json:"login_type"`
	Identifier	string	`gorm:"identifier" json:"identifier"`
	Credential	string	`gorm:"credential" json:"credential"`
	Status		int	`gorm:"type:int" json:"status"`
}



func (UserAuth) TableName() string {
	return "user_auth"
}