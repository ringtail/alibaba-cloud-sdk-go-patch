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

// DescribeAdminInfo invokes the lto.DescribeAdminInfo API synchronously
func (client *Client) DescribeAdminInfo(request *DescribeAdminInfoRequest) (response *DescribeAdminInfoResponse, err error) {
	response = CreateDescribeAdminInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAdminInfoWithChan invokes the lto.DescribeAdminInfo API asynchronously
func (client *Client) DescribeAdminInfoWithChan(request *DescribeAdminInfoRequest) (<-chan *DescribeAdminInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeAdminInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAdminInfo(request)
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

// DescribeAdminInfoWithCallback invokes the lto.DescribeAdminInfo API asynchronously
func (client *Client) DescribeAdminInfoWithCallback(request *DescribeAdminInfoRequest, callback func(response *DescribeAdminInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAdminInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeAdminInfo(request)
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

// DescribeAdminInfoRequest is the request struct for api DescribeAdminInfo
type DescribeAdminInfoRequest struct {
	*requests.RpcRequest
}

// DescribeAdminInfoResponse is the response struct for api DescribeAdminInfo
type DescribeAdminInfoResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateDescribeAdminInfoRequest creates a request to invoke DescribeAdminInfo API
func CreateDescribeAdminInfoRequest() (request *DescribeAdminInfoRequest) {
	request = &DescribeAdminInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("lto", "2021-07-07", "DescribeAdminInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeAdminInfoResponse creates a response to parse from DescribeAdminInfo response
func CreateDescribeAdminInfoResponse() (response *DescribeAdminInfoResponse) {
	response = &DescribeAdminInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
