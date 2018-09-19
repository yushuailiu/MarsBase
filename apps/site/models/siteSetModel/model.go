package siteSetModel

import "github.com/yushuailiu/MarsBase/pkg/database/mysql"

type SiteSet struct {
	mysql.Model
	SiteName	string	`json:"site_name"`
	MetaKeywords	string	`json:"meta_keywords"`
	MetaDescription		string	`json:"meta_description"`
	SiteVersion	string	`json:"site_version"`
	Logo	string	`json:"logo"`
	Favicon	string	`json:"favicon"`
	Copyright string	`json:"copyright"`
}



func (SiteSet) TableName() string {
	return "site_set"
}