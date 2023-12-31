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

// DescribeGroupedTags invokes the sas.DescribeGroupedTags API synchronously
func (client *Client) DescribeGroupedTags(request *DescribeGroupedTagsRequest) (response *DescribeGroupedTagsResponse, err error) {
	response = CreateDescribeGroupedTagsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGroupedTagsWithChan invokes the sas.DescribeGroupedTags API asynchronously
func (client *Client) DescribeGroupedTagsWithChan(request *DescribeGroupedTagsRequest) (<-chan *DescribeGroupedTagsResponse, <-chan error) {
	responseChan := make(chan *DescribeGroupedTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGroupedTags(request)
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

// DescribeGroupedTagsWithCallback invokes the sas.DescribeGroupedTags API asynchronously
func (client *Client) DescribeGroupedTagsWithCallback(request *DescribeGroupedTagsRequest, callback func(response *DescribeGroupedTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGroupedTagsResponse
		var err error
		defer close(result)
		response, err = client.DescribeGroupedTags(request)
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

// DescribeGroupedTagsRequest is the request struct for api DescribeGroupedTags
type DescribeGroupedTagsRequest struct {
	*requests.RpcRequest
	MachineTypes string `position:"Query" name:"MachineTypes"`
	SourceIp     string `position:"Query" name:"SourceIp"`
}

// DescribeGroupedTagsResponse is the response struct for api DescribeGroupedTags
type DescribeGroupedTagsResponse struct {
	*responses.BaseResponse
	HttpStatusCode int            `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Success        bool           `json:"Success" xml:"Success"`
	Count          int            `json:"Count" xml:"Count"`
	GroupedFileds  []GroupedFiled `json:"GroupedFileds" xml:"GroupedFileds"`
}

// CreateDescribeGroupedTagsRequest creates a request to invoke DescribeGroupedTags API
func CreateDescribeGroupedTagsRequest() (request *DescribeGroupedTagsRequest) {
	request = &DescribeGroupedTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeGroupedTags", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGroupedTagsResponse creates a response to parse from DescribeGroupedTags response
func CreateDescribeGroupedTagsResponse() (response *DescribeGroupedTagsResponse) {
	response = &DescribeGroupedTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
