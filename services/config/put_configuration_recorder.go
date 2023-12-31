package config

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

// PutConfigurationRecorder invokes the config.PutConfigurationRecorder API synchronously
func (client *Client) PutConfigurationRecorder(request *PutConfigurationRecorderRequest) (response *PutConfigurationRecorderResponse, err error) {
	response = CreatePutConfigurationRecorderResponse()
	err = client.DoAction(request, response)
	return
}

// PutConfigurationRecorderWithChan invokes the config.PutConfigurationRecorder API asynchronously
func (client *Client) PutConfigurationRecorderWithChan(request *PutConfigurationRecorderRequest) (<-chan *PutConfigurationRecorderResponse, <-chan error) {
	responseChan := make(chan *PutConfigurationRecorderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutConfigurationRecorder(request)
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

// PutConfigurationRecorderWithCallback invokes the config.PutConfigurationRecorder API asynchronously
func (client *Client) PutConfigurationRecorderWithCallback(request *PutConfigurationRecorderRequest, callback func(response *PutConfigurationRecorderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutConfigurationRecorderResponse
		var err error
		defer close(result)
		response, err = client.PutConfigurationRecorder(request)
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

// PutConfigurationRecorderRequest is the request struct for api PutConfigurationRecorder
type PutConfigurationRecorderRequest struct {
	*requests.RpcRequest
	ResourceTypes string `position:"Body" name:"ResourceTypes"`
}

// PutConfigurationRecorderResponse is the response struct for api PutConfigurationRecorder
type PutConfigurationRecorderResponse struct {
	*responses.BaseResponse
	RequestId             string                `json:"RequestId" xml:"RequestId"`
	ConfigurationRecorder ConfigurationRecorder `json:"ConfigurationRecorder" xml:"ConfigurationRecorder"`
}

// CreatePutConfigurationRecorderRequest creates a request to invoke PutConfigurationRecorder API
func CreatePutConfigurationRecorderRequest() (request *PutConfigurationRecorderRequest) {
	request = &PutConfigurationRecorderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2019-01-08", "PutConfigurationRecorder", "", "")
	request.Method = requests.POST
	return
}

// CreatePutConfigurationRecorderResponse creates a response to parse from PutConfigurationRecorder response
func CreatePutConfigurationRecorderResponse() (response *PutConfigurationRecorderResponse) {
	response = &PutConfigurationRecorderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
