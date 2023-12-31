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

// DescribeComponents invokes the sae.DescribeComponents API synchronously
func (client *Client) DescribeComponents(request *DescribeComponentsRequest) (response *DescribeComponentsResponse, err error) {
	response = CreateDescribeComponentsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeComponentsWithChan invokes the sae.DescribeComponents API asynchronously
func (client *Client) DescribeComponentsWithChan(request *DescribeComponentsRequest) (<-chan *DescribeComponentsResponse, <-chan error) {
	responseChan := make(chan *DescribeComponentsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeComponents(request)
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

// DescribeComponentsWithCallback invokes the sae.DescribeComponents API asynchronously
func (client *Client) DescribeComponentsWithCallback(request *DescribeComponentsRequest, callback func(response *DescribeComponentsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeComponentsResponse
		var err error
		defer close(result)
		response, err = client.DescribeComponents(request)
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

// DescribeComponentsRequest is the request struct for api DescribeComponents
type DescribeComponentsRequest struct {
	*requests.RoaRequest
	AppId string `position:"Query" name:"AppId"`
	Type  string `position:"Query" name:"Type"`
}

// DescribeComponentsResponse is the response struct for api DescribeComponents
type DescribeComponentsResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Message   string     `json:"Message" xml:"Message"`
	TraceId   string     `json:"TraceId" xml:"TraceId"`
	ErrorCode string     `json:"ErrorCode" xml:"ErrorCode"`
	Code      string     `json:"Code" xml:"Code"`
	Success   bool       `json:"Success" xml:"Success"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateDescribeComponentsRequest creates a request to invoke DescribeComponents API
func CreateDescribeComponentsRequest() (request *DescribeComponentsRequest) {
	request = &DescribeComponentsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DescribeComponents", "/pop/v1/sam/resource/components", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeComponentsResponse creates a response to parse from DescribeComponents response
func CreateDescribeComponentsResponse() (response *DescribeComponentsResponse) {
	response = &DescribeComponentsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
