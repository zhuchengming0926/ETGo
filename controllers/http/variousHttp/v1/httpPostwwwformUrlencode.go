/**
 * @File: httpPostwwwformUrlencode.go
 * @Author: zhuchengming
 * @Description:普通的post表单请求，Content-Type=application/x-www-form-urlencoded
 * @Date: 2021/2/26 17:33
 */

package v1

import (
	"ETGo/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HttpPostWwwFormUrlEncodeDemo(ctx *gin.Context)  {
	type ReqParam struct {
		Name   string    `form:"name" binding:"required"`
		Age    int64     `form:"age" binding:"required" validate:"gt=0"`
	}
	//gin框架获取方式,做了兼容，对form-data方式和x-www-form-urlencoded
	req := ReqParam{}
	if errParams := ctx.ShouldBind(&req); errParams != nil {
		helper.Render.RenderJsonFail(ctx, errParams)
	}
	validate := validator.New()
	validErr := validate.Struct(validate)
	if validErr != nil {
		helper.Render.RenderJsonFail(ctx, validErr)
	}

	helper.Render.RenderJsonSucc(ctx, req)
	return

	//普通方式获取方式
	//r := ctx.Request
	//r.ParseForm()
	//name := r.PostForm["name"]
	//fmt.Println(name)
}

