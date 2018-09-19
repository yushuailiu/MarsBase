package userRepository

import (
	. "github.com/yushuailiu/MarsBase/pkg/database/mysql"
	"github.com/yushuailiu/MarsBase/apps/user/models/userAuthModel"
	"github.com/yushuailiu/MarsBase/apps/user/models/userModel"
	"net"
)

func AuthExist(identifier string, credential string, loginType string) int {
	var userAuth userAuthModel.UserAuth
	DB.Select("user_id").
		Where("identifier = ? and credential = ? and loginType = ?", identifier,
			credential, loginType).Count(&userAuth)
	return userAuth.UserId
}

func CanRegisterByEmailAndUsername(email string, username string) bool {
	var count int
	DB.Model(&userAuthModel.UserAuth{}).Where("loginType = 'email' and identifier = ?", email).
		Or("where loginType='username' and identifier = ?", username).Count(&count)
	return count > 0
}

func UserExists(id int) (user *userModel.UserInfo) {
	DB.Where("id = ?", id).First(&user)
	return
}

func RegisterUserByEmailAndUsername(email string, username string, password string) int {

	return 0
}


type listUserInfo struct {
	userModel.UserInfo
	RoleName string `gorm:"column:role_name" json:"role_name"`
	Ip int `gorm:"column:last_ip" json:"ip,omitempty"`
	LastIp string `json:"last_ip"`
}

func UserInfo(id int) (*listUserInfo) {

	var user listUserInfo

	DB.Table("user_info").Select("user_info.*, user_role.name as role_name").
		Joins("join user_role on user_role.id=user_info.role_id").
		Where("user_info.id = ?", id).Scan(&user)
	return &user
}

func UserList(keyword string, page, pageSize int) (users []*listUserInfo)  {

	query := DB.Table("user_info").Select("user_info.*, user_role.name as role_name").
		Joins("join user_role on user_role.id=user_info.role_id")
	if len(keyword) > 0 {
		query.
			Where("user_info.nickname like ?", "%" + keyword + "%")
	}
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users)

	for index, user := range users{
		if user.Ip >= 0 {
			users[index].LastIp = net.IPv4(byte(user.Ip >> 24), byte(user.Ip >> 16), byte(user.Ip >> 8), byte(user.Ip)).String()
		}
	}

	return
}

func FindUserByLoginInfo(loginType, identifier, credential string) (user *userAuthModel.UserAuth) {
	DB.Where("loginType = ? and identifier = ? and credential = ?",
		loginType, identifier, credential).First(user)
	return
}

func DeleteUser(id int)  {
	tx := DB.Begin()

	tx.Where("id = ?", id).Delete(userModel.UserInfo{})

	tx.Where("user_id = ?", id).Delete(userAuthModel.UserAuth{})

	tx.Commit()
	return
}

func Count(keyword string) (count int) {
	query := DB.Model(&userModel.UserInfo{})
	if len(keyword) > 0 {
		query.Where("keyword like ?", "%" + keyword + "%")
	}
	query.Count(&count)
	return
}