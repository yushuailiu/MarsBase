package functions

import (
	"github.com/kataras/iris"
	"io/ioutil"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"fmt"
)

func ReadRequestBody(ctx iris.Context, target interface{}) error {
	body,err := ioutil.ReadAll(ctx.Request().Body)

	if err != nil {
		return fmt.Errorf("参数错误！")
	}

	err = json.Unmarshal(body, target)

	if err != nil {
		return fmt.Errorf("参数错误！")
	}

	_, err = govalidator.ValidateStruct(target)

	if err != nil {
		return err
	}
	return nil
}