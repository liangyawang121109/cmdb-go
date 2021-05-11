package router

import (
	"cmdb-go/cmdb"
	"cmdb-go/handcors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	R := gin.Default()
	R.Use(handcors.Cors())
	cmdbGroup := R.Group("/cmdb")
	{
		cmdbGroup.GET("/Ecsinfo", cmdb.EcsInfoGet)
		cmdbGroup.GET("/Slbinfo", cmdb.SlbInfoGet)
		cmdbGroup.GET("/containerInfo", cmdb.GetContainerInfo)
		cmdbGroup.GET("/queryContainer/:containerTag", cmdb.FilterContainer)
		cmdbGroup.DELETE("/containerDelete/:containerName/:containerNode/:containerIds", cmdb.DeleteContainer)
		cmdbGroup.DELETE("/delContainerOnTag/:containerTag/:containerEnv", cmdb.DelContainerOnTag)
		cmdbGroup.PUT("/containerRollover/:containerName", cmdb.RolloverContainer)
		cmdbGroup.PUT("/containerUpdate/:containerNode/:containerName/:containerPort", cmdb.UpdateContainer)
		cmdbGroup.PUT("/containerMonitor/:containerName/:containerStatus", cmdb.UpdateContainerStatus)
		cmdbGroup.POST("/containerCreate/:containerEnv/:runUser/:containerPort/:containerTag/:containerName/:containerProject/:containerNode/:containerIds/:containerStatus", cmdb.CreateContainerInfo)

	}
	R.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound,gin.H{
			"message":"NotFound",
		})
	})
	return R
}
