package cmdb

import (
	"cmdb-go/alimodel"
	sqlmodel "cmdb-go/commonorm"
	containermodel "cmdb-go/containermodel"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetContainerInfo(c *gin.Context) {
	c.JSON(200,gin.H{
		"data":sqlmodel.SqlGet(),
	})
}

func DeleteContainer(c *gin.Context)  {
	containername := c.Param("containerName")
	nodeIp := c.Param("containerNode")
	containerid := c.Param("containerIds")
	sqlmodel.SqlDelete(containername)
	containermodel.ContainerDelete(nodeIp,containerid)
}

func RolloverContainer(c *gin.Context)  {
	containername := c.Param("containerName")
	sqlmodel.SqlRollover(containername)

}

func UpdateContainer(c *gin.Context)  {
	containernode := c.Param("containerNode")
	fmt.Println(containernode)
	fmt.Printf("%T",containernode)
	containername := c.Param("containerName")
	containerport := c.Param("containerPort")
	sqlmodel.SqlUpdate(containernode, containername, containerport)

}

func UpdateContainerStatus(c *gin.Context)  {
	containername := c.Param("containerName")
	containerstatus := c.Param("containerStatus")
	ContianerStatus,_ := strconv.ParseBool(containerstatus)
	sqlmodel.ContainerStatusMonitor(containername, ContianerStatus)
}

func CreateContainerInfo(c *gin.Context)  {
	containerenv := c.Param("containerEnv")
	runuser := c.Param("runUser")
	containerport := c.Param("containerPort")
	containertag := c.Param("containerTag")
	containername := c.Param("containerName")
	containerporject := c.Param("containerProject")
	containernode := c.Param("containerNode")
	containerids := c.Param("containerIds")
	containerstatus := c.Param("containerStatus")
	ContianerStatus,_ := strconv.ParseBool(containerstatus)
	sqlmodel.SqlCreate(containerenv,runuser,containerport,containertag,containername,containerporject,containernode,containerids, ContianerStatus)
}

func FilterContainer(c *gin.Context)  {
	containertag := c.Param("containerTag")
	c.JSON(200,gin.H{
		"filterdata":sqlmodel.SqlFind(containertag),
	})
}

func EcsInfoGet(c *gin.Context)  {
	ecsInfo,err := alimodel.AliInstanceList()
	if err != nil {
		fmt.Println("Ecs数据获取异常",err)
	}
	c.JSON(200,gin.H{
		"data":ecsInfo,
	})
}

func SlbInfoGet(c *gin.Context)  {
	slbInfo,err := alimodel.AliSlbInfo()
	fmt.Println(slbInfo)
	if err != nil {
		fmt.Println("Slb数据获取异常",err)
	}
	c.JSON(200,gin.H{
		"data":slbInfo,
	})
}

func DelContainerOnTag(c *gin.Context) {
	containerTag := c.Param("containerTag")
	containerEnv := c.Param("containerEnv")
	containerData := sqlmodel.SqlGet().([]sqlmodel.ContainerInfo)
	for i := 0; i < len(containerData); i++ {
		if containerTag == containerData[i].ContainerTag && containerEnv == containerData[i].ContainerEnv {
			sqlmodel.SqlDelete(containerData[i].ContainerName)
			containermodel.ContainerDelete(containerData[i].ContainerNode,containerData[i].ContainerIds)
			fmt.Println("应该被删除的有", containerData[i])
		}
	}
}
