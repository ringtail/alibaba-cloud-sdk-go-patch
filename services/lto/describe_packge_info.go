package lto

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

// DescribePackgeInfo invokes the lto.DescribePackgeInfo API synchronously
func (client *Client) DescribePackgeInfo(request *DescribePackgeInfoRequest) (response *DescribePackgeInfoResponse, err error) {
	response = CreateDescribePackgeInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePackgeInfoWithChan invokes the lto.DescribePackgeInfo API asynchronously
func (client *Client) DescribePackgeInfoWithChan(request *DescribePackgeInfoRequest) (<-chan *DescribePackgeInfoResponse, <-chan error) {
	responseChan := make(chan *DescribePackgeInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePackgeInfo(request)
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

// DescribePackgeInfoWithCallback invokes the lto.DescribePackgeInfo API asynchronously
func (client *Client) DescribePackgeInfoWithCallback(request *DescribePackgeInfoRequest, callback func(response *DescribePackgeInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePackgeInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribePackgeInfo(request)
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

// DescribePackgeInfoRequest is the request struct for api DescribePackgeInfo
type DescribePackgeInfoRequest struct {
	*requests.RpcRequest
	AccountId string `position:"Query" name:"AccountId"`
}

// DescribePackgeInfoResponse is the response struct for api DescribePackgeInfo
type DescribePackgeInfoResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateDescribePackgeInfoRequest creates a request to invoke DescribePackgeInfo API
func CreateDescribePackgeInfoRequest() (request *DescribePackgeInfoRequest) {
	request = &DescribePackgeInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("lto", "2021-07-07", "DescribePackgeInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribePackgeInfoResponse creates a response to parse from DescribePackgeInfo response
func CreateDescribePackgeInfoResponse() (response *DescribePackgeInfoResponse) {
	response = &DescribePackgeInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
