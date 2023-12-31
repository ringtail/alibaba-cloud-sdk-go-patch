package eventbridge

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

// GetEventBus invokes the eventbridge.GetEventBus API synchronously
func (client *Client) GetEventBus(request *GetEventBusRequest) (response *GetEventBusResponse, err error) {
	response = CreateGetEventBusResponse()
	err = client.DoAction(request, response)
	return
}

// GetEventBusWithChan invokes the eventbridge.GetEventBus API asynchronously
func (client *Client) GetEventBusWithChan(request *GetEventBusRequest) (<-chan *GetEventBusResponse, <-chan error) {
	responseChan := make(chan *GetEventBusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetEventBus(request)
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

// GetEventBusWithCallback invokes the eventbridge.GetEventBus API asynchronously
func (client *Client) GetEventBusWithCallback(request *GetEventBusRequest, callback func(response *GetEventBusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetEventBusResponse
		var err error
		defer close(result)
		response, err = client.GetEventBus(request)
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

// GetEventBusRequest is the request struct for api GetEventBus
type GetEventBusRequest struct {
	*requests.RpcRequest
	EventBusName string `position:"Query" name:"EventBusName"`
}

// GetEventBusResponse is the response struct for api GetEventBus
type GetEventBusResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetEventBusRequest creates a request to invoke GetEventBus API
func CreateGetEventBusRequest() (request *GetEventBusRequest) {
	request = &GetEventBusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eventbridge", "2020-04-01", "GetEventBus", "", "")
	request.Method = requests.POST
	return
}

// CreateGetEventBusResponse creates a response to parse from GetEventBus response
func CreateGetEventBusResponse() (response *GetEventBusResponse) {
	response = &GetEventBusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
