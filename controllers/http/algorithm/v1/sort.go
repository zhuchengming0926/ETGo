/**
 * @File: sort.go
 * @Author: zhuchengming
 * @Description:
 * @Date: 2021/2/18 20:07
 */

package v1

import (
	"ETGo/components/dto/dtoAlgorithm"
	"ETGo/elog"
	"ETGo/helper"
	"ETGo/service/svAlgorithm"
	"github.com/gin-gonic/gin"
)

func VariousSort(ctx *gin.Context)  {
	req := dtoAlgorithm.SortReq{}
	if errParams := ctx.ShouldBindJSON(&req); errParams != nil {
		elog.Warnf(ctx, "参数有错, err is %s", errParams.Error())
		helper.Render.RenderJsonFail(ctx, errParams)
		return
	}
	res, err := svAlgorithm.VariousSort(&req)
	if err != nil {
		helper.Render.RenderJsonFail(ctx, err)
	}
	helper.Render.RenderJsonSucc(ctx, res)
	return
}

