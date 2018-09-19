package articleCategoryModel

import "github.com/yushuailiu/MarsBase/pkg/database/mysql"

type ArticleCategory struct {
	mysql.Model
	Name	string	`gorm:"column:name" sql:"type:varchar(128);not null;"`
	Del	int	`gorm:"column:del" sql:"type:int;default:0" json:"del"`
}

func (ArticleCategory) TableName() string {
	return "article_category"
}
