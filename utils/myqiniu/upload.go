package myqiniu

import (
	"os"
	"github.com/yushuailiu/MarsBase/pkg/config"
	"github.com/qiniu/api.v7/storage"
	"github.com/qiniu/api.v7/auth/qbox"
	"fmt"
	"context"
)

func UploadFile(file string, key string) (url string, err error) {
	info, err := os.Stat(file)
	if err != nil || info.IsDir() {
		return "", err
	}

	bucket := config.GetConfig().Section("qiniu").Key("bucket").String()
	accessKey := config.GetConfig().Section("qiniu").Key("accessKey").String()
	secretKey := config.GetConfig().Section("qiniu").Key("secretKey").String()
	host := config.GetConfig().Section("qiniu").Key("host").String()

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}

	// 保存额外参数
	//putExtra := storage.PutExtra{
	//	Params: map[string]string{
	//		"x:name": "github logo",
	//	},
	//}
	putExtra := storage.PutExtra{}
	err = formUploader.PutFile(context.Background(), &ret, upToken, key, file, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(host + ret.Key,ret.Hash)

	return host + ret.Key,nil
}