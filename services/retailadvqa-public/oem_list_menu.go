package retailadvqa_public

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

// OemListMenu invokes the retailadvqa_public.OemListMenu API synchronously
func (client *Client) OemListMenu(request *OemListMenuRequest) (response *OemListMenuResponse, err error) {
	response = CreateOemListMenuResponse()
	err = client.DoAction(request, response)
	return
}

// OemListMenuWithChan invokes the retailadvqa_public.OemListMenu API asynchronously
func (client *Client) OemListMenuWithChan(request *OemListMenuRequest) (<-chan *OemListMenuResponse, <-chan error) {
	responseChan := make(chan *OemListMenuResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OemListMenu(request)
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

// OemListMenuWithCallback invokes the retailadvqa_public.OemListMenu API asynchronously
func (client *Client) OemListMenuWithCallback(request *OemListMenuRequest, callback func(response *OemListMenuResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OemListMenuResponse
		var err error
		defer close(result)
		response, err = client.OemListMenu(request)
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

// OemListMenuRequest is the request struct for api OemListMenu
type OemListMenuRequest struct {
	*requests.RpcRequest
	AccessId    string `position:"Query" name:"AccessId"`
	TenantId    string `position:"Query" name:"TenantId"`
	RoleSign    string `position:"Query" name:"RoleSign"`
	Version     string `position:"Query" name:"Version"`
	WorkspaceId string `position:"Query" name:"WorkspaceId"`
}

// OemListMenuResponse is the response struct for api OemListMenu
type OemListMenuResponse struct {
	*responses.BaseResponse
	RequestId string                  `json:"RequestId" xml:"RequestId"`
	ErrorDesc string                  `json:"ErrorDesc" xml:"ErrorDesc"`
	TraceId   string                  `json:"TraceId" xml:"TraceId"`
	ErrorCode string                  `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool                    `json:"Success" xml:"Success"`
	Data      []DataItemInOemListMenu `json:"Data" xml:"Data"`
}

// CreateOemListMenuRequest creates a request to invoke OemListMenu API
func CreateOemListMenuRequest() (request *OemListMenuRequest) {
	request = &OemListMenuRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "OemListMenu", "", "")
	request.Method = requests.POST
	return
}

// CreateOemListMenuResponse creates a response to parse from OemListMenu response
func CreateOemListMenuResponse() (response *OemListMenuResponse) {
	response = &OemListMenuResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
