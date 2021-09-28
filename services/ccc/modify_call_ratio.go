package ccc

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

// ModifyCallRatio invokes the ccc.ModifyCallRatio API synchronously
func (client *Client) ModifyCallRatio(request *ModifyCallRatioRequest) (response *ModifyCallRatioResponse, err error) {
	response = CreateModifyCallRatioResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyCallRatioWithChan invokes the ccc.ModifyCallRatio API asynchronously
func (client *Client) ModifyCallRatioWithChan(request *ModifyCallRatioRequest) (<-chan *ModifyCallRatioResponse, <-chan error) {
	responseChan := make(chan *ModifyCallRatioResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCallRatio(request)
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

// ModifyCallRatioWithCallback invokes the ccc.ModifyCallRatio API asynchronously
func (client *Client) ModifyCallRatioWithCallback(request *ModifyCallRatioRequest, callback func(response *ModifyCallRatioResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCallRatioResponse
		var err error
		defer close(result)
		response, err = client.ModifyCallRatio(request)
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

// ModifyCallRatioRequest is the request struct for api ModifyCallRatio
type ModifyCallRatioRequest struct {
	*requests.RpcRequest
	InstanceId   string           `position:"Query" name:"InstanceId"`
	SkillGroupId string           `position:"Query" name:"SkillGroupId"`
	JobGroupId   string           `position:"Query" name:"JobGroupId"`
	Ratio        requests.Integer `position:"Query" name:"Ratio"`
}

// ModifyCallRatioResponse is the response struct for api ModifyCallRatio
type ModifyCallRatioResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateModifyCallRatioRequest creates a request to invoke ModifyCallRatio API
func CreateModifyCallRatioRequest() (request *ModifyCallRatioRequest) {
	request = &ModifyCallRatioRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ModifyCallRatio", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyCallRatioResponse creates a response to parse from ModifyCallRatio response
func CreateModifyCallRatioResponse() (response *ModifyCallRatioResponse) {
	response = &ModifyCallRatioResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
