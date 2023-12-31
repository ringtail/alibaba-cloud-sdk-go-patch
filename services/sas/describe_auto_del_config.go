package sas

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

// DescribeAutoDelConfig invokes the sas.DescribeAutoDelConfig API synchronously
func (client *Client) DescribeAutoDelConfig(request *DescribeAutoDelConfigRequest) (response *DescribeAutoDelConfigResponse, err error) {
	response = CreateDescribeAutoDelConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAutoDelConfigWithChan invokes the sas.DescribeAutoDelConfig API asynchronously
func (client *Client) DescribeAutoDelConfigWithChan(request *DescribeAutoDelConfigRequest) (<-chan *DescribeAutoDelConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeAutoDelConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAutoDelConfig(request)
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

// DescribeAutoDelConfigWithCallback invokes the sas.DescribeAutoDelConfig API asynchronously
func (client *Client) DescribeAutoDelConfigWithCallback(request *DescribeAutoDelConfigRequest, callback func(response *DescribeAutoDelConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAutoDelConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeAutoDelConfig(request)
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

// DescribeAutoDelConfigRequest is the request struct for api DescribeAutoDelConfig
type DescribeAutoDelConfigRequest struct {
	*requests.RpcRequest
	SourceIp                   string `position:"Query" name:"SourceIp"`
	ResourceDirectoryAccountId string `position:"Query" name:"ResourceDirectoryAccountId"`
}

// DescribeAutoDelConfigResponse is the response struct for api DescribeAutoDelConfig
type DescribeAutoDelConfigResponse struct {
	*responses.BaseResponse
	Days      int    `json:"Days" xml:"Days"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeAutoDelConfigRequest creates a request to invoke DescribeAutoDelConfig API
func CreateDescribeAutoDelConfigRequest() (request *DescribeAutoDelConfigRequest) {
	request = &DescribeAutoDelConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeAutoDelConfig", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAutoDelConfigResponse creates a response to parse from DescribeAutoDelConfig response
func CreateDescribeAutoDelConfigResponse() (response *DescribeAutoDelConfigResponse) {
	response = &DescribeAutoDelConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
