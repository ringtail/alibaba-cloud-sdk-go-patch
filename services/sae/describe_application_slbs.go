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

// DescribeApplicationSlbs invokes the sae.DescribeApplicationSlbs API synchronously
func (client *Client) DescribeApplicationSlbs(request *DescribeApplicationSlbsRequest) (response *DescribeApplicationSlbsResponse, err error) {
	response = CreateDescribeApplicationSlbsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApplicationSlbsWithChan invokes the sae.DescribeApplicationSlbs API asynchronously
func (client *Client) DescribeApplicationSlbsWithChan(request *DescribeApplicationSlbsRequest) (<-chan *DescribeApplicationSlbsResponse, <-chan error) {
	responseChan := make(chan *DescribeApplicationSlbsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApplicationSlbs(request)
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

// DescribeApplicationSlbsWithCallback invokes the sae.DescribeApplicationSlbs API asynchronously
func (client *Client) DescribeApplicationSlbsWithCallback(request *DescribeApplicationSlbsRequest, callback func(response *DescribeApplicationSlbsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApplicationSlbsResponse
		var err error
		defer close(result)
		response, err = client.DescribeApplicationSlbs(request)
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

// DescribeApplicationSlbsRequest is the request struct for api DescribeApplicationSlbs
type DescribeApplicationSlbsRequest struct {
	*requests.RoaRequest
	AppId string `position:"Query" name:"AppId"`
}

// DescribeApplicationSlbsResponse is the response struct for api DescribeApplicationSlbs
type DescribeApplicationSlbsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeApplicationSlbsRequest creates a request to invoke DescribeApplicationSlbs API
func CreateDescribeApplicationSlbsRequest() (request *DescribeApplicationSlbsRequest) {
	request = &DescribeApplicationSlbsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DescribeApplicationSlbs", "/pop/v1/sam/app/slb", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeApplicationSlbsResponse creates a response to parse from DescribeApplicationSlbs response
func CreateDescribeApplicationSlbsResponse() (response *DescribeApplicationSlbsResponse) {
	response = &DescribeApplicationSlbsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
