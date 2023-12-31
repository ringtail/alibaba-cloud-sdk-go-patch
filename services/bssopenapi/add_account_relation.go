package bssopenapi

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

// AddAccountRelation invokes the bssopenapi.AddAccountRelation API synchronously
func (client *Client) AddAccountRelation(request *AddAccountRelationRequest) (response *AddAccountRelationResponse, err error) {
	response = CreateAddAccountRelationResponse()
	err = client.DoAction(request, response)
	return
}

// AddAccountRelationWithChan invokes the bssopenapi.AddAccountRelation API asynchronously
func (client *Client) AddAccountRelationWithChan(request *AddAccountRelationRequest) (<-chan *AddAccountRelationResponse, <-chan error) {
	responseChan := make(chan *AddAccountRelationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddAccountRelation(request)
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

// AddAccountRelationWithCallback invokes the bssopenapi.AddAccountRelation API asynchronously
func (client *Client) AddAccountRelationWithCallback(request *AddAccountRelationRequest, callback func(response *AddAccountRelationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddAccountRelationResponse
		var err error
		defer close(result)
		response, err = client.AddAccountRelation(request)
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

// AddAccountRelationRequest is the request struct for api AddAccountRelation
type AddAccountRelationRequest struct {
	*requests.RpcRequest
	ChildNick       string           `position:"Query" name:"ChildNick"`
	RelationType    string           `position:"Query" name:"RelationType"`
	ParentUserId    requests.Integer `position:"Query" name:"ParentUserId"`
	ChildUserId     requests.Integer `position:"Query" name:"ChildUserId"`
	RequestId       string           `position:"Query" name:"RequestId"`
	PermissionCodes *[]string        `position:"Query" name:"PermissionCodes"  type:"Repeated"`
	RoleCodes       *[]string        `position:"Query" name:"RoleCodes"  type:"Repeated"`
}

// AddAccountRelationResponse is the response struct for api AddAccountRelation
type AddAccountRelationResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateAddAccountRelationRequest creates a request to invoke AddAccountRelation API
func CreateAddAccountRelationRequest() (request *AddAccountRelationRequest) {
	request = &AddAccountRelationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "AddAccountRelation", "bssopenapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddAccountRelationResponse creates a response to parse from AddAccountRelation response
func CreateAddAccountRelationResponse() (response *AddAccountRelationResponse) {
	response = &AddAccountRelationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
