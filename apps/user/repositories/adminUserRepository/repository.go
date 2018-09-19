package adminUserRepository

import (
	. "github.com/yushuailiu/MarsBase/pkg/database/mysql"
	"github.com/yushuailiu/MarsBase/apps/user/models/adminUserModel"
	"net"
	. "github.com/yushuailiu/MarsBase/apps/user/models/adminUserRoleModel"
	"github.com/yushuailiu/MarsBase/pkg/logging"
)
type omit *struct{}

type ListUserInfo struct {
	*adminUserModel.AdminUser
	RoleName string `gorm:"column:role_name" json:"role_name"`
	Ip       int    `gorm:"column:last_ip" json:"ip,omitempty"`
	LastIp   string `json:"last_ip"`
	Password omit `gorm:"column:password" json:"password,omitempty"`
}

type UserDetailInfo struct {
	*adminUserModel.AdminUser
	RoleName string `gorm:"column:role_name" json:"role_name"`
	Ip       int    `gorm:"column:last_ip" json:"ip,omitempty"`
	LastIp   string `json:"last_ip"`
}

var allowSortFields = map[string]interface{}{
	"id":nil,
}

var sortTypes = map[string]string{
	"descend": "desc",
	"ascend": "asc",
}

func RoleList() (roles []*AdminUserRole) {
	DB.Find(&roles)
	return
}

func FindUserByUsername(username string) (*UserDetailInfo) {

	var user UserDetailInfo

	DB.Table("admin_user_info").Select("admin_user_info.*, admin_user_role.name as role_name").
		Joins("join admin_user_role on admin_user_role.id=admin_user_info.role_id").
		Where("admin_user_info.username = ?", username).Scan(&user)
	return &user
}

func UserExistsByUsername(username string) bool {
	count := 0
	DB.Model(adminUserModel.AdminUser{}).Where("username = ?", username).Count(&count)
	return count > 0
}

func UserExists(id int) (*adminUserModel.AdminUser) {
	var user adminUserModel.AdminUser
	DB.Where("id = ?", id).First(&user)
	return &user
}

func RegisterUserByEmailAndUsername(email string, username string, password string) int {

	return 0
}

func UserInfo(id int) (*ListUserInfo) {

	var user ListUserInfo

	DB.Table("admin_user_info").Select("admin_user_info.*, admin_user_role.name as role_name").
		Joins("join admin_user_role on admin_user_role.id=admin_user_info.role_id").
		Where("admin_user_info.id = ?", id).Scan(&user)
	return &user
}

func UserList(page, pageSize int, status int,
	sex int, username, field, order string) (users []*ListUserInfo) {
	query := DB.Table("admin_user_info").Select("admin_user_info.*, admin_user_role.name as role_name").
		Joins("join admin_user_role on admin_user_role.id=admin_user_info.role_id")

	if status != -1 {
		query = query.Where("admin_user_info.status = ?", status)
	}
	if sex != -1 {
		query = query.Where("admin_user_info.sex = ?", sex)
	}

	if len(username) > 0 {
		query = query.
			Where("admin_user_info.username like ?", "%"+username+"%")
	}

	if _,ok := allowSortFields[field]; ok {
		if  order, ok := sortTypes[order]; ok {
			query = query.Order(field + " " + order)
		}
	}


	query = query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users)

	for index, user := range users {
		if user.Ip >= 0 {
			users[index].LastIp = net.IPv4(byte(user.Ip>>24), byte(user.Ip>>16), byte(user.Ip>>8), byte(user.Ip)).String()
		}
	}

	return
}

func DeleteUser(id int) {
	tx := DB.Begin()

	tx.Where("id = ?", id).Delete(adminUserModel.AdminUser{})

	tx.Where("user_id = ?", id).Delete(adminUserModel.AdminUser{})

	tx.Commit()
	return
}

func Count(status, sex int, username string) (count int) {
	query := DB.Model(&adminUserModel.AdminUser{})
	if status != -1 {
		query = query.Where("status = ?", status)
	}
	if sex != -1 {
		query = query.Where("sex = ?", sex)
	}
	if len(username) > 0 {
		query = query.Where("username like ?", "%"+username+"%")
	}
	query.Count(&count)
	return
}

func UpdateUserStatus(ids []int, status int) (bool) {

	if status != 0 && status != 1 {
		return false
	}

	err := DB.Model(adminUserModel.AdminUser{}).Where("id in (?)", ids).Update("status", status).Error

	if err != nil {
		logging.Log.Error(err)
		return false
	}

	return true
}
