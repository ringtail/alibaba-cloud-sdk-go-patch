package dataworks_public

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

// CreateTableLevel invokes the dataworks_public.CreateTableLevel API synchronously
func (client *Client) CreateTableLevel(request *CreateTableLevelRequest) (response *CreateTableLevelResponse, err error) {
	response = CreateCreateTableLevelResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTableLevelWithChan invokes the dataworks_public.CreateTableLevel API asynchronously
func (client *Client) CreateTableLevelWithChan(request *CreateTableLevelRequest) (<-chan *CreateTableLevelResponse, <-chan error) {
	responseChan := make(chan *CreateTableLevelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTableLevel(request)
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

// CreateTableLevelWithCallback invokes the dataworks_public.CreateTableLevel API asynchronously
func (client *Client) CreateTableLevelWithCallback(request *CreateTableLevelRequest, callback func(response *CreateTableLevelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTableLevelResponse
		var err error
		defer close(result)
		response, err = client.CreateTableLevel(request)
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

// CreateTableLevelRequest is the request struct for api CreateTableLevel
type CreateTableLevelRequest struct {
	*requests.RpcRequest
	LevelType   requests.Integer `position:"Query" name:"LevelType"`
	Description string           `position:"Query" name:"Description"`
	Name        string           `position:"Query" name:"Name"`
	ProjectId   requests.Integer `position:"Query" name:"ProjectId"`
}

// CreateTableLevelResponse is the response struct for api CreateTableLevel
type CreateTableLevelResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool   `json:"Success" xml:"Success"`
	LevelId        int64  `json:"LevelId" xml:"LevelId"`
}

// CreateCreateTableLevelRequest creates a request to invoke CreateTableLevel API
func CreateCreateTableLevelRequest() (request *CreateTableLevelRequest) {
	request = &CreateTableLevelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "CreateTableLevel", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateTableLevelResponse creates a response to parse from CreateTableLevel response
func CreateCreateTableLevelResponse() (response *CreateTableLevelResponse) {
	response = &CreateTableLevelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
