// This file is auto-generated, don't edit it. Thanks.
package alimodel

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs "github.com/alibabacloud-go/ecs-20140526/v2/client"
	slb20140515 "github.com/alibabacloud-go/slb-20140515/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

type InstanceInfo struct {
	InstanceName string
	InstanceId string
	InstanceCpu int32
	InstanceMemory int32
	InstanceIp string
}

type SlbInfo struct {
	SlbStatus string
	SlbName string
	SlbAddr string
	SlbCreateTime string
}

var config = &openapi.Config{}

var id string = "你的aliid"

var sec string = "你的alisec"

func  AliInstanceList() (info interface{}, _err error) {
	// 1. 初始化配置

	// 您的AccessKey ID
	config.AccessKeyId = &id
	// 您的AccessKey Secret
	config.AccessKeySecret = &sec
	//设置请求地址
	config.Endpoint = tea.String("ecs.aliyuncs.com")
	// 设置连接超时为5000毫秒
	config.ConnectTimeout = tea.Int(5000)
	// 设置读超时为5000毫秒
	config.ReadTimeout = tea.Int(5000)
	// 2. 初始化客户端
	client, _err := ecs.NewClient(config)
	if _err != nil {
		fmt.Println(_err)
	}

	regionIds := "cn-beijing"
	describeInstancesRequest := &ecs.DescribeInstancesRequest{
		PageSize: tea.Int32(100),
		RegionId: &regionIds,
	}
	resp, _err := client.DescribeInstances(describeInstancesRequest)
	if _err != nil {
		fmt.Println(_err)
	}

	instances := resp.Body.Instances.Instance
	var infos []InstanceInfo
	instanceinfo := InstanceInfo{}

	for _, instance := range instances {
		instanceinfo.InstanceName = tea.StringValue(instance.InstanceName)
		instanceinfo.InstanceId = tea.StringValue(instance.InstanceId)
		instanceinfo.InstanceIp = tea.StringValue(instance.VpcAttributes.PrivateIpAddress.IpAddress[0])
		instanceinfo.InstanceCpu = tea.Int32Value(instance.Cpu)
		instanceinfo.InstanceMemory = tea.Int32Value(instance.Memory)
		infos = append(infos, instanceinfo)

	}
	return infos,_err
}

func AliSlbInfo () (slblist interface{}, _err error) {
	// 您的AccessKey ID
	config.AccessKeyId = &id
	// 您的AccessKey Secret
	config.AccessKeySecret = &sec
	//设置请求地址
	config.Endpoint = tea.String("slb.aliyuncs.com")
	client, _err := slb20140515.NewClient(config)
	if _err != nil {
		fmt.Println(_err)
	}

	describeLoadBalancersRequest := &slb20140515.DescribeLoadBalancersRequest{
		RegionId: tea.String("cn-beijing"),

	}
	// 复制代码运行请自行打印 API 的返回值
	res, _err := client.DescribeLoadBalancers(describeLoadBalancersRequest)
	if _err != nil {
		fmt.Println(_err)
	}
	slbs := res.Body.LoadBalancers.LoadBalancer
	var Slbinfos []SlbInfo
	slbinfo := SlbInfo{}
	for _,v := range slbs {
		slbinfo.SlbName = tea.StringValue(v.LoadBalancerName)
		slbinfo.SlbAddr = tea.StringValue(v.Address)
		slbinfo.SlbStatus = tea.StringValue(v.LoadBalancerStatus)
		slbinfo.SlbCreateTime = tea.StringValue(v.CreateTime)
		Slbinfos = append(Slbinfos, slbinfo)
	}
	return Slbinfos, _err
}
