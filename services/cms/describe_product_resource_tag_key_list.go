package cms

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

// DescribeProductResourceTagKeyList invokes the cms.DescribeProductResourceTagKeyList API synchronously
func (client *Client) DescribeProductResourceTagKeyList(request *DescribeProductResourceTagKeyListRequest) (response *DescribeProductResourceTagKeyListResponse, err error) {
	response = CreateDescribeProductResourceTagKeyListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeProductResourceTagKeyListWithChan invokes the cms.DescribeProductResourceTagKeyList API asynchronously
func (client *Client) DescribeProductResourceTagKeyListWithChan(request *DescribeProductResourceTagKeyListRequest) (<-chan *DescribeProductResourceTagKeyListResponse, <-chan error) {
	responseChan := make(chan *DescribeProductResourceTagKeyListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeProductResourceTagKeyList(request)
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

// DescribeProductResourceTagKeyListWithCallback invokes the cms.DescribeProductResourceTagKeyList API asynchronously
func (client *Client) DescribeProductResourceTagKeyListWithCallback(request *DescribeProductResourceTagKeyListRequest, callback func(response *DescribeProductResourceTagKeyListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeProductResourceTagKeyListResponse
		var err error
		defer close(result)
		response, err = client.DescribeProductResourceTagKeyList(request)
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

// DescribeProductResourceTagKeyListRequest is the request struct for api DescribeProductResourceTagKeyList
type DescribeProductResourceTagKeyListRequest struct {
	*requests.RpcRequest
	NextToken string `position:"Query" name:"NextToken"`
}

// DescribeProductResourceTagKeyListResponse is the response struct for api DescribeProductResourceTagKeyList
type DescribeProductResourceTagKeyListResponse struct {
	*responses.BaseResponse
	Code      string                                     `json:"Code" xml:"Code"`
	Message   string                                     `json:"Message" xml:"Message"`
	NextToken string                                     `json:"NextToken" xml:"NextToken"`
	RequestId string                                     `json:"RequestId" xml:"RequestId"`
	Success   bool                                       `json:"Success" xml:"Success"`
	TagKeys   TagKeysInDescribeProductResourceTagKeyList `json:"TagKeys" xml:"TagKeys"`
}

// CreateDescribeProductResourceTagKeyListRequest creates a request to invoke DescribeProductResourceTagKeyList API
func CreateDescribeProductResourceTagKeyListRequest() (request *DescribeProductResourceTagKeyListRequest) {
	request = &DescribeProductResourceTagKeyListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeProductResourceTagKeyList", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeProductResourceTagKeyListResponse creates a response to parse from DescribeProductResourceTagKeyList response
func CreateDescribeProductResourceTagKeyListResponse() (response *DescribeProductResourceTagKeyListResponse) {
	response = &DescribeProductResourceTagKeyListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
