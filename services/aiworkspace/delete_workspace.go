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

// DeleteWorkspace invokes the aiworkspace.DeleteWorkspace API synchronously
func (client *Client) DeleteWorkspace(request *DeleteWorkspaceRequest) (response *DeleteWorkspaceResponse, err error) {
	response = CreateDeleteWorkspaceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteWorkspaceWithChan invokes the aiworkspace.DeleteWorkspace API asynchronously
func (client *Client) DeleteWorkspaceWithChan(request *DeleteWorkspaceRequest) (<-chan *DeleteWorkspaceResponse, <-chan error) {
	responseChan := make(chan *DeleteWorkspaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteWorkspace(request)
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

// DeleteWorkspaceWithCallback invokes the aiworkspace.DeleteWorkspace API asynchronously
func (client *Client) DeleteWorkspaceWithCallback(request *DeleteWorkspaceRequest, callback func(response *DeleteWorkspaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteWorkspaceResponse
		var err error
		defer close(result)
		response, err = client.DeleteWorkspace(request)
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

// DeleteWorkspaceRequest is the request struct for api DeleteWorkspace
type DeleteWorkspaceRequest struct {
	*requests.RoaRequest
	WorkspaceId string `position:"Path" name:"WorkspaceId"`
}

// DeleteWorkspaceResponse is the response struct for api DeleteWorkspace
type DeleteWorkspaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteWorkspaceRequest creates a request to invoke DeleteWorkspace API
func CreateDeleteWorkspaceRequest() (request *DeleteWorkspaceRequest) {
	request = &DeleteWorkspaceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("AIWorkSpace", "2021-02-04", "DeleteWorkspace", "/api/v1/workspaces/[WorkspaceId]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteWorkspaceResponse creates a response to parse from DeleteWorkspace response
func CreateDeleteWorkspaceResponse() (response *DeleteWorkspaceResponse) {
	response = &DeleteWorkspaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
