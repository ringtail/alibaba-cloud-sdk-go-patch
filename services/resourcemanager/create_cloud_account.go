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

// CreateCloudAccount invokes the resourcemanager.CreateCloudAccount API synchronously
func (client *Client) CreateCloudAccount(request *CreateCloudAccountRequest) (response *CreateCloudAccountResponse, err error) {
	response = CreateCreateCloudAccountResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCloudAccountWithChan invokes the resourcemanager.CreateCloudAccount API asynchronously
func (client *Client) CreateCloudAccountWithChan(request *CreateCloudAccountRequest) (<-chan *CreateCloudAccountResponse, <-chan error) {
	responseChan := make(chan *CreateCloudAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCloudAccount(request)
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

// CreateCloudAccountWithCallback invokes the resourcemanager.CreateCloudAccount API asynchronously
func (client *Client) CreateCloudAccountWithCallback(request *CreateCloudAccountRequest, callback func(response *CreateCloudAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCloudAccountResponse
		var err error
		defer close(result)
		response, err = client.CreateCloudAccount(request)
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

// CreateCloudAccountRequest is the request struct for api CreateCloudAccount
type CreateCloudAccountRequest struct {
	*requests.RpcRequest
	ParentFolderId string `position:"Query" name:"ParentFolderId"`
	DisplayName    string `position:"Query" name:"DisplayName"`
	Email          string `position:"Query" name:"Email"`
	PayerAccountId string `position:"Query" name:"PayerAccountId"`
}

// CreateCloudAccountResponse is the response struct for api CreateCloudAccount
type CreateCloudAccountResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Account   Account `json:"Account" xml:"Account"`
}

// CreateCreateCloudAccountRequest creates a request to invoke CreateCloudAccount API
func CreateCreateCloudAccountRequest() (request *CreateCloudAccountRequest) {
	request = &CreateCloudAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "CreateCloudAccount", "resourcemanager", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateCloudAccountResponse creates a response to parse from CreateCloudAccount response
func CreateCreateCloudAccountResponse() (response *CreateCloudAccountResponse) {
	response = &CreateCloudAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
