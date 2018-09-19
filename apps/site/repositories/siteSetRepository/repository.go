package siteSetRepository

import (
	"github.com/yushuailiu/MarsBase/pkg/database/mysql"
	"github.com/yushuailiu/MarsBase/apps/site/models/siteSetModel"
)

type Param struct {
	SiteName	string	`json:"site_name" valid:"required~缺少站点名字段,runelength(1|15)~站点名不合法"`
	MetaKeywords	string	`json:"meta_keywords" valid:"required~站点关键字必须,runelength(1|50)~站点关键字不合法"`
	MetaDescription		int	`json:"meta_description" valid:"required~站点描述必须,runelength(1|200)~站点描述不合法"`
	SiteVersion	string	`json:"site_version" valid:"required~站点版本必须,runelength(1|10)~站点版本不合法"`
	Logo	string	`json:"logo" valid:"required~站点logo必须,length(1|256)~站点logo不合法"`
	Favicon	string	`json:"favicon" valid:"required~站点favicon必须,length(1|256)~站点logo不合法"`
	Copyright string	`json:"copyright" valid:"required~站点Copyright必须,runelength(1|50)~站点copyright不合法"`
}

func GetSiteSet() *siteSetModel.SiteSet {
	var siteSet siteSetModel.SiteSet
	mysql.DB.Model(siteSetModel.SiteSet{}).First(&siteSet)
	return &siteSet
}

func UpdateSiteSet(name, keywords, description, version, logo, favicon, copyright string) (error) {
	var siteSet siteSetModel.SiteSet
	err := mysql.DB.Where("id = ?", 1).Attrs(siteSetModel.SiteSet{
		SiteName: name,
		MetaKeywords: keywords,
		MetaDescription: description,
		SiteVersion: version,
		Logo: logo,
		Favicon: favicon,
		Copyright: copyright,
	}).FirstOrCreate(&siteSet).Error
	if err != nil {
		return err
	}
	return nil
}