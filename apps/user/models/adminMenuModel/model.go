package adminMenuModel


import "github.com/yushuailiu/MarsBase/pkg/database/mysql"

type AdminMenu struct {
	mysql.Model
	Name	string	`gorm:"type:varchar(128)" json:"name"`
	ParentName	string	`gorm:"type:varchar(128)" json:"parent_name"`
	App	string	`gorm:"type:varchar(128)" json:"app"`
	Page		string `gorm:"type:varchar(128)" json:"page"`
	Status		int	`gorm:"type:int" json:"status"`
}



func (AdminMenu) TableName() string {
	return "admin_menu"
}