package imageHandler

import (
	"github.com/kataras/iris"
	"mime/multipart"
	"time"
	"strconv"
	"os"
	"path/filepath"
	"io"
	"github.com/yushuailiu/MarsBase/pkg/myhttp"
	"github.com/yushuailiu/MarsBase/pkg/config"
	"github.com/yushuailiu/MarsBase/utils/myqiniu"
	"github.com/yushuailiu/MarsBase/pkg/logging"
)

// todo 根据图片用途分目录，加上日期分目录
func UploadImage(ctx iris.Context)  {
	maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()

	err := ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}

	form := ctx.Request().MultipartForm

	urls := make([]string, 0)
	files := form.File
	for _, item := range files {
		_, localPath, err := saveFile(item[0])
		if err != nil {
			logging.Log.Error(err)
			continue
		}
		url, err := myqiniu.UploadFile(localPath, filepath.Base(localPath))
		if err != nil {
			logging.Log.Error(err)
			continue
		}
		urls = append(urls, url)
	}

	myhttp.DefaultSuccess(ctx, urls)
}

func saveFile(file *multipart.FileHeader) (int64, string, error) {
	unixTime :=	time.Now().Unix()
	filename := strconv.FormatInt(unixTime, 10) + file.Filename

	uploadPath := config.GetConfig().Section("").Key("uploadPath").String()

	src, err := file.Open()
	if err != nil {
		return 0,filename,err
	}

	defer src.Close()

	filePath := filepath.Join(uploadPath, filename)
	// 调试暂时放这里
	out, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, os.FileMode(0666))

	if err != nil {
		return 0, filename, err
	}
	defer out.Close()
	length,err := io.Copy(out, src)
	return length,filePath, err
}