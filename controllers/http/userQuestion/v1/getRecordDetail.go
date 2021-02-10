/**************************************************************************
 *
 **************************************************************************/

/**
 * @File: getRecordDetail.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description:
 * @Date: 2021/2/8 16:56
 */

package v1

import (
	"ETGo/components/dto/dtoUserQuestion"
	"ETGo/elog"
	"ETGo/helper"
	"ETGo/service/svUserQuestion"
	"github.com/gin-gonic/gin"
)

func GetRecordDetail(ctx *gin.Context)  {
	req := dtoUserQuestion.DetailReq{}
	if errParams := ctx.ShouldBindJSON(&req); errParams != nil {
		elog.Warnf(ctx, "你爹出日志了吗")
		helper.Render.RenderJsonFail(ctx, errParams)
		return
	}
	res, err := svUserQuestion.UserFeedBackDetail(req.Id)
	if err != nil {
		helper.Render.RenderJsonFail(ctx, err)
	}
	helper.Render.RenderJsonSucc(ctx, res)
	return
}

