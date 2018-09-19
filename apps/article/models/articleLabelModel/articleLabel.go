package articleLabelModel

import "github.com/yushuailiu/MarsBase/pkg/database/mysql"

type ArticleLabel struct {
	mysql.Model
	Name	string	`gorm:"column:name" sql:"type:varchar(128);not null;" json:"name"`
	Del	int	`gorm:"column:del" sql:"type:int;default:0" json:"del"`
}

func (ArticleLabel) TableName() string {
	return "article_label"
}
