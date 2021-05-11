package containerModel

import (
	"fmt"
	"net/http"
)

func ContainerDelete(containerNode,containerId string) {
	/*
	传递docker remote api 以停止正在运行的容器 并删除它
	 */
	deleteContainerUrl := fmt.Sprintf("http://%s:2375/containers/%s",containerNode,containerId)
	stopContainerUrl := fmt.Sprintf("http://%s:2375/containers/%s/stop",containerNode,containerId)
	stopContainer,_ := http.NewRequest("POST",stopContainerUrl,nil)
	deleteContainer,_ := http.NewRequest("DELETE",deleteContainerUrl,nil)
	stopRes,_ := http.DefaultClient.Do(stopContainer)
	deleteRes,_ := http.DefaultClient.Do(deleteContainer)
	defer stopRes.Body.Close()
	defer deleteRes.Body.Close()
	fmt.Println(deleteRes.StatusCode)
}
