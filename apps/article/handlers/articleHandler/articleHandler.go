package articleHandler

import (
	"github.com/kataras/iris"
	"github.com/yushuailiu/MarsBase/apps/article/respositories/articleRepository"
	"github.com/yushuailiu/MarsBase/pkg/myhttp"
)

func List(ctx iris.Context) {
	page := ctx.URLParamIntDefault("page", 1)
	pageSize := ctx.URLParamIntDefault("pageSize", 10)
	keyword := ctx.URLParamTrim("keyword")

	list := articleRepository.ArticleList(page, pageSize, keyword)

	myhttp.Success(ctx, "", iris.Map{
		"list":		list,
		"page":		page,
		"pageSize":	pageSize,
		"count":	articleRepository.ArticleCountByKeyword(keyword),
	})
	return
}

func Detail(ctx iris.Context) {
	id := ctx.URLParamIntDefault("id", 0)

	if id <= 0 {
		myhttp.DefaultNotFound(ctx)
		return
	}
}

func Add(ctx iris.Context) {
	title := ctx.URLParamTrim("title")
	content := ctx.URLParamTrim("content")
	categoryId := ctx.URLParamTrim("category_id")
	intro := ctx.URLParamTrim("intro")

	println(title, content, categoryId, intro)
}
