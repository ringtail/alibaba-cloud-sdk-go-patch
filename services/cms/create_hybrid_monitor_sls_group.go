package cms

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

// CreateHybridMonitorSLSGroup invokes the cms.CreateHybridMonitorSLSGroup API synchronously
func (client *Client) CreateHybridMonitorSLSGroup(request *CreateHybridMonitorSLSGroupRequest) (response *CreateHybridMonitorSLSGroupResponse, err error) {
	response = CreateCreateHybridMonitorSLSGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateHybridMonitorSLSGroupWithChan invokes the cms.CreateHybridMonitorSLSGroup API asynchronously
func (client *Client) CreateHybridMonitorSLSGroupWithChan(request *CreateHybridMonitorSLSGroupRequest) (<-chan *CreateHybridMonitorSLSGroupResponse, <-chan error) {
	responseChan := make(chan *CreateHybridMonitorSLSGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateHybridMonitorSLSGroup(request)
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

// CreateHybridMonitorSLSGroupWithCallback invokes the cms.CreateHybridMonitorSLSGroup API asynchronously
func (client *Client) CreateHybridMonitorSLSGroupWithCallback(request *CreateHybridMonitorSLSGroupRequest, callback func(response *CreateHybridMonitorSLSGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateHybridMonitorSLSGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateHybridMonitorSLSGroup(request)
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

// CreateHybridMonitorSLSGroupRequest is the request struct for api CreateHybridMonitorSLSGroup
type CreateHybridMonitorSLSGroupRequest struct {
	*requests.RpcRequest
	SLSGroupDescription string                                       `position:"Query" name:"SLSGroupDescription"`
	SLSGroupConfig      *[]CreateHybridMonitorSLSGroupSLSGroupConfig `position:"Query" name:"SLSGroupConfig"  type:"Repeated"`
	SLSGroupName        string                                       `position:"Query" name:"SLSGroupName"`
}

// CreateHybridMonitorSLSGroupSLSGroupConfig is a repeated param struct in CreateHybridMonitorSLSGroupRequest
type CreateHybridMonitorSLSGroupSLSGroupConfig struct {
	SLSLogstore string `name:"SLSLogstore"`
	SLSUserId   string `name:"SLSUserId"`
	SLSProject  string `name:"SLSProject"`
	SLSRegion   string `name:"SLSRegion"`
}

// CreateHybridMonitorSLSGroupResponse is the response struct for api CreateHybridMonitorSLSGroup
type CreateHybridMonitorSLSGroupResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
}

// CreateCreateHybridMonitorSLSGroupRequest creates a request to invoke CreateHybridMonitorSLSGroup API
func CreateCreateHybridMonitorSLSGroupRequest() (request *CreateHybridMonitorSLSGroupRequest) {
	request = &CreateHybridMonitorSLSGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "CreateHybridMonitorSLSGroup", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateHybridMonitorSLSGroupResponse creates a response to parse from CreateHybridMonitorSLSGroup response
func CreateCreateHybridMonitorSLSGroupResponse() (response *CreateHybridMonitorSLSGroupResponse) {
	response = &CreateHybridMonitorSLSGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
