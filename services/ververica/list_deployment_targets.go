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

// ListDeploymentTargets invokes the ververica.ListDeploymentTargets API synchronously
func (client *Client) ListDeploymentTargets(request *ListDeploymentTargetsRequest) (response *ListDeploymentTargetsResponse, err error) {
	response = CreateListDeploymentTargetsResponse()
	err = client.DoAction(request, response)
	return
}

// ListDeploymentTargetsWithChan invokes the ververica.ListDeploymentTargets API asynchronously
func (client *Client) ListDeploymentTargetsWithChan(request *ListDeploymentTargetsRequest) (<-chan *ListDeploymentTargetsResponse, <-chan error) {
	responseChan := make(chan *ListDeploymentTargetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDeploymentTargets(request)
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

// ListDeploymentTargetsWithCallback invokes the ververica.ListDeploymentTargets API asynchronously
func (client *Client) ListDeploymentTargetsWithCallback(request *ListDeploymentTargetsRequest, callback func(response *ListDeploymentTargetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDeploymentTargetsResponse
		var err error
		defer close(result)
		response, err = client.ListDeploymentTargets(request)
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

// ListDeploymentTargetsRequest is the request struct for api ListDeploymentTargets
type ListDeploymentTargetsRequest struct {
	*requests.RoaRequest
	Workspace string `position:"Path" name:"workspace"`
	Namespace string `position:"Path" name:"namespace"`
}

// ListDeploymentTargetsResponse is the response struct for api ListDeploymentTargets
type ListDeploymentTargetsResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	RequestId string `json:"requestId" xml:"requestId"`
	Data      string `json:"data" xml:"data"`
}

// CreateListDeploymentTargetsRequest creates a request to invoke ListDeploymentTargets API
func CreateListDeploymentTargetsRequest() (request *ListDeploymentTargetsRequest) {
	request = &ListDeploymentTargetsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "ListDeploymentTargets", "/pop/workspaces/[workspace]/api/v1/namespaces/[namespace]/deployment-targets", "", "")
	request.Method = requests.GET
	return
}

// CreateListDeploymentTargetsResponse creates a response to parse from ListDeploymentTargets response
func CreateListDeploymentTargetsResponse() (response *ListDeploymentTargetsResponse) {
	response = &ListDeploymentTargetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
