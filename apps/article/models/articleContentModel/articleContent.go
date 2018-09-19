package articleContentModel

import "github.com/yushuailiu/MarsBase/pkg/database/mysql"

type ArticleContent struct {
	mysql.Model
	Title		string	`gorm:"column:title" sql:"type:varchar(256);not null;"`
	Content		string	`gorm:"column:content;" sql:"type:text;" json:"content"`
	CategoryId	int	`gorm:"column:category_id" sql:"type:int;" json:"category_id"`
	LabelId		int	`gorm:"column:label_id" sql:"type:int" json:"label_id"`
	Intro		string	`gorm:"column:intro" sql:"type:varchar(1024)" json:"intro"`
	Author		int	`gorm:"column:author" sql:"type:int" json:"author"`
	HitsNum		int	`gorm:"column:hits_num" sql:"type:int" json:"hits_num"`
	VotesNum	int	`gorm:"column:votes_num" sql:"type:int" json:"votes_num"`
	CommentsNum	int	`gorm:"column:comments_num" sql:"type:int" json:"comments_num"`
	OnTop		int	`gorm:"column:ontop" sql:"type:int" json:"ontop"`
	Status		int	`gorm:"column:status" sql:"type:int" json:"status"`
	Extra		string	`gorm:"column:extra" sql:"type:varchar(1024)" json:"extra"`
}

func (ArticleContent) TableName() string {
	return "article_content"
}
