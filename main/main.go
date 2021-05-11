package main

import (
	sqlmodel "cmdb-go/commonorm"
	"cmdb-go/configmodel"
	"cmdb-go/router"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	configmodel.ConfigGet()
	sqlmodel.CheckDbconn(sqlmodel.Err)
	// sqlmodel.Sqlmodel()
	// 分离了router 调用routers下的router文件
	router.Router().Run(configmodel.ConfigGet().ServerPort)
}
