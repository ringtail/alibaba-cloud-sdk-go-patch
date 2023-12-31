package companyreg

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

// QueryAvailableNumbers invokes the companyreg.QueryAvailableNumbers API synchronously
func (client *Client) QueryAvailableNumbers(request *QueryAvailableNumbersRequest) (response *QueryAvailableNumbersResponse, err error) {
	response = CreateQueryAvailableNumbersResponse()
	err = client.DoAction(request, response)
	return
}

// QueryAvailableNumbersWithChan invokes the companyreg.QueryAvailableNumbers API asynchronously
func (client *Client) QueryAvailableNumbersWithChan(request *QueryAvailableNumbersRequest) (<-chan *QueryAvailableNumbersResponse, <-chan error) {
	responseChan := make(chan *QueryAvailableNumbersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryAvailableNumbers(request)
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

// QueryAvailableNumbersWithCallback invokes the companyreg.QueryAvailableNumbers API asynchronously
func (client *Client) QueryAvailableNumbersWithCallback(request *QueryAvailableNumbersRequest, callback func(response *QueryAvailableNumbersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryAvailableNumbersResponse
		var err error
		defer close(result)
		response, err = client.QueryAvailableNumbers(request)
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

// QueryAvailableNumbersRequest is the request struct for api QueryAvailableNumbers
type QueryAvailableNumbersRequest struct {
	*requests.RpcRequest
	BizType string `position:"Query" name:"BizType"`
}

// QueryAvailableNumbersResponse is the response struct for api QueryAvailableNumbers
type QueryAvailableNumbersResponse struct {
	*responses.BaseResponse
	ErrorMsg  string   `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	ErrorCode string   `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool     `json:"Success" xml:"Success"`
	Data      []string `json:"Data" xml:"Data"`
}

// CreateQueryAvailableNumbersRequest creates a request to invoke QueryAvailableNumbers API
func CreateQueryAvailableNumbersRequest() (request *QueryAvailableNumbersRequest) {
	request = &QueryAvailableNumbersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-03-06", "QueryAvailableNumbers", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryAvailableNumbersResponse creates a response to parse from QueryAvailableNumbers response
func CreateQueryAvailableNumbersResponse() (response *QueryAvailableNumbersResponse) {
	response = &QueryAvailableNumbersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
