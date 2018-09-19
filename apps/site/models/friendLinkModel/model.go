package friendLinkModel

import "github.com/yushuailiu/MarsBase/pkg/database/mysql"

type FriendLink struct {
	mysql.Model
	Name	string	`json:"name"`
	Url	string	`json:"url"`
	Sort	int	`json:"sort"`
	Status	int	`json:"status"`
}

func (FriendLink) TableName() string {
	return "site_friend_link"
}