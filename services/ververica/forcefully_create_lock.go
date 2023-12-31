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

// ForcefullyCreateLock invokes the ververica.ForcefullyCreateLock API synchronously
func (client *Client) ForcefullyCreateLock(request *ForcefullyCreateLockRequest) (response *ForcefullyCreateLockResponse, err error) {
	response = CreateForcefullyCreateLockResponse()
	err = client.DoAction(request, response)
	return
}

// ForcefullyCreateLockWithChan invokes the ververica.ForcefullyCreateLock API asynchronously
func (client *Client) ForcefullyCreateLockWithChan(request *ForcefullyCreateLockRequest) (<-chan *ForcefullyCreateLockResponse, <-chan error) {
	responseChan := make(chan *ForcefullyCreateLockResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ForcefullyCreateLock(request)
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

// ForcefullyCreateLockWithCallback invokes the ververica.ForcefullyCreateLock API asynchronously
func (client *Client) ForcefullyCreateLockWithCallback(request *ForcefullyCreateLockRequest, callback func(response *ForcefullyCreateLockResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ForcefullyCreateLockResponse
		var err error
		defer close(result)
		response, err = client.ForcefullyCreateLock(request)
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

// ForcefullyCreateLockRequest is the request struct for api ForcefullyCreateLock
type ForcefullyCreateLockRequest struct {
	*requests.RoaRequest
	Workspace    string `position:"Path" name:"workspace"`
	ResourceId   string `position:"Query" name:"resourceId"`
	Namespace    string `position:"Path" name:"namespace"`
	ResourceType string `position:"Query" name:"resourceType"`
}

// ForcefullyCreateLockResponse is the response struct for api ForcefullyCreateLock
type ForcefullyCreateLockResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	RequestId string `json:"requestId" xml:"requestId"`
	Data      string `json:"data" xml:"data"`
}

// CreateForcefullyCreateLockRequest creates a request to invoke ForcefullyCreateLock API
func CreateForcefullyCreateLockRequest() (request *ForcefullyCreateLockRequest) {
	request = &ForcefullyCreateLockRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "ForcefullyCreateLock", "/pop/workspaces/[workspace]/api/v1/namespaces/[namespace]/locks", "", "")
	request.Method = requests.POST
	return
}

// CreateForcefullyCreateLockResponse creates a response to parse from ForcefullyCreateLock response
func CreateForcefullyCreateLockResponse() (response *ForcefullyCreateLockResponse) {
	response = &ForcefullyCreateLockResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
