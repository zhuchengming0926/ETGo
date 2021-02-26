/**
 * @File: httpPostMultiForm.go
 * @Author: zhuchengming
 * @Description:
 * @Date: 2021/2/26 17:34
 */

package v1

import (
	"ETGo/components/utils"
	"ETGo/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func HttpPostMultiFormDemo(ctx *gin.Context)  {
	form, err := ctx.MultipartForm()
	if err != nil {
		helper.Render.RenderJsonFail(ctx, err)
		return
	}
	files := form.File["awardFile"]

	fileSavePath := "./statics/img/"
	if !utils.Exists(fileSavePath) {
		err := os.Mkdir(fileSavePath, os.ModePerm)
		if err != nil {
			helper.Render.RenderJsonFail(ctx, err)
			return
		}
	}

	for _, f := range files {
		fmt.Println(f.Filename)
		ctx.SaveUploadedFile(f, strings.Join([]string{fileSavePath, f.Filename}, ""))
	}
	helper.Render.RenderJsonSucc(ctx, "ok")
}


