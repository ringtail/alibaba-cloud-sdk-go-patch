package ccc

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

// ListPhoneTags invokes the ccc.ListPhoneTags API synchronously
func (client *Client) ListPhoneTags(request *ListPhoneTagsRequest) (response *ListPhoneTagsResponse, err error) {
	response = CreateListPhoneTagsResponse()
	err = client.DoAction(request, response)
	return
}

// ListPhoneTagsWithChan invokes the ccc.ListPhoneTags API asynchronously
func (client *Client) ListPhoneTagsWithChan(request *ListPhoneTagsRequest) (<-chan *ListPhoneTagsResponse, <-chan error) {
	responseChan := make(chan *ListPhoneTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPhoneTags(request)
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

// ListPhoneTagsWithCallback invokes the ccc.ListPhoneTags API asynchronously
func (client *Client) ListPhoneTagsWithCallback(request *ListPhoneTagsRequest, callback func(response *ListPhoneTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPhoneTagsResponse
		var err error
		defer close(result)
		response, err = client.ListPhoneTags(request)
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

// ListPhoneTagsRequest is the request struct for api ListPhoneTags
type ListPhoneTagsRequest struct {
	*requests.RpcRequest
	NumberGroupIds *[]string        `position:"Query" name:"NumberGroupIds"  type:"Repeated"`
	CurrentPage    requests.Integer `position:"Query" name:"CurrentPage"`
	OutboundOnly   requests.Boolean `position:"Query" name:"OutboundOnly"`
	Number         string           `position:"Query" name:"Number"`
	InstanceId     string           `position:"Query" name:"InstanceId"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
}

// ListPhoneTagsResponse is the response struct for api ListPhoneTags
type ListPhoneTagsResponse struct {
	*responses.BaseResponse
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	Success        bool         `json:"Success" xml:"Success"`
	Code           string       `json:"Code" xml:"Code"`
	Message        string       `json:"Message" xml:"Message"`
	HttpStatusCode int          `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PhoneNumbers   PhoneNumbers `json:"PhoneNumbers" xml:"PhoneNumbers"`
}

// CreateListPhoneTagsRequest creates a request to invoke ListPhoneTags API
func CreateListPhoneTagsRequest() (request *ListPhoneTagsRequest) {
	request = &ListPhoneTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListPhoneTags", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListPhoneTagsResponse creates a response to parse from ListPhoneTags response
func CreateListPhoneTagsResponse() (response *ListPhoneTagsResponse) {
	response = &ListPhoneTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
