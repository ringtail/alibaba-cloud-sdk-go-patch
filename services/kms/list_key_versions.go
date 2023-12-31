package kms

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

// ListKeyVersions invokes the kms.ListKeyVersions API synchronously
func (client *Client) ListKeyVersions(request *ListKeyVersionsRequest) (response *ListKeyVersionsResponse, err error) {
	response = CreateListKeyVersionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListKeyVersionsWithChan invokes the kms.ListKeyVersions API asynchronously
func (client *Client) ListKeyVersionsWithChan(request *ListKeyVersionsRequest) (<-chan *ListKeyVersionsResponse, <-chan error) {
	responseChan := make(chan *ListKeyVersionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListKeyVersions(request)
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

// ListKeyVersionsWithCallback invokes the kms.ListKeyVersions API asynchronously
func (client *Client) ListKeyVersionsWithCallback(request *ListKeyVersionsRequest, callback func(response *ListKeyVersionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListKeyVersionsResponse
		var err error
		defer close(result)
		response, err = client.ListKeyVersions(request)
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

// ListKeyVersionsRequest is the request struct for api ListKeyVersions
type ListKeyVersionsRequest struct {
	*requests.RpcRequest
	KeyId      string           `position:"Query" name:"KeyId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListKeyVersionsResponse is the response struct for api ListKeyVersions
type ListKeyVersionsResponse struct {
	*responses.BaseResponse
	PageSize    int         `json:"PageSize" xml:"PageSize"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	PageNumber  int         `json:"PageNumber" xml:"PageNumber"`
	TotalCount  int         `json:"TotalCount" xml:"TotalCount"`
	KeyVersions KeyVersions `json:"KeyVersions" xml:"KeyVersions"`
}

// CreateListKeyVersionsRequest creates a request to invoke ListKeyVersions API
func CreateListKeyVersionsRequest() (request *ListKeyVersionsRequest) {
	request = &ListKeyVersionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "ListKeyVersions", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListKeyVersionsResponse creates a response to parse from ListKeyVersions response
func CreateListKeyVersionsResponse() (response *ListKeyVersionsResponse) {
	response = &ListKeyVersionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
