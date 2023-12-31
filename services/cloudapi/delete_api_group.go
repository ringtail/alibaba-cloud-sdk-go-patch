package cloudapi

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

// DeleteApiGroup invokes the cloudapi.DeleteApiGroup API synchronously
func (client *Client) DeleteApiGroup(request *DeleteApiGroupRequest) (response *DeleteApiGroupResponse, err error) {
	response = CreateDeleteApiGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteApiGroupWithChan invokes the cloudapi.DeleteApiGroup API asynchronously
func (client *Client) DeleteApiGroupWithChan(request *DeleteApiGroupRequest) (<-chan *DeleteApiGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteApiGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteApiGroup(request)
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

// DeleteApiGroupWithCallback invokes the cloudapi.DeleteApiGroup API asynchronously
func (client *Client) DeleteApiGroupWithCallback(request *DeleteApiGroupRequest, callback func(response *DeleteApiGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteApiGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteApiGroup(request)
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

// DeleteApiGroupRequest is the request struct for api DeleteApiGroup
type DeleteApiGroupRequest struct {
	*requests.RpcRequest
	GroupId       string               `position:"Query" name:"GroupId"`
	SecurityToken string               `position:"Query" name:"SecurityToken"`
	Tag           *[]DeleteApiGroupTag `position:"Query" name:"Tag"  type:"Repeated"`
}

// DeleteApiGroupTag is a repeated param struct in DeleteApiGroupRequest
type DeleteApiGroupTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DeleteApiGroupResponse is the response struct for api DeleteApiGroup
type DeleteApiGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteApiGroupRequest creates a request to invoke DeleteApiGroup API
func CreateDeleteApiGroupRequest() (request *DeleteApiGroupRequest) {
	request = &DeleteApiGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteApiGroup", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteApiGroupResponse creates a response to parse from DeleteApiGroup response
func CreateDeleteApiGroupResponse() (response *DeleteApiGroupResponse) {
	response = &DeleteApiGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
