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

// ListRecentlyRecycledDirectories invokes the nas.ListRecentlyRecycledDirectories API synchronously
func (client *Client) ListRecentlyRecycledDirectories(request *ListRecentlyRecycledDirectoriesRequest) (response *ListRecentlyRecycledDirectoriesResponse, err error) {
	response = CreateListRecentlyRecycledDirectoriesResponse()
	err = client.DoAction(request, response)
	return
}

// ListRecentlyRecycledDirectoriesWithChan invokes the nas.ListRecentlyRecycledDirectories API asynchronously
func (client *Client) ListRecentlyRecycledDirectoriesWithChan(request *ListRecentlyRecycledDirectoriesRequest) (<-chan *ListRecentlyRecycledDirectoriesResponse, <-chan error) {
	responseChan := make(chan *ListRecentlyRecycledDirectoriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRecentlyRecycledDirectories(request)
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

// ListRecentlyRecycledDirectoriesWithCallback invokes the nas.ListRecentlyRecycledDirectories API asynchronously
func (client *Client) ListRecentlyRecycledDirectoriesWithCallback(request *ListRecentlyRecycledDirectoriesRequest, callback func(response *ListRecentlyRecycledDirectoriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRecentlyRecycledDirectoriesResponse
		var err error
		defer close(result)
		response, err = client.ListRecentlyRecycledDirectories(request)
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

// ListRecentlyRecycledDirectoriesRequest is the request struct for api ListRecentlyRecycledDirectories
type ListRecentlyRecycledDirectoriesRequest struct {
	*requests.RpcRequest
	NextToken    string           `position:"Query" name:"NextToken"`
	FileSystemId string           `position:"Query" name:"FileSystemId"`
	MaxResults   requests.Integer `position:"Query" name:"MaxResults"`
}

// ListRecentlyRecycledDirectoriesResponse is the response struct for api ListRecentlyRecycledDirectories
type ListRecentlyRecycledDirectoriesResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	NextToken string  `json:"NextToken" xml:"NextToken"`
	Entries   []Entry `json:"Entries" xml:"Entries"`
}

// CreateListRecentlyRecycledDirectoriesRequest creates a request to invoke ListRecentlyRecycledDirectories API
func CreateListRecentlyRecycledDirectoriesRequest() (request *ListRecentlyRecycledDirectoriesRequest) {
	request = &ListRecentlyRecycledDirectoriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "ListRecentlyRecycledDirectories", "nas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListRecentlyRecycledDirectoriesResponse creates a response to parse from ListRecentlyRecycledDirectories response
func CreateListRecentlyRecycledDirectoriesResponse() (response *ListRecentlyRecycledDirectoriesResponse) {
	response = &ListRecentlyRecycledDirectoriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
