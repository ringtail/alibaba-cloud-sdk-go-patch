package dts

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

// CreateDedicatedClusterMonitorRule invokes the dts.CreateDedicatedClusterMonitorRule API synchronously
func (client *Client) CreateDedicatedClusterMonitorRule(request *CreateDedicatedClusterMonitorRuleRequest) (response *CreateDedicatedClusterMonitorRuleResponse, err error) {
	response = CreateCreateDedicatedClusterMonitorRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDedicatedClusterMonitorRuleWithChan invokes the dts.CreateDedicatedClusterMonitorRule API asynchronously
func (client *Client) CreateDedicatedClusterMonitorRuleWithChan(request *CreateDedicatedClusterMonitorRuleRequest) (<-chan *CreateDedicatedClusterMonitorRuleResponse, <-chan error) {
	responseChan := make(chan *CreateDedicatedClusterMonitorRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDedicatedClusterMonitorRule(request)
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

// CreateDedicatedClusterMonitorRuleWithCallback invokes the dts.CreateDedicatedClusterMonitorRule API asynchronously
func (client *Client) CreateDedicatedClusterMonitorRuleWithCallback(request *CreateDedicatedClusterMonitorRuleRequest, callback func(response *CreateDedicatedClusterMonitorRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDedicatedClusterMonitorRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateDedicatedClusterMonitorRule(request)
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

// CreateDedicatedClusterMonitorRuleRequest is the request struct for api CreateDedicatedClusterMonitorRule
type CreateDedicatedClusterMonitorRuleRequest struct {
	*requests.RpcRequest
	CpuAlarmThreshold  requests.Integer `position:"Query" name:"CpuAlarmThreshold"`
	Phones             string           `position:"Query" name:"Phones"`
	DedicatedClusterId string           `position:"Query" name:"DedicatedClusterId"`
	DiskAlarmThreshold requests.Integer `position:"Query" name:"DiskAlarmThreshold"`
	MemAlarmThreshold  requests.Integer `position:"Query" name:"MemAlarmThreshold"`
	DuAlarmThreshold   requests.Integer `position:"Query" name:"DuAlarmThreshold"`
	OwnerId            string           `position:"Query" name:"OwnerId"`
	NoticeSwitch       requests.Integer `position:"Query" name:"NoticeSwitch"`
	InstanceId         string           `position:"Query" name:"InstanceId"`
}

// CreateDedicatedClusterMonitorRuleResponse is the response struct for api CreateDedicatedClusterMonitorRule
type CreateDedicatedClusterMonitorRuleResponse struct {
	*responses.BaseResponse
	HttpStatusCode string `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	Success        string `json:"Success" xml:"Success"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
}

// CreateCreateDedicatedClusterMonitorRuleRequest creates a request to invoke CreateDedicatedClusterMonitorRule API
func CreateCreateDedicatedClusterMonitorRuleRequest() (request *CreateDedicatedClusterMonitorRuleRequest) {
	request = &CreateDedicatedClusterMonitorRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "CreateDedicatedClusterMonitorRule", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDedicatedClusterMonitorRuleResponse creates a response to parse from CreateDedicatedClusterMonitorRule response
func CreateCreateDedicatedClusterMonitorRuleResponse() (response *CreateDedicatedClusterMonitorRuleResponse) {
	response = &CreateDedicatedClusterMonitorRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
