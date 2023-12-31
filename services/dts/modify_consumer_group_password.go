package dts

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

// ModifyConsumerGroupPassword invokes the dts.ModifyConsumerGroupPassword API synchronously
func (client *Client) ModifyConsumerGroupPassword(request *ModifyConsumerGroupPasswordRequest) (response *ModifyConsumerGroupPasswordResponse, err error) {
	response = CreateModifyConsumerGroupPasswordResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyConsumerGroupPasswordWithChan invokes the dts.ModifyConsumerGroupPassword API asynchronously
func (client *Client) ModifyConsumerGroupPasswordWithChan(request *ModifyConsumerGroupPasswordRequest) (<-chan *ModifyConsumerGroupPasswordResponse, <-chan error) {
	responseChan := make(chan *ModifyConsumerGroupPasswordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyConsumerGroupPassword(request)
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

// ModifyConsumerGroupPasswordWithCallback invokes the dts.ModifyConsumerGroupPassword API asynchronously
func (client *Client) ModifyConsumerGroupPasswordWithCallback(request *ModifyConsumerGroupPasswordRequest, callback func(response *ModifyConsumerGroupPasswordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyConsumerGroupPasswordResponse
		var err error
		defer close(result)
		response, err = client.ModifyConsumerGroupPassword(request)
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

// ModifyConsumerGroupPasswordRequest is the request struct for api ModifyConsumerGroupPassword
type ModifyConsumerGroupPasswordRequest struct {
	*requests.RpcRequest
	ConsumerGroupID          string `position:"Query" name:"ConsumerGroupID"`
	ConsumerGroupPassword    string `position:"Query" name:"ConsumerGroupPassword"`
	AccountId                string `position:"Query" name:"AccountId"`
	ConsumerGroupUserName    string `position:"Query" name:"ConsumerGroupUserName"`
	ConsumerGroupName        string `position:"Query" name:"ConsumerGroupName"`
	SubscriptionInstanceId   string `position:"Query" name:"SubscriptionInstanceId"`
	OwnerId                  string `position:"Query" name:"OwnerId"`
	ConsumerGroupNewPassword string `position:"Query" name:"consumerGroupNewPassword"`
}

// ModifyConsumerGroupPasswordResponse is the response struct for api ModifyConsumerGroupPassword
type ModifyConsumerGroupPasswordResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	ErrCode    string `json:"ErrCode" xml:"ErrCode"`
	Success    string `json:"Success" xml:"Success"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
}

// CreateModifyConsumerGroupPasswordRequest creates a request to invoke ModifyConsumerGroupPassword API
func CreateModifyConsumerGroupPasswordRequest() (request *ModifyConsumerGroupPasswordRequest) {
	request = &ModifyConsumerGroupPasswordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "ModifyConsumerGroupPassword", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyConsumerGroupPasswordResponse creates a response to parse from ModifyConsumerGroupPassword response
func CreateModifyConsumerGroupPasswordResponse() (response *ModifyConsumerGroupPasswordResponse) {
	response = &ModifyConsumerGroupPasswordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
