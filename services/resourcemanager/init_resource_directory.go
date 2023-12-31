package resourcemanager

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

// InitResourceDirectory invokes the resourcemanager.InitResourceDirectory API synchronously
func (client *Client) InitResourceDirectory(request *InitResourceDirectoryRequest) (response *InitResourceDirectoryResponse, err error) {
	response = CreateInitResourceDirectoryResponse()
	err = client.DoAction(request, response)
	return
}

// InitResourceDirectoryWithChan invokes the resourcemanager.InitResourceDirectory API asynchronously
func (client *Client) InitResourceDirectoryWithChan(request *InitResourceDirectoryRequest) (<-chan *InitResourceDirectoryResponse, <-chan error) {
	responseChan := make(chan *InitResourceDirectoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InitResourceDirectory(request)
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

// InitResourceDirectoryWithCallback invokes the resourcemanager.InitResourceDirectory API asynchronously
func (client *Client) InitResourceDirectoryWithCallback(request *InitResourceDirectoryRequest, callback func(response *InitResourceDirectoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InitResourceDirectoryResponse
		var err error
		defer close(result)
		response, err = client.InitResourceDirectory(request)
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

// InitResourceDirectoryRequest is the request struct for api InitResourceDirectory
type InitResourceDirectoryRequest struct {
	*requests.RpcRequest
}

// InitResourceDirectoryResponse is the response struct for api InitResourceDirectory
type InitResourceDirectoryResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	ResourceDirectory ResourceDirectory `json:"ResourceDirectory" xml:"ResourceDirectory"`
}

// CreateInitResourceDirectoryRequest creates a request to invoke InitResourceDirectory API
func CreateInitResourceDirectoryRequest() (request *InitResourceDirectoryRequest) {
	request = &InitResourceDirectoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "InitResourceDirectory", "resourcemanager", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInitResourceDirectoryResponse creates a response to parse from InitResourceDirectory response
func CreateInitResourceDirectoryResponse() (response *InitResourceDirectoryResponse) {
	response = &InitResourceDirectoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
