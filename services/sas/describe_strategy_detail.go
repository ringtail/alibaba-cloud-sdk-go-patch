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

// DescribeStrategyDetail invokes the sas.DescribeStrategyDetail API synchronously
func (client *Client) DescribeStrategyDetail(request *DescribeStrategyDetailRequest) (response *DescribeStrategyDetailResponse, err error) {
	response = CreateDescribeStrategyDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStrategyDetailWithChan invokes the sas.DescribeStrategyDetail API asynchronously
func (client *Client) DescribeStrategyDetailWithChan(request *DescribeStrategyDetailRequest) (<-chan *DescribeStrategyDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeStrategyDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStrategyDetail(request)
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

// DescribeStrategyDetailWithCallback invokes the sas.DescribeStrategyDetail API asynchronously
func (client *Client) DescribeStrategyDetailWithCallback(request *DescribeStrategyDetailRequest, callback func(response *DescribeStrategyDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStrategyDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeStrategyDetail(request)
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

// DescribeStrategyDetailRequest is the request struct for api DescribeStrategyDetail
type DescribeStrategyDetailRequest struct {
	*requests.RpcRequest
	SourceIp                   string `position:"Query" name:"SourceIp"`
	Id                         string `position:"Query" name:"Id"`
	Lang                       string `position:"Query" name:"Lang"`
	ResourceDirectoryAccountId string `position:"Query" name:"ResourceDirectoryAccountId"`
}

// DescribeStrategyDetailResponse is the response struct for api DescribeStrategyDetail
type DescribeStrategyDetailResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Strategy  Strategy `json:"Strategy" xml:"Strategy"`
}

// CreateDescribeStrategyDetailRequest creates a request to invoke DescribeStrategyDetail API
func CreateDescribeStrategyDetailRequest() (request *DescribeStrategyDetailRequest) {
	request = &DescribeStrategyDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeStrategyDetail", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeStrategyDetailResponse creates a response to parse from DescribeStrategyDetail response
func CreateDescribeStrategyDetailResponse() (response *DescribeStrategyDetailResponse) {
	response = &DescribeStrategyDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
