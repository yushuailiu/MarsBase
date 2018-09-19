package siteSetHandler

import (
	"github.com/kataras/iris"
	"github.com/yushuailiu/MarsBase/apps/site/repositories/siteSetRepository"
	"github.com/yushuailiu/MarsBase/pkg/myhttp"
	"io/ioutil"
	"encoding/json"
)

func Hello(ctx iris.Context) {
	myhttp.DefaultSuccess(ctx, "hello world!")
}
func GetSiteSet(ctx iris.Context) {
	siteSet := siteSetRepository.GetSiteSet()

	myhttp.DefaultSuccess(ctx, siteSet)
}

func UpdateSiteSet(ctx iris.Context) {
	type Param struct {
		SiteName	string	`json:"site_name" valid:"required~缺少站点名字段,runelength(1|15)~站点名不合法"`
		MetaKeywords	string	`json:"meta_keywords" valid:"required~站点关键字必须,runelength(1|50)~站点关键字不合法"`
		MetaDescription		string	`json:"meta_description" valid:"required~站点描述必须,runelength(1|200)~站点描述不合法"`
		SiteVersion	string	`json:"site_version" valid:"required~站点版本必须,runelength(1|10)~站点版本不合法"`
		Logo	string	`json:"logo" valid:"required~站点logo必须,length(1|256)~站点logo不合法"`
		Favicon	string	`json:"favicon" valid:"required~站点favicon必须,length(1|256)~站点logo不合法"`
		Copyright string	`json:"copyright" valid:"required~站点Copyright必须,runelength(1|50)~站点copyright不合法"`
	}

	var param Param

	body, err := ioutil.ReadAll(ctx.Request().Body)
	err = json.Unmarshal(body, &param)

	if err != nil {
		myhttp.ParamError(ctx, err.Error(), nil)
		return
	}

	err = siteSetRepository.UpdateSiteSet(param.SiteName,
		param.MetaKeywords, param.MetaDescription, param.SiteVersion,
		param.Logo, param.Favicon, param.Copyright)

	if err != nil {
		myhttp.DefaultParamError(ctx)
		return
	}
	myhttp.DefaultSuccess(ctx, nil)
}