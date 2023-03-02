package ltl

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

// CreateMPCoSPhaseGroup invokes the ltl.CreateMPCoSPhaseGroup API synchronously
func (client *Client) CreateMPCoSPhaseGroup(request *CreateMPCoSPhaseGroupRequest) (response *CreateMPCoSPhaseGroupResponse, err error) {
	response = CreateCreateMPCoSPhaseGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMPCoSPhaseGroupWithChan invokes the ltl.CreateMPCoSPhaseGroup API asynchronously
func (client *Client) CreateMPCoSPhaseGroupWithChan(request *CreateMPCoSPhaseGroupRequest) (<-chan *CreateMPCoSPhaseGroupResponse, <-chan error) {
	responseChan := make(chan *CreateMPCoSPhaseGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMPCoSPhaseGroup(request)
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

// CreateMPCoSPhaseGroupWithCallback invokes the ltl.CreateMPCoSPhaseGroup API asynchronously
func (client *Client) CreateMPCoSPhaseGroupWithCallback(request *CreateMPCoSPhaseGroupRequest, callback func(response *CreateMPCoSPhaseGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMPCoSPhaseGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateMPCoSPhaseGroup(request)
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

// CreateMPCoSPhaseGroupRequest is the request struct for api CreateMPCoSPhaseGroup
type CreateMPCoSPhaseGroupRequest struct {
	*requests.RpcRequest
	Name       string `position:"Query" name:"Name"`
	ApiVersion string `position:"Query" name:"ApiVersion"`
	Remark     string `position:"Query" name:"Remark"`
	BizChainId string `position:"Query" name:"BizChainId"`
}

// CreateMPCoSPhaseGroupResponse is the response struct for api CreateMPCoSPhaseGroup
type CreateMPCoSPhaseGroupResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateMPCoSPhaseGroupRequest creates a request to invoke CreateMPCoSPhaseGroup API
func CreateCreateMPCoSPhaseGroupRequest() (request *CreateMPCoSPhaseGroupRequest) {
	request = &CreateMPCoSPhaseGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ltl", "2019-05-10", "CreateMPCoSPhaseGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateMPCoSPhaseGroupResponse creates a response to parse from CreateMPCoSPhaseGroup response
func CreateCreateMPCoSPhaseGroupResponse() (response *CreateMPCoSPhaseGroupResponse) {
	response = &CreateMPCoSPhaseGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
