package actiontrail

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

// StopLogging invokes the actiontrail.StopLogging API synchronously
func (client *Client) StopLogging(request *StopLoggingRequest) (response *StopLoggingResponse, err error) {
	response = CreateStopLoggingResponse()
	err = client.DoAction(request, response)
	return
}

// StopLoggingWithChan invokes the actiontrail.StopLogging API asynchronously
func (client *Client) StopLoggingWithChan(request *StopLoggingRequest) (<-chan *StopLoggingResponse, <-chan error) {
	responseChan := make(chan *StopLoggingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopLogging(request)
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

// StopLoggingWithCallback invokes the actiontrail.StopLogging API asynchronously
func (client *Client) StopLoggingWithCallback(request *StopLoggingRequest, callback func(response *StopLoggingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopLoggingResponse
		var err error
		defer close(result)
		response, err = client.StopLogging(request)
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

// StopLoggingRequest is the request struct for api StopLogging
type StopLoggingRequest struct {
	*requests.RpcRequest
	Name string `position:"Query" name:"Name"`
}

// StopLoggingResponse is the response struct for api StopLogging
type StopLoggingResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopLoggingRequest creates a request to invoke StopLogging API
func CreateStopLoggingRequest() (request *StopLoggingRequest) {
	request = &StopLoggingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Actiontrail", "2020-07-06", "StopLogging", "", "")
	request.Method = requests.GET
	return
}

// CreateStopLoggingResponse creates a response to parse from StopLogging response
func CreateStopLoggingResponse() (response *StopLoggingResponse) {
	response = &StopLoggingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
