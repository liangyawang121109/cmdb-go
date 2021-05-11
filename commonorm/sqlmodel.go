package sqlmodel

import (
	"cmdb-go/configmodel"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

type ContainerInfo struct {
	gorm.Model
	ContainerEnv string
	RunUser string
	ContainerProject string
	ContainerTag string
	ContainerName string
	ContainerPort string
	ContainerNode string
	ContainerIds string
	ContainerStatus bool
}

func Dbconnect() (db interface{},err error){
	config := configmodel.ConfigGet()
	databaseUrl := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",config.User,config.Passwd,config.DatabaseUrl,config.Database)
	Db, DbErr := gorm.Open(config.Name, databaseUrl)
	return Db,DbErr
}

var Db,Err = Dbconnect()
var DbConnection = Db.(*gorm.DB)
func SqlGet() interface{} {
	var container  []ContainerInfo
	DbConnection.Find(&container)
	return container
}

func CheckDbconn(err error)  {

	if err != nil {
		fmt.Println("数据库链接异常",err)
		os.Exit(1)

	}
}


func Sqlmodel() {
	//数据表初始化
	defer DbConnection.Close()
	if Err != nil {
		fmt.Println("数据库连接异常", Err)
		os.Exit(1)
	}
	DbConnection.AutoMigrate(&ContainerInfo{})
}

func SqlCreate(containerenv, runuser, containerport, containertag, containername, containerporject,containernode,containerids string,containerstatus bool)  {
	u1 := ContainerInfo{
		ContainerEnv: containerenv,
		RunUser: runuser,
		ContainerPort: containerport,
		ContainerTag: containertag,
		ContainerName: containername,
		ContainerProject: containerporject,
		ContainerNode: containernode,
		ContainerIds: containerids,
		ContainerStatus: containerstatus}
	DbConnection.Create(&u1)
}


func ContainerStatusMonitor(containerName string, containerStatus bool)  {
	var containerinfos ContainerInfo
	DbConnection.Debug().Model(&containerinfos).Where("container_name=?", containerName).Update("container_status", containerStatus)
}

func SqlUpdate(containerNode, containerName, containerPort string) {
	var containerinfos ContainerInfo
	DbConnection.Debug().Model(&containerinfos).Where("container_name=?", containerName).Updates(map[string]interface{}{"container_port":containerPort,"container_node": containerNode})
}

func SqlRollover(containerName string) {
	var containerinfos ContainerInfo
	DbConnection.Debug().Model(&containerinfos).Where("container_name=?", containerName).Update("container_name", containerName)
}

func SqlDelete(containerName string)  {
	//使用了gorm的model后字都会有deletedat字段 那么将自动获得软删除功能
	//按照匹配到的数据行删除数据 根据容器名称 默认gorm是软删除 也就是在数据库的deletedat字段添加时间
	//那么gorm默认就不会查询这个字段不为空的数据 这就是软删除
	//db.Debug().Where("container_name=?", containerName).Delete(ContainerInfo{})

	//物理删除 也就是硬删除
	DbConnection.Debug().Unscoped().Where("container_name=?", containerName).Delete(ContainerInfo{})
}

func SqlFind(containerTag string) interface{} {
	var querycontainerinfo []ContainerInfo
	DbConnection.Debug().Where("container_tag=?", containerTag).Find(&querycontainerinfo)
	return querycontainerinfo
}

//func main() {
//	fmt.Println(SqlGet())
//}
