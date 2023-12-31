package live

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

// BatchGetOnlineUsers invokes the live.BatchGetOnlineUsers API synchronously
func (client *Client) BatchGetOnlineUsers(request *BatchGetOnlineUsersRequest) (response *BatchGetOnlineUsersResponse, err error) {
	response = CreateBatchGetOnlineUsersResponse()
	err = client.DoAction(request, response)
	return
}

// BatchGetOnlineUsersWithChan invokes the live.BatchGetOnlineUsers API asynchronously
func (client *Client) BatchGetOnlineUsersWithChan(request *BatchGetOnlineUsersRequest) (<-chan *BatchGetOnlineUsersResponse, <-chan error) {
	responseChan := make(chan *BatchGetOnlineUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchGetOnlineUsers(request)
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

// BatchGetOnlineUsersWithCallback invokes the live.BatchGetOnlineUsers API asynchronously
func (client *Client) BatchGetOnlineUsersWithCallback(request *BatchGetOnlineUsersRequest, callback func(response *BatchGetOnlineUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchGetOnlineUsersResponse
		var err error
		defer close(result)
		response, err = client.BatchGetOnlineUsers(request)
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

// BatchGetOnlineUsersRequest is the request struct for api BatchGetOnlineUsers
type BatchGetOnlineUsersRequest struct {
	*requests.RpcRequest
	GroupId string `position:"Body" name:"GroupId"`
	UserIds string `position:"Body" name:"UserIds"`
	AppId   string `position:"Body" name:"AppId"`
}

// BatchGetOnlineUsersResponse is the response struct for api BatchGetOnlineUsers
type BatchGetOnlineUsersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateBatchGetOnlineUsersRequest creates a request to invoke BatchGetOnlineUsers API
func CreateBatchGetOnlineUsersRequest() (request *BatchGetOnlineUsersRequest) {
	request = &BatchGetOnlineUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "BatchGetOnlineUsers", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBatchGetOnlineUsersResponse creates a response to parse from BatchGetOnlineUsers response
func CreateBatchGetOnlineUsersResponse() (response *BatchGetOnlineUsersResponse) {
	response = &BatchGetOnlineUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
