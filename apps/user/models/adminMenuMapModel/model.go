package adminMenuMapModel
import "github.com/yushuailiu/MarsBase/pkg/database/mysql"

type AdminMenuMap struct {
	mysql.Model
	MenuId	string	`gorm:"type:int" json:"menu_id"`
	RoleId	string	`gorm:"type:int" json:"role_id"`
	Status		int	`gorm:"type:int" json:"status"`
}



func (AdminMenuMap) TableName() string {
	return "admin_menu_map"
}