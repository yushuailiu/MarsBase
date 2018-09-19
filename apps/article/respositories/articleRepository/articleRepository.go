package articleRepository

import (
	. "github.com/yushuailiu/MarsBase/pkg/database/mysql"
	. "github.com/yushuailiu/MarsBase/apps/article/models/articleContentModel"
)

func ArticleList(page int, pageSize int, keyword string) (articles []*ArticleContent) {
	if keyword == "" {
		DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&articles)
	} else {
		DB.Where("search_text LIKE ?", "%"+keyword+"%").
			Offset((page - 1) * pageSize).Limit(pageSize).Find(&articles)
	}
	return
}

func ArticleCountByKeyword(keyword string) (count int) {
	if keyword == "" {
		DB.Model(&ArticleContent{}).Count(&count)
	} else {
		DB.Model(&ArticleContent{}).Where("search_text like ?", "%"+keyword+"%").Count(&count)
	}
	return
}
