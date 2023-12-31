package dfs

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

// ListMountPoints invokes the dfs.ListMountPoints API synchronously
func (client *Client) ListMountPoints(request *ListMountPointsRequest) (response *ListMountPointsResponse, err error) {
	response = CreateListMountPointsResponse()
	err = client.DoAction(request, response)
	return
}

// ListMountPointsWithChan invokes the dfs.ListMountPoints API asynchronously
func (client *Client) ListMountPointsWithChan(request *ListMountPointsRequest) (<-chan *ListMountPointsResponse, <-chan error) {
	responseChan := make(chan *ListMountPointsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMountPoints(request)
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

// ListMountPointsWithCallback invokes the dfs.ListMountPoints API asynchronously
func (client *Client) ListMountPointsWithCallback(request *ListMountPointsRequest, callback func(response *ListMountPointsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMountPointsResponse
		var err error
		defer close(result)
		response, err = client.ListMountPoints(request)
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

// ListMountPointsRequest is the request struct for api ListMountPoints
type ListMountPointsRequest struct {
	*requests.RpcRequest
	InputRegionId string           `position:"Query" name:"InputRegionId"`
	Limit         requests.Integer `position:"Query" name:"Limit"`
	FileSystemId  string           `position:"Query" name:"FileSystemId"`
	OrderBy       string           `position:"Query" name:"OrderBy"`
	StartOffset   requests.Integer `position:"Query" name:"StartOffset"`
	OrderType     string           `position:"Query" name:"OrderType"`
}

// ListMountPointsResponse is the response struct for api ListMountPoints
type ListMountPointsResponse struct {
	*responses.BaseResponse
	TotalCount  int          `json:"TotalCount" xml:"TotalCount"`
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	MountPoints []MountPoint `json:"MountPoints" xml:"MountPoints"`
}

// CreateListMountPointsRequest creates a request to invoke ListMountPoints API
func CreateListMountPointsRequest() (request *ListMountPointsRequest) {
	request = &ListMountPointsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DFS", "2018-06-20", "ListMountPoints", "alidfs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListMountPointsResponse creates a response to parse from ListMountPoints response
func CreateListMountPointsResponse() (response *ListMountPointsResponse) {
	response = &ListMountPointsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
