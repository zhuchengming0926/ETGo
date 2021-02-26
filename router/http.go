

/**
 * @File: http.go
 * @Author: zhuchengming
 * @Description:
 * @Date: 2021/2/7 20:06
 */

package router

import (
	agv1 "ETGo/controllers/http/algorithm/v1"
	uqv1 "ETGo/controllers/http/userQuestion/v1"
	httpv1 "ETGo/controllers/http/variousHttp/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Http(router *gin.Engine)  {
	routerGroupRoot := router.Group("/et")
	{
		routerGroupVer := routerGroupRoot.Group("/v1")
		{
			//业务
			routerGroupModuleUQ := routerGroupVer.Group("/userquestion")
			{
				routerGroupModuleUQ.POST("/getDetail", uqv1.GetRecordDetail)
			}

			//算法
			routerGroupAlgorithm := routerGroupVer.Group("/algorithm")
			{
				routerGroupAlgorithm.POST("/sort", agv1.VariousSort)
			}

			//各类http请求参数解析
			routerGroupHttp := routerGroupVer.Group("/variousHttp")
			{
				routerGroupHttp.GET("/getdemo", httpv1.HttpGetDemo)
				routerGroupHttp.POST("/postjson", httpv1.HttpPostJsonDemo)
				routerGroupHttp.POST("/postmultiform", httpv1.HttpPostMultiFormDemo)
				routerGroupHttp.POST("/postwwwurlencode", httpv1.HttpPostWwwFormUrlEncodeDemo)
				routerGroupHttp.POST("/postxml", httpv1.HttpPostXmlDemo)
			}
			
			//首页测试
			routerGroupVer.GET("/index", func(ctx *gin.Context) {
				ctx.HTML(http.StatusOK, "index.html", nil)
			})
			
			//上传文件测试
			routerGroupVer.POST("/uploadFile", func(ctx *gin.Context) {
				//获取表单数据
				f, err := ctx.FormFile("f1")
				if err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"error":err})
					return
				} else {
					ctx.SaveUploadedFile(f, f.Filename)
					ctx.JSON(http.StatusOK, gin.H{"message":"OK"})
				}
			})
		}
	}
}

