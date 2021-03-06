package exporter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func DiscoveryAlicloudMysql(filePath, exporterType string) {
	var nodeinfolist []NodeInfo
	var nodeinfo NodeInfo
	var flag bool = false

	totalcount := GetInstancesTotalCount(exporterType)
	ecsClient := NewEcsClient()

	for i := 0; i <= (totalcount / PAGESIZE); i++ {
		request := ecs.CreateDescribeInstancesRequest()
		request.PageSize = requests.NewInteger(PAGESIZE)
		request.PageNumber = requests.NewInteger(i + 1)
		request.Tag1Key = "Monitoring"
		request.Tag1Value = "true"
		request.Tag2Key = "Component"
		request.Tag2Value = "mysql"
		request.Status = "Running"
		response, err := ecsClient.DescribeInstances(request)
		if err != nil {
			panic(err)
		}

		for _, v := range response.Instances.Instance {
			//fmt.Println(x)
			for _, y := range v.Tags.Tag {
				if y.TagKey == "Env" {
					nodeinfo.Labels.Env = y.TagValue
				} else if y.TagKey == "Job" {
					nodeinfo.Labels.Job = y.TagValue
				} else if y.TagKey == "Component" {
					nodeinfo.Labels.Component = y.TagValue
				} else if y.TagKey == "Service" {
					nodeinfo.Labels.Service = y.TagValue
				}

				if nodeinfo.Labels.Job == "" {
					nodeinfo.Labels.Job = "mysql"
				}
			}

			for m, n := range nodeinfolist {
				if n.Labels.Env == nodeinfo.Labels.Env && n.Labels.Service == nodeinfo.Labels.Service && n.Labels.Service != "" {
					nodeinfolist[m].Targets = append(nodeinfolist[m].Targets, v.InstanceName+":9104")
					flag = true
					break
				} else {
					flag = false
				}
			}

			if flag == false {
				nodeinfo.Targets = append(nodeinfo.Targets, v.InstanceName+":9104")
				nodeinfolist = append(nodeinfolist, nodeinfo)
			}
			nodeinfo.Targets = nil
			nodeinfo.Labels.Env = ""
			nodeinfo.Labels.Job = ""
			nodeinfo.Labels.Component = ""
			nodeinfo.Labels.Service = ""
		}
	}
	jsonScrapeConfig, err := json.MarshalIndent(nodeinfolist, "", "\t")
	if err != nil {
		fmt.Println("json err", err)
	}
	err = ioutil.WriteFile(filePath, jsonScrapeConfig, 0644)
	if err != nil {
		panic(err)
	}
}
