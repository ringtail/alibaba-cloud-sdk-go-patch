package swas_open

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

// RebootInstance invokes the swas_open.RebootInstance API synchronously
func (client *Client) RebootInstance(request *RebootInstanceRequest) (response *RebootInstanceResponse, err error) {
	response = CreateRebootInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// RebootInstanceWithChan invokes the swas_open.RebootInstance API asynchronously
func (client *Client) RebootInstanceWithChan(request *RebootInstanceRequest) (<-chan *RebootInstanceResponse, <-chan error) {
	responseChan := make(chan *RebootInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RebootInstance(request)
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

// RebootInstanceWithCallback invokes the swas_open.RebootInstance API asynchronously
func (client *Client) RebootInstanceWithCallback(request *RebootInstanceRequest, callback func(response *RebootInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RebootInstanceResponse
		var err error
		defer close(result)
		response, err = client.RebootInstance(request)
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

// RebootInstanceRequest is the request struct for api RebootInstance
type RebootInstanceRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Query" name:"ClientToken"`
	InstanceId  string `position:"Query" name:"InstanceId"`
}

// RebootInstanceResponse is the response struct for api RebootInstance
type RebootInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRebootInstanceRequest creates a request to invoke RebootInstance API
func CreateRebootInstanceRequest() (request *RebootInstanceRequest) {
	request = &RebootInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "RebootInstance", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRebootInstanceResponse creates a response to parse from RebootInstance response
func CreateRebootInstanceResponse() (response *RebootInstanceResponse) {
	response = &RebootInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
