package viapi_regen

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

// DeleteTrainTask invokes the viapi_regen.DeleteTrainTask API synchronously
func (client *Client) DeleteTrainTask(request *DeleteTrainTaskRequest) (response *DeleteTrainTaskResponse, err error) {
	response = CreateDeleteTrainTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTrainTaskWithChan invokes the viapi_regen.DeleteTrainTask API asynchronously
func (client *Client) DeleteTrainTaskWithChan(request *DeleteTrainTaskRequest) (<-chan *DeleteTrainTaskResponse, <-chan error) {
	responseChan := make(chan *DeleteTrainTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTrainTask(request)
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

// DeleteTrainTaskWithCallback invokes the viapi_regen.DeleteTrainTask API asynchronously
func (client *Client) DeleteTrainTaskWithCallback(request *DeleteTrainTaskRequest, callback func(response *DeleteTrainTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTrainTaskResponse
		var err error
		defer close(result)
		response, err = client.DeleteTrainTask(request)
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

// DeleteTrainTaskRequest is the request struct for api DeleteTrainTask
type DeleteTrainTaskRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Body" name:"Id"`
}

// DeleteTrainTaskResponse is the response struct for api DeleteTrainTask
type DeleteTrainTaskResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDeleteTrainTaskRequest creates a request to invoke DeleteTrainTask API
func CreateDeleteTrainTaskRequest() (request *DeleteTrainTaskRequest) {
	request = &DeleteTrainTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi-regen", "2021-11-19", "DeleteTrainTask", "selflearning", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteTrainTaskResponse creates a response to parse from DeleteTrainTask response
func CreateDeleteTrainTaskResponse() (response *DeleteTrainTaskResponse) {
	response = &DeleteTrainTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
