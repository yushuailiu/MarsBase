package articleLabelMapModel

import "github.com/yushuailiu/MarsBase/pkg/database/mysql"

type ArticleLabelMap struct {
	mysql.Model
	Name	string	`gorm:"column:name" sql:"type:varchar(128);not null;" json:"name"`
	Del	int	`gorm:"column:del" sql:"type:int;default:0" json:"del" json:"del"`
}

func (ArticleLabelMap) TableName() string {
	return "article_label_map"
}
