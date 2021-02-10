/**************************************************************************
 *
 **************************************************************************/

/**
 * @File: http.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description:
 * @Date: 2021/2/7 20:06
 */

package router

import (
	v1 "ETGo/controllers/http/userQuestion/v1"
	"github.com/gin-gonic/gin"
)

func Http(router *gin.Engine)  {
	routerGroupRoot := router.Group("/et")
	{
		routerGroupVer := routerGroupRoot.Group("/v1")
		{
			routerGroupModule := routerGroupVer.Group("/userquestion")
			{
				routerGroupModule.POST("/getDetail", v1.GetRecordDetail)
			}
		}
	}
}

