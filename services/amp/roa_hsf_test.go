package amp

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

// RoaHsfTest invokes the amp.RoaHsfTest API synchronously
func (client *Client) RoaHsfTest(request *RoaHsfTestRequest) (response *RoaHsfTestResponse, err error) {
	response = CreateRoaHsfTestResponse()
	err = client.DoAction(request, response)
	return
}

// RoaHsfTestWithChan invokes the amp.RoaHsfTest API asynchronously
func (client *Client) RoaHsfTestWithChan(request *RoaHsfTestRequest) (<-chan *RoaHsfTestResponse, <-chan error) {
	responseChan := make(chan *RoaHsfTestResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RoaHsfTest(request)
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

// RoaHsfTestWithCallback invokes the amp.RoaHsfTest API asynchronously
func (client *Client) RoaHsfTestWithCallback(request *RoaHsfTestRequest, callback func(response *RoaHsfTestResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RoaHsfTestResponse
		var err error
		defer close(result)
		response, err = client.RoaHsfTest(request)
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

// RoaHsfTestRequest is the request struct for api RoaHsfTest
type RoaHsfTestRequest struct {
	*requests.RoaRequest
	Body RoaHsfTestBody `position:"Body" name:"body"  type:"Struct"`
}

// RoaHsfTestBody is a repeated param struct in RoaHsfTestRequest
type RoaHsfTestBody struct {
	Id string `name:"id"`
}

// RoaHsfTestResponse is the response struct for api RoaHsfTest
type RoaHsfTestResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateRoaHsfTestRequest creates a request to invoke RoaHsfTest API
func CreateRoaHsfTestRequest() (request *RoaHsfTestRequest) {
	request = &RoaHsfTestRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("amp", "2020-07-08", "RoaHsfTest", "/roa/hsf/test", "ServiceCode", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRoaHsfTestResponse creates a response to parse from RoaHsfTest response
func CreateRoaHsfTestResponse() (response *RoaHsfTestResponse) {
	response = &RoaHsfTestResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
