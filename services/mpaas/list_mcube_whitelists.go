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

// ListMcubeWhitelists invokes the mpaas.ListMcubeWhitelists API synchronously
func (client *Client) ListMcubeWhitelists(request *ListMcubeWhitelistsRequest) (response *ListMcubeWhitelistsResponse, err error) {
	response = CreateListMcubeWhitelistsResponse()
	err = client.DoAction(request, response)
	return
}

// ListMcubeWhitelistsWithChan invokes the mpaas.ListMcubeWhitelists API asynchronously
func (client *Client) ListMcubeWhitelistsWithChan(request *ListMcubeWhitelistsRequest) (<-chan *ListMcubeWhitelistsResponse, <-chan error) {
	responseChan := make(chan *ListMcubeWhitelistsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMcubeWhitelists(request)
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

// ListMcubeWhitelistsWithCallback invokes the mpaas.ListMcubeWhitelists API asynchronously
func (client *Client) ListMcubeWhitelistsWithCallback(request *ListMcubeWhitelistsRequest, callback func(response *ListMcubeWhitelistsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMcubeWhitelistsResponse
		var err error
		defer close(result)
		response, err = client.ListMcubeWhitelists(request)
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

// ListMcubeWhitelistsRequest is the request struct for api ListMcubeWhitelists
type ListMcubeWhitelistsRequest struct {
	*requests.RpcRequest
	TenantId    string `position:"Body" name:"TenantId"`
	AppId       string `position:"Body" name:"AppId"`
	WorkspaceId string `position:"Body" name:"WorkspaceId"`
}

// ListMcubeWhitelistsResponse is the response struct for api ListMcubeWhitelists
type ListMcubeWhitelistsResponse struct {
	*responses.BaseResponse
	ResultMessage       string              `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode          string              `json:"ResultCode" xml:"ResultCode"`
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	ListWhitelistResult ListWhitelistResult `json:"ListWhitelistResult" xml:"ListWhitelistResult"`
}

// CreateListMcubeWhitelistsRequest creates a request to invoke ListMcubeWhitelists API
func CreateListMcubeWhitelistsRequest() (request *ListMcubeWhitelistsRequest) {
	request = &ListMcubeWhitelistsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "ListMcubeWhitelists", "", "")
	request.Method = requests.POST
	return
}

// CreateListMcubeWhitelistsResponse creates a response to parse from ListMcubeWhitelists response
func CreateListMcubeWhitelistsResponse() (response *ListMcubeWhitelistsResponse) {
	response = &ListMcubeWhitelistsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
