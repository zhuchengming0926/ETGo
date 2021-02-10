

/**
 * @File: render.go
 * @Author: zhuchengming
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

func (r RenderS) RenderJsonAbort(ctx *gin.Context) {
	var renderJson  = DefaultRender{
		ErrNo:    501,
		ErrMsg:   "system internal error",
		Data:     gin.H{},
	}

	ctx.AbortWithStatusJSON(http.StatusOK, renderJson)
	return
}

func (r RenderS) RenderJsonFail(ctx *gin.Context, err error) {
	var renderJson = DefaultRender{
		ErrNo:    2001,
		ErrMsg:   err.Error(),
		Data:     gin.H{},
	}

	ctx.JSON(http.StatusOK, renderJson)
	return
}