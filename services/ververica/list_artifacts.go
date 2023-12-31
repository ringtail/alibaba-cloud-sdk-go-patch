package ververica

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

// ListArtifacts invokes the ververica.ListArtifacts API synchronously
func (client *Client) ListArtifacts(request *ListArtifactsRequest) (response *ListArtifactsResponse, err error) {
	response = CreateListArtifactsResponse()
	err = client.DoAction(request, response)
	return
}

// ListArtifactsWithChan invokes the ververica.ListArtifacts API asynchronously
func (client *Client) ListArtifactsWithChan(request *ListArtifactsRequest) (<-chan *ListArtifactsResponse, <-chan error) {
	responseChan := make(chan *ListArtifactsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListArtifacts(request)
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

// ListArtifactsWithCallback invokes the ververica.ListArtifacts API asynchronously
func (client *Client) ListArtifactsWithCallback(request *ListArtifactsRequest, callback func(response *ListArtifactsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListArtifactsResponse
		var err error
		defer close(result)
		response, err = client.ListArtifacts(request)
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

// ListArtifactsRequest is the request struct for api ListArtifacts
type ListArtifactsRequest struct {
	*requests.RoaRequest
	Workspace string `position:"Path" name:"workspace"`
	Namespace string `position:"Path" name:"namespace"`
}

// ListArtifactsResponse is the response struct for api ListArtifacts
type ListArtifactsResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	RequestId string `json:"requestId" xml:"requestId"`
	Data      string `json:"data" xml:"data"`
}

// CreateListArtifactsRequest creates a request to invoke ListArtifacts API
func CreateListArtifactsRequest() (request *ListArtifactsRequest) {
	request = &ListArtifactsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "ListArtifacts", "/pop/workspaces/[workspace]/artifacts/v1/namespaces/[namespace]/artifacts:list", "", "")
	request.Method = requests.GET
	return
}

// CreateListArtifactsResponse creates a response to parse from ListArtifacts response
func CreateListArtifactsResponse() (response *ListArtifactsResponse) {
	response = &ListArtifactsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
