package sae

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

// ListLogConfigs invokes the sae.ListLogConfigs API synchronously
func (client *Client) ListLogConfigs(request *ListLogConfigsRequest) (response *ListLogConfigsResponse, err error) {
	response = CreateListLogConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// ListLogConfigsWithChan invokes the sae.ListLogConfigs API asynchronously
func (client *Client) ListLogConfigsWithChan(request *ListLogConfigsRequest) (<-chan *ListLogConfigsResponse, <-chan error) {
	responseChan := make(chan *ListLogConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListLogConfigs(request)
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

// ListLogConfigsWithCallback invokes the sae.ListLogConfigs API asynchronously
func (client *Client) ListLogConfigsWithCallback(request *ListLogConfigsRequest, callback func(response *ListLogConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListLogConfigsResponse
		var err error
		defer close(result)
		response, err = client.ListLogConfigs(request)
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

// ListLogConfigsRequest is the request struct for api ListLogConfigs
type ListLogConfigsRequest struct {
	*requests.RoaRequest
	AppId       string           `position:"Query" name:"AppId"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
}

// ListLogConfigsResponse is the response struct for api ListLogConfigs
type ListLogConfigsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListLogConfigsRequest creates a request to invoke ListLogConfigs API
func CreateListLogConfigsRequest() (request *ListLogConfigsRequest) {
	request = &ListLogConfigsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "ListLogConfigs", "/pop/v1/sam/log/listLogConfigs", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListLogConfigsResponse creates a response to parse from ListLogConfigs response
func CreateListLogConfigsResponse() (response *ListLogConfigsResponse) {
	response = &ListLogConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
