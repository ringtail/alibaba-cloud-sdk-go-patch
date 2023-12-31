package avatar

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

// SendText invokes the avatar.SendText API synchronously
func (client *Client) SendText(request *SendTextRequest) (response *SendTextResponse, err error) {
	response = CreateSendTextResponse()
	err = client.DoAction(request, response)
	return
}

// SendTextWithChan invokes the avatar.SendText API asynchronously
func (client *Client) SendTextWithChan(request *SendTextRequest) (<-chan *SendTextResponse, <-chan error) {
	responseChan := make(chan *SendTextResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendText(request)
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

// SendTextWithCallback invokes the avatar.SendText API asynchronously
func (client *Client) SendTextWithCallback(request *SendTextRequest, callback func(response *SendTextResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendTextResponse
		var err error
		defer close(result)
		response, err = client.SendText(request)
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

// SendTextRequest is the request struct for api SendText
type SendTextRequest struct {
	*requests.RpcRequest
	Feedback        requests.Boolean        `position:"Query" name:"Feedback"`
	UniqueCode      string                  `position:"Query" name:"UniqueCode"`
	StreamExtension SendTextStreamExtension `position:"Query" name:"StreamExtension"  type:"Struct"`
	TenantId        requests.Integer        `position:"Query" name:"TenantId"`
	Interrupt       requests.Boolean        `position:"Query" name:"Interrupt"`
	SessionId       string                  `position:"Query" name:"SessionId"`
	Text            string                  `position:"Query" name:"Text"`
}

// SendTextStreamExtension is a repeated param struct in SendTextRequest
type SendTextStreamExtension struct {
	IsStream string `name:"IsStream"`
	Index    string `name:"Index"`
	Position string `name:"Position"`
}

// SendTextResponse is the response struct for api SendText
type SendTextResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSendTextRequest creates a request to invoke SendText API
func CreateSendTextRequest() (request *SendTextRequest) {
	request = &SendTextRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("avatar", "2022-01-30", "SendText", "", "")
	request.Method = requests.POST
	return
}

// CreateSendTextResponse creates a response to parse from SendText response
func CreateSendTextResponse() (response *SendTextResponse) {
	response = &SendTextResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
