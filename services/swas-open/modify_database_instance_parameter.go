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

// ModifyDatabaseInstanceParameter invokes the swas_open.ModifyDatabaseInstanceParameter API synchronously
func (client *Client) ModifyDatabaseInstanceParameter(request *ModifyDatabaseInstanceParameterRequest) (response *ModifyDatabaseInstanceParameterResponse, err error) {
	response = CreateModifyDatabaseInstanceParameterResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDatabaseInstanceParameterWithChan invokes the swas_open.ModifyDatabaseInstanceParameter API asynchronously
func (client *Client) ModifyDatabaseInstanceParameterWithChan(request *ModifyDatabaseInstanceParameterRequest) (<-chan *ModifyDatabaseInstanceParameterResponse, <-chan error) {
	responseChan := make(chan *ModifyDatabaseInstanceParameterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDatabaseInstanceParameter(request)
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

// ModifyDatabaseInstanceParameterWithCallback invokes the swas_open.ModifyDatabaseInstanceParameter API asynchronously
func (client *Client) ModifyDatabaseInstanceParameterWithCallback(request *ModifyDatabaseInstanceParameterRequest, callback func(response *ModifyDatabaseInstanceParameterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDatabaseInstanceParameterResponse
		var err error
		defer close(result)
		response, err = client.ModifyDatabaseInstanceParameter(request)
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

// ModifyDatabaseInstanceParameterRequest is the request struct for api ModifyDatabaseInstanceParameter
type ModifyDatabaseInstanceParameterRequest struct {
	*requests.RpcRequest
	ClientToken        string           `position:"Query" name:"ClientToken"`
	DatabaseInstanceId string           `position:"Query" name:"DatabaseInstanceId"`
	ForceRestart       requests.Boolean `position:"Query" name:"ForceRestart"`
	Parameters         string           `position:"Query" name:"Parameters"`
}

// ModifyDatabaseInstanceParameterResponse is the response struct for api ModifyDatabaseInstanceParameter
type ModifyDatabaseInstanceParameterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDatabaseInstanceParameterRequest creates a request to invoke ModifyDatabaseInstanceParameter API
func CreateModifyDatabaseInstanceParameterRequest() (request *ModifyDatabaseInstanceParameterRequest) {
	request = &ModifyDatabaseInstanceParameterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "ModifyDatabaseInstanceParameter", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDatabaseInstanceParameterResponse creates a response to parse from ModifyDatabaseInstanceParameter response
func CreateModifyDatabaseInstanceParameterResponse() (response *ModifyDatabaseInstanceParameterResponse) {
	response = &ModifyDatabaseInstanceParameterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
