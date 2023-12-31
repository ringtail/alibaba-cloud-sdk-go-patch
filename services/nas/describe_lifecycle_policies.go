package nas

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

// DescribeLifecyclePolicies invokes the nas.DescribeLifecyclePolicies API synchronously
func (client *Client) DescribeLifecyclePolicies(request *DescribeLifecyclePoliciesRequest) (response *DescribeLifecyclePoliciesResponse, err error) {
	response = CreateDescribeLifecyclePoliciesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLifecyclePoliciesWithChan invokes the nas.DescribeLifecyclePolicies API asynchronously
func (client *Client) DescribeLifecyclePoliciesWithChan(request *DescribeLifecyclePoliciesRequest) (<-chan *DescribeLifecyclePoliciesResponse, <-chan error) {
	responseChan := make(chan *DescribeLifecyclePoliciesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLifecyclePolicies(request)
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

// DescribeLifecyclePoliciesWithCallback invokes the nas.DescribeLifecyclePolicies API asynchronously
func (client *Client) DescribeLifecyclePoliciesWithCallback(request *DescribeLifecyclePoliciesRequest, callback func(response *DescribeLifecyclePoliciesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLifecyclePoliciesResponse
		var err error
		defer close(result)
		response, err = client.DescribeLifecyclePolicies(request)
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

// DescribeLifecyclePoliciesRequest is the request struct for api DescribeLifecyclePolicies
type DescribeLifecyclePoliciesRequest struct {
	*requests.RpcRequest
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	FileSystemId string           `position:"Query" name:"FileSystemId"`
}

// DescribeLifecyclePoliciesResponse is the response struct for api DescribeLifecyclePolicies
type DescribeLifecyclePoliciesResponse struct {
	*responses.BaseResponse
	TotalCount        int               `json:"TotalCount" xml:"TotalCount"`
	PageSize          int               `json:"PageSize" xml:"PageSize"`
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	PageNumber        int               `json:"PageNumber" xml:"PageNumber"`
	LifecyclePolicies []LifecyclePolicy `json:"LifecyclePolicies" xml:"LifecyclePolicies"`
}

// CreateDescribeLifecyclePoliciesRequest creates a request to invoke DescribeLifecyclePolicies API
func CreateDescribeLifecyclePoliciesRequest() (request *DescribeLifecyclePoliciesRequest) {
	request = &DescribeLifecyclePoliciesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "DescribeLifecyclePolicies", "nas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeLifecyclePoliciesResponse creates a response to parse from DescribeLifecyclePolicies response
func CreateDescribeLifecyclePoliciesResponse() (response *DescribeLifecyclePoliciesResponse) {
	response = &DescribeLifecyclePoliciesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
