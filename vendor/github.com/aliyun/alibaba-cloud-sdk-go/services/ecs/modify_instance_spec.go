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

func (client *Client) ModifyInstanceSpec(request *ModifyInstanceSpecRequest) (response *ModifyInstanceSpecResponse, err error) {
	response = CreateModifyInstanceSpecResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyInstanceSpecWithChan(request *ModifyInstanceSpecRequest) (<-chan *ModifyInstanceSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceSpec(request)
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

func (client *Client) ModifyInstanceSpecWithCallback(request *ModifyInstanceSpecRequest, callback func(response *ModifyInstanceSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceSpec(request)
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

type ModifyInstanceSpecRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                  requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount             string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken                      string           `position:"Query" name:"ClientToken"`
	AllowMigrateAcrossZone           requests.Boolean `position:"Query" name:"AllowMigrateAcrossZone"`
	OwnerAccount                     string           `position:"Query" name:"OwnerAccount"`
	InternetMaxBandwidthOut          requests.Integer `position:"Query" name:"InternetMaxBandwidthOut"`
	OwnerId                          requests.Integer `position:"Query" name:"OwnerId"`
	TemporaryInternetMaxBandwidthOut requests.Integer `position:"Query" name:"Temporary.InternetMaxBandwidthOut"`
	SystemDiskCategory               string           `position:"Query" name:"SystemDisk.Category"`
	TemporaryStartTime               string           `position:"Query" name:"Temporary.StartTime"`
	Async                            requests.Boolean `position:"Query" name:"Async"`
	InstanceId                       string           `position:"Query" name:"InstanceId"`
	InstanceType                     string           `position:"Query" name:"InstanceType"`
	TemporaryEndTime                 string           `position:"Query" name:"Temporary.EndTime"`
	InternetMaxBandwidthIn           requests.Integer `position:"Query" name:"InternetMaxBandwidthIn"`
}

type ModifyInstanceSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyInstanceSpecRequest() (request *ModifyInstanceSpecRequest) {
	request = &ModifyInstanceSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceSpec", "ecs", "openAPI")
	return
}

func CreateModifyInstanceSpecResponse() (response *ModifyInstanceSpecResponse) {
	response = &ModifyInstanceSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
