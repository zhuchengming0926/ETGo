/**
 * @File: httpPostOctetTream.go
 * @Author: zhuchengming
 * @Description:
 * @Date: 2021/3/31 17:14
 */

package v1

import (
	"ETGo/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

//这个Content-type:application/octet-stream 只是说明上传文件的方式，
//而enctype="multipart/form-data" 这是是说参数传递时的打包方式
func HttpPostOctetTream(ctx *gin.Context)  {
	fh, err := ctx.FormFile("fls")
	if err != nil {
		helper.Render.RenderJsonFail(ctx, err)
		return
	}
	fmt.Println(jsoniter.MarshalToString(*fh))

	err = ctx.SaveUploadedFile(fh, "./statics/img/" + fh.Filename)
	if err != nil {
		helper.Render.RenderJsonFail(ctx, err)
		return
	}
	helper.Render.RenderJsonSucc(ctx, "ok")
}

