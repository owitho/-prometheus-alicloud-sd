package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
	response = CreateDescribeClustersResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeClustersWithChan(request *DescribeClustersRequest) (<-chan *DescribeClustersResponse, <-chan error) {
	responseChan := make(chan *DescribeClustersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClusters(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeClustersWithCallback(request *DescribeClustersRequest, callback func(response *DescribeClustersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClustersResponse
		var err error
		defer close(result)
		response, err = client.DescribeClusters(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type DescribeClustersRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeClustersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Clusters  struct {
		Cluster []struct {
			ClusterId string `json:"ClusterId" xml:"ClusterId"`
		} `json:"Cluster" xml:"Cluster"`
	} `json:"Clusters" xml:"Clusters"`
}

func CreateDescribeClustersRequest() (request *DescribeClustersRequest) {
	request = &DescribeClustersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeClusters", "ecs", "openAPI")
	return
}

func CreateDescribeClustersResponse() (response *DescribeClustersResponse) {
	response = &DescribeClustersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
