package sas

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

// OperateSuspiciousTargetConfig invokes the sas.OperateSuspiciousTargetConfig API synchronously
func (client *Client) OperateSuspiciousTargetConfig(request *OperateSuspiciousTargetConfigRequest) (response *OperateSuspiciousTargetConfigResponse, err error) {
	response = CreateOperateSuspiciousTargetConfigResponse()
	err = client.DoAction(request, response)
	return
}

// OperateSuspiciousTargetConfigWithChan invokes the sas.OperateSuspiciousTargetConfig API asynchronously
func (client *Client) OperateSuspiciousTargetConfigWithChan(request *OperateSuspiciousTargetConfigRequest) (<-chan *OperateSuspiciousTargetConfigResponse, <-chan error) {
	responseChan := make(chan *OperateSuspiciousTargetConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OperateSuspiciousTargetConfig(request)
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

// OperateSuspiciousTargetConfigWithCallback invokes the sas.OperateSuspiciousTargetConfig API asynchronously
func (client *Client) OperateSuspiciousTargetConfigWithCallback(request *OperateSuspiciousTargetConfigRequest, callback func(response *OperateSuspiciousTargetConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OperateSuspiciousTargetConfigResponse
		var err error
		defer close(result)
		response, err = client.OperateSuspiciousTargetConfig(request)
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

// OperateSuspiciousTargetConfigRequest is the request struct for api OperateSuspiciousTargetConfig
type OperateSuspiciousTargetConfigRequest struct {
	*requests.RpcRequest
	TargetType                 string `position:"Query" name:"TargetType"`
	Type                       string `position:"Query" name:"Type"`
	TargetOperations           string `position:"Query" name:"TargetOperations"`
	SourceIp                   string `position:"Query" name:"SourceIp"`
	Lang                       string `position:"Query" name:"Lang"`
	ResourceDirectoryAccountId string `position:"Query" name:"ResourceDirectoryAccountId"`
}

// OperateSuspiciousTargetConfigResponse is the response struct for api OperateSuspiciousTargetConfig
type OperateSuspiciousTargetConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateOperateSuspiciousTargetConfigRequest creates a request to invoke OperateSuspiciousTargetConfig API
func CreateOperateSuspiciousTargetConfigRequest() (request *OperateSuspiciousTargetConfigRequest) {
	request = &OperateSuspiciousTargetConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "OperateSuspiciousTargetConfig", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOperateSuspiciousTargetConfigResponse creates a response to parse from OperateSuspiciousTargetConfig response
func CreateOperateSuspiciousTargetConfigResponse() (response *OperateSuspiciousTargetConfigResponse) {
	response = &OperateSuspiciousTargetConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
