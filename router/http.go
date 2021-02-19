

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
	"github.com/gin-gonic/gin"
)

func Http(router *gin.Engine)  {
	routerGroupRoot := router.Group("/et")
	{
		routerGroupVer := routerGroupRoot.Group("/v1")
		{
			routerGroupModuleUQ := routerGroupVer.Group("/userquestion")
			{
				routerGroupModuleUQ.POST("/getDetail", uqv1.GetRecordDetail)
			}

			routerGroupAlgorithm := routerGroupVer.Group("/algorithm")
			{
				routerGroupAlgorithm.POST("/sort", agv1.VariousSort)
			}
		}
	}
}

