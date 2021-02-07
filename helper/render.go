/**************************************************************************
 * Copyright (c) 2020 Zuoyebang Inc. All Rights Reserved
 **************************************************************************/

/**
 * @File: render.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description:
 * @Date: 2021/2/7 20:11
 */

package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Render RenderS

type RenderS struct {
}

type DefaultRender struct {
	ErrNo  int         `json:"errNo"`
	ErrMsg string      `json:"errStr"`
	Data   interface{} `json:"data"`
}

func (r RenderS) RenderJson(ctx *gin.Context, code int, msg string, data interface{}) {
	renderJson := DefaultRender{code, msg, data}
	ctx.JSON(http.StatusOK, renderJson)
}

func (r RenderS) RenderJsonSucc(ctx *gin.Context, data interface{}) {
	renderJson := DefaultRender{0, "succ", data}
	ctx.JSON(http.StatusOK, renderJson)
}