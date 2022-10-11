package aiworkspace

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

// ListPermissions invokes the aiworkspace.ListPermissions API synchronously
func (client *Client) ListPermissions(request *ListPermissionsRequest) (response *ListPermissionsResponse, err error) {
	response = CreateListPermissionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListPermissionsWithChan invokes the aiworkspace.ListPermissions API asynchronously
func (client *Client) ListPermissionsWithChan(request *ListPermissionsRequest) (<-chan *ListPermissionsResponse, <-chan error) {
	responseChan := make(chan *ListPermissionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPermissions(request)
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

// ListPermissionsWithCallback invokes the aiworkspace.ListPermissions API asynchronously
func (client *Client) ListPermissionsWithCallback(request *ListPermissionsRequest, callback func(response *ListPermissionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPermissionsResponse
		var err error
		defer close(result)
		response, err = client.ListPermissions(request)
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

// ListPermissionsRequest is the request struct for api ListPermissions
type ListPermissionsRequest struct {
	*requests.RoaRequest
	WorkspaceId string `position:"Path" name:"WorkspaceId"`
}

// ListPermissionsResponse is the response struct for api ListPermissions
type ListPermissionsResponse struct {
	*responses.BaseResponse
	RequestId   string            `json:"RequestId" xml:"RequestId"`
	TotalCount  int64             `json:"TotalCount" xml:"TotalCount"`
	Permissions []PermissionsItem `json:"Permissions" xml:"Permissions"`
}

// CreateListPermissionsRequest creates a request to invoke ListPermissions API
func CreateListPermissionsRequest() (request *ListPermissionsRequest) {
	request = &ListPermissionsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("AIWorkSpace", "2021-02-04", "ListPermissions", "/api/v1/workspaces/[WorkspaceId]/permissions", "", "")
	request.Method = requests.GET
	return
}

// CreateListPermissionsResponse creates a response to parse from ListPermissions response
func CreateListPermissionsResponse() (response *ListPermissionsResponse) {
	response = &ListPermissionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
