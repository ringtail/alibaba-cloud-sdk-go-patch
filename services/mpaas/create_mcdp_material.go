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

// CreateMcdpMaterial invokes the mpaas.CreateMcdpMaterial API synchronously
func (client *Client) CreateMcdpMaterial(request *CreateMcdpMaterialRequest) (response *CreateMcdpMaterialResponse, err error) {
	response = CreateCreateMcdpMaterialResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMcdpMaterialWithChan invokes the mpaas.CreateMcdpMaterial API asynchronously
func (client *Client) CreateMcdpMaterialWithChan(request *CreateMcdpMaterialRequest) (<-chan *CreateMcdpMaterialResponse, <-chan error) {
	responseChan := make(chan *CreateMcdpMaterialResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMcdpMaterial(request)
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

// CreateMcdpMaterialWithCallback invokes the mpaas.CreateMcdpMaterial API asynchronously
func (client *Client) CreateMcdpMaterialWithCallback(request *CreateMcdpMaterialRequest, callback func(response *CreateMcdpMaterialResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMcdpMaterialResponse
		var err error
		defer close(result)
		response, err = client.CreateMcdpMaterial(request)
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

// CreateMcdpMaterialRequest is the request struct for api CreateMcdpMaterial
type CreateMcdpMaterialRequest struct {
	*requests.RpcRequest
	TenantId                                 string `position:"Body" name:"TenantId"`
	MpaasMappcenterMcdpMaterialCreateJsonStr string `position:"Body" name:"MpaasMappcenterMcdpMaterialCreateJsonStr"`
	AppId                                    string `position:"Body" name:"AppId"`
	WorkspaceId                              string `position:"Body" name:"WorkspaceId"`
}

// CreateMcdpMaterialResponse is the response struct for api CreateMcdpMaterial
type CreateMcdpMaterialResponse struct {
	*responses.BaseResponse
	ResultMessage string                            `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string                            `json:"ResultCode" xml:"ResultCode"`
	RequestId     string                            `json:"RequestId" xml:"RequestId"`
	ResultContent ResultContentInCreateMcdpMaterial `json:"ResultContent" xml:"ResultContent"`
}

// CreateCreateMcdpMaterialRequest creates a request to invoke CreateMcdpMaterial API
func CreateCreateMcdpMaterialRequest() (request *CreateMcdpMaterialRequest) {
	request = &CreateMcdpMaterialRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "CreateMcdpMaterial", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateMcdpMaterialResponse creates a response to parse from CreateMcdpMaterial response
func CreateCreateMcdpMaterialResponse() (response *CreateMcdpMaterialResponse) {
	response = &CreateMcdpMaterialResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
