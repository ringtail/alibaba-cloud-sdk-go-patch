package mpaas

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

// StartUserAppAsyncEnhanceInMsa invokes the mpaas.StartUserAppAsyncEnhanceInMsa API synchronously
func (client *Client) StartUserAppAsyncEnhanceInMsa(request *StartUserAppAsyncEnhanceInMsaRequest) (response *StartUserAppAsyncEnhanceInMsaResponse, err error) {
	response = CreateStartUserAppAsyncEnhanceInMsaResponse()
	err = client.DoAction(request, response)
	return
}

// StartUserAppAsyncEnhanceInMsaWithChan invokes the mpaas.StartUserAppAsyncEnhanceInMsa API asynchronously
func (client *Client) StartUserAppAsyncEnhanceInMsaWithChan(request *StartUserAppAsyncEnhanceInMsaRequest) (<-chan *StartUserAppAsyncEnhanceInMsaResponse, <-chan error) {
	responseChan := make(chan *StartUserAppAsyncEnhanceInMsaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartUserAppAsyncEnhanceInMsa(request)
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

// StartUserAppAsyncEnhanceInMsaWithCallback invokes the mpaas.StartUserAppAsyncEnhanceInMsa API asynchronously
func (client *Client) StartUserAppAsyncEnhanceInMsaWithCallback(request *StartUserAppAsyncEnhanceInMsaRequest, callback func(response *StartUserAppAsyncEnhanceInMsaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartUserAppAsyncEnhanceInMsaResponse
		var err error
		defer close(result)
		response, err = client.StartUserAppAsyncEnhanceInMsa(request)
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

// StartUserAppAsyncEnhanceInMsaRequest is the request struct for api StartUserAppAsyncEnhanceInMsa
type StartUserAppAsyncEnhanceInMsaRequest struct {
	*requests.RpcRequest
	RunMode             string           `position:"Body" name:"RunMode"`
	Classes             string           `position:"Body" name:"Classes"`
	NativeDebugger      requests.Integer `position:"Body" name:"NativeDebugger"`
	DalvikDebugger      requests.Integer `position:"Body" name:"DalvikDebugger"`
	TotalSwitch         requests.Boolean `position:"Body" name:"TotalSwitch"`
	Root                requests.Integer `position:"Body" name:"Root"`
	TenantId            string           `position:"Body" name:"TenantId"`
	JavaHook            requests.Integer `position:"Body" name:"JavaHook"`
	Id                  requests.Integer `position:"Body" name:"Id"`
	ApkProtector        requests.Boolean `position:"Body" name:"ApkProtector"`
	TaskType            string           `position:"Body" name:"TaskType"`
	PackageTampered     requests.Integer `position:"Body" name:"PackageTampered"`
	SoFileList          string           `position:"Body" name:"SoFileList"`
	MemoryDump          requests.Integer `position:"Body" name:"MemoryDump"`
	EmulatorEnvironment requests.Integer `position:"Body" name:"EmulatorEnvironment"`
	NativeHook          requests.Integer `position:"Body" name:"NativeHook"`
	AppId               string           `position:"Body" name:"AppId"`
	AssetsFileList      string           `position:"Body" name:"AssetsFileList"`
	WorkspaceId         string           `position:"Body" name:"WorkspaceId"`
}

// StartUserAppAsyncEnhanceInMsaResponse is the response struct for api StartUserAppAsyncEnhanceInMsa
type StartUserAppAsyncEnhanceInMsaResponse struct {
	*responses.BaseResponse
	ResultMessage string                                       `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string                                       `json:"ResultCode" xml:"ResultCode"`
	RequestId     string                                       `json:"RequestId" xml:"RequestId"`
	ResultContent ResultContentInStartUserAppAsyncEnhanceInMsa `json:"ResultContent" xml:"ResultContent"`
}

// CreateStartUserAppAsyncEnhanceInMsaRequest creates a request to invoke StartUserAppAsyncEnhanceInMsa API
func CreateStartUserAppAsyncEnhanceInMsaRequest() (request *StartUserAppAsyncEnhanceInMsaRequest) {
	request = &StartUserAppAsyncEnhanceInMsaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "StartUserAppAsyncEnhanceInMsa", "", "")
	request.Method = requests.POST
	return
}

// CreateStartUserAppAsyncEnhanceInMsaResponse creates a response to parse from StartUserAppAsyncEnhanceInMsa response
func CreateStartUserAppAsyncEnhanceInMsaResponse() (response *StartUserAppAsyncEnhanceInMsaResponse) {
	response = &StartUserAppAsyncEnhanceInMsaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
