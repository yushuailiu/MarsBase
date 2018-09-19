package friendLinkHandler

import (
	"github.com/kataras/iris"
	"io/ioutil"
	"github.com/yushuailiu/MarsBase/pkg/myhttp"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/yushuailiu/MarsBase/apps/site/models/friendLinkModel"
	"github.com/yushuailiu/MarsBase/pkg/database/mysql"
	"github.com/yushuailiu/MarsBase/pkg/logging"
	"unicode/utf8"
	"reflect"
)

type Param struct {
	Name	string	`json:"name" valid:"required~链接名必须,runelength(1|20)~链接名"`
	Url	string	`json:"url" valid:"required~链接地址必须,url~非法链接,runelength(1|200)~链接地址过长"`
	Sort	int	`json:"sort"`
	Status	int	`json:"status" valid:"range(0|1)"`
}

func Delete(ctx iris.Context) {
	id,err := ctx.Params().GetInt("id")

	if err != nil  {
		myhttp.DefaultParamError(ctx)
		return
	}

	err = mysql.DB.Where("id = ?", id).Delete(&friendLinkModel.FriendLink{}).Error
	if err != nil {
		logging.Log.Error(err)
		myhttp.DefaultSystemError(ctx)
		return
	}
	myhttp.DefaultSuccess(ctx, nil)
}
func UpdateStatus(ctx iris.Context) {
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		myhttp.DefaultParamError(ctx)
		return
	}

	var param map[string]interface{}

	err = json.Unmarshal(body, &param)

	var id,status int
	tId, ok := param["id"]

	if !ok || reflect.TypeOf(tId).Kind() != reflect.Float64 || tId.(float64) < 0 {
		myhttp.ParamError(ctx, "非法id", nil)
		return
	}
	id = int(tId.(float64))

	tStatus, ok := param["status"]

	if !ok || reflect.TypeOf(tStatus).Kind() != reflect.Float64 || !govalidator.InRangeFloat64(tStatus.(float64), 0, 1) {
		myhttp.ParamError(ctx, "状态值应为整型", nil)
		return
	}

	status = int(tStatus.(float64))


	err = mysql.DB.Model(friendLinkModel.FriendLink{}).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		logging.Log.Error(err)
		myhttp.DefaultParamError(ctx)
		return
	}

	myhttp.DefaultSuccess(ctx, nil)
}

func Update(ctx iris.Context) {
	id,err := ctx.Params().GetInt("id")
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		myhttp.DefaultParamError(ctx)
		return
	}

	var param map[string]interface{}

	err = json.Unmarshal(body, &param)
	if err != nil {
		myhttp.DefaultParamError(ctx)
		return
	}

	friendLink := friendLinkModel.FriendLink{}
	mysql.DB.Where("id = ?", id).First(&friendLink)

	if friendLink.ID == 0 {
		myhttp.DefaultNotFound(ctx)
		return
	}

	name, ok := param["name"]
	if ok && reflect.TypeOf(name).Kind() == reflect.String{
		tName := name.(string)
		if len(tName) == 0 || utf8.RuneCountInString(tName) > 20 {
			myhttp.ParamError(ctx, "链接名不合法", nil)
			return
		}
		friendLink.Name = tName
	}

	url, ok := param["url"]
	if ok && reflect.TypeOf(url).Kind() == reflect.String {
		tUrl := url.(string)
		if !govalidator.IsURL(tUrl) {
			myhttp.ParamError(ctx, "非法链接", nil)
			return
		}
		friendLink.Url = tUrl
	}

	sort, ok := param["sort"]

	if ok && reflect.TypeOf(sort).Kind() == reflect.Float64 {
		tSort := int(sort.(float64))
		if tSort < 0 {
			myhttp.ParamError(ctx, "非法排序字段", nil)
			return
		}
		friendLink.Sort = tSort
	}

	status, ok := param["status"]
	if ok && reflect.TypeOf(status).Kind() == reflect.Int {
		tStatus := status.(int)
		if govalidator.InRangeInt(tStatus, 0, 2) {
			myhttp.ParamError(ctx, "非法状态值", nil)
			return
		}
		friendLink.Status = tStatus
	}

	err = mysql.DB.Save(&friendLink).Error

	if err != nil {
		myhttp.DefaultParamError(ctx)
		return
	}
	myhttp.DefaultSuccess(ctx, nil)
}

//func Update(ctx iris.Context) {
//	id,err := ctx.Params().GetInt("id")
//
//	if err != nil  {
//		myhttp.DefaultParamError(ctx)
//		return
//	}
//
//	var param Param
//	err = functions.ReadRequestBody(ctx, &param)
//
//	if err != nil {
//		myhttp.ParamError(ctx, err.Error(), nil)
//		return
//	}
//
//	count := mysql.DB.Model(friendLinkModel.FriendLink{}).Where("id = ?", id).
//		Update(friendLinkModel.FriendLink{
//			Name: param.Name,
//			Url: param.Url,
//			Sort: param.Sort,
//	}).RowsAffected
//
//	if count == 0 {
//		myhttp.DefaultNotFound(ctx)
//		return
//	}
//	myhttp.DefaultSuccess(ctx, nil)
//}


func List(ctx iris.Context) {
	page := ctx.URLParamIntDefault("currentPage", 1)
	pageSize := ctx.URLParamIntDefault("pageSize", 10)
	var friendLinks []*friendLinkModel.FriendLink
	err := mysql.DB.Model(&friendLinkModel.FriendLink{}).
		Limit(pageSize).Offset((page - 1) * pageSize).Find(&friendLinks).Error

	count := 0
	mysql.DB.Model(friendLinkModel.FriendLink{}).Count(&count)
	if err != nil {
		logging.Log.Error(err)
		myhttp.DefaultParamError(ctx)
		return
	}
	myhttp.DefaultSuccess(ctx, iris.Map{
		"list": friendLinks,
		"pagination": iris.Map{
			"total": count,
			"current": page,
			"pageSize": pageSize,
		},
	})
}
func Add(ctx iris.Context) {

	body,err := ioutil.ReadAll(ctx.Request().Body)

	if err != nil {
		myhttp.DefaultParamError(ctx)
		return
	}

	var param Param
	err = json.Unmarshal(body, &param)

	if err != nil {
		myhttp.DefaultParamError(ctx)
		return
	}

	_, err = govalidator.ValidateStruct(&param)

	if err != nil {
		myhttp.ParamError(ctx, err.Error(), nil)
		return
	}

	var friendLink friendLinkModel.FriendLink

	friendLink.Name = param.Name
	friendLink.Url = param.Url
	friendLink.Sort = param.Sort
	friendLink.Status = param.Status

	err = mysql.DB.Create(&friendLink).Error

	if err != nil {
		logging.Log.Error(err)
		myhttp.DefaultParamError(ctx)
		return
	}
	myhttp.DefaultSuccess(ctx, nil)
}