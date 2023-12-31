package dms_enterprise

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

// CreateProxyAccess invokes the dms_enterprise.CreateProxyAccess API synchronously
func (client *Client) CreateProxyAccess(request *CreateProxyAccessRequest) (response *CreateProxyAccessResponse, err error) {
	response = CreateCreateProxyAccessResponse()
	err = client.DoAction(request, response)
	return
}

// CreateProxyAccessWithChan invokes the dms_enterprise.CreateProxyAccess API asynchronously
func (client *Client) CreateProxyAccessWithChan(request *CreateProxyAccessRequest) (<-chan *CreateProxyAccessResponse, <-chan error) {
	responseChan := make(chan *CreateProxyAccessResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateProxyAccess(request)
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

// CreateProxyAccessWithCallback invokes the dms_enterprise.CreateProxyAccess API asynchronously
func (client *Client) CreateProxyAccessWithCallback(request *CreateProxyAccessRequest, callback func(response *CreateProxyAccessResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateProxyAccessResponse
		var err error
		defer close(result)
		response, err = client.CreateProxyAccess(request)
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

// CreateProxyAccessRequest is the request struct for api CreateProxyAccess
type CreateProxyAccessRequest struct {
	*requests.RpcRequest
	IndepAccount  string           `position:"Query" name:"IndepAccount"`
	UserId        requests.Integer `position:"Query" name:"UserId"`
	Tid           requests.Integer `position:"Query" name:"Tid"`
	ProxyId       requests.Integer `position:"Query" name:"ProxyId"`
	IndepPassword string           `position:"Query" name:"IndepPassword"`
}

// CreateProxyAccessResponse is the response struct for api CreateProxyAccess
type CreateProxyAccessResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Success       bool   `json:"Success" xml:"Success"`
	ErrorMessage  string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode     string `json:"ErrorCode" xml:"ErrorCode"`
	ProxyAccessId int64  `json:"ProxyAccessId" xml:"ProxyAccessId"`
}

// CreateCreateProxyAccessRequest creates a request to invoke CreateProxyAccess API
func CreateCreateProxyAccessRequest() (request *CreateProxyAccessRequest) {
	request = &CreateProxyAccessRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "CreateProxyAccess", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateProxyAccessResponse creates a response to parse from CreateProxyAccess response
func CreateCreateProxyAccessResponse() (response *CreateProxyAccessResponse) {
	response = &CreateProxyAccessResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
