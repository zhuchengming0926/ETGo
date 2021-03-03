/**
 * @File: httpPostJson.go
 * @Author: zhuchengming
 * @Description:
 * @Date: 2021/2/26 17:33
 */

package v1

import (
	"ETGo/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HttpPostJsonDemo(ctx *gin.Context)  {
	type ReqParam struct {
		Name    string    `json:"name" binding:"required"`
		Age     int64     `json:"age" validate:"gt=0"`  //age不能传小于等于0的数据
	}

	req := ReqParam{}
	if errParams := ctx.ShouldBind(&req); errParams != nil {
		helper.Render.RenderJsonFail(ctx, errParams)
		return
	}
	validate := validator.New()
	validErr := validate.Struct(req)
	if validErr != nil {
		helper.Render.RenderJsonFail(ctx, validErr)
		return
	}

	helper.Render.RenderJsonSucc(ctx, req)
	return
}