package dcdn

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

// RollbackDcdnStagingConfig invokes the dcdn.RollbackDcdnStagingConfig API synchronously
func (client *Client) RollbackDcdnStagingConfig(request *RollbackDcdnStagingConfigRequest) (response *RollbackDcdnStagingConfigResponse, err error) {
	response = CreateRollbackDcdnStagingConfigResponse()
	err = client.DoAction(request, response)
	return
}

// RollbackDcdnStagingConfigWithChan invokes the dcdn.RollbackDcdnStagingConfig API asynchronously
func (client *Client) RollbackDcdnStagingConfigWithChan(request *RollbackDcdnStagingConfigRequest) (<-chan *RollbackDcdnStagingConfigResponse, <-chan error) {
	responseChan := make(chan *RollbackDcdnStagingConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RollbackDcdnStagingConfig(request)
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

// RollbackDcdnStagingConfigWithCallback invokes the dcdn.RollbackDcdnStagingConfig API asynchronously
func (client *Client) RollbackDcdnStagingConfigWithCallback(request *RollbackDcdnStagingConfigRequest, callback func(response *RollbackDcdnStagingConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RollbackDcdnStagingConfigResponse
		var err error
		defer close(result)
		response, err = client.RollbackDcdnStagingConfig(request)
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

// RollbackDcdnStagingConfigRequest is the request struct for api RollbackDcdnStagingConfig
type RollbackDcdnStagingConfigRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
}

// RollbackDcdnStagingConfigResponse is the response struct for api RollbackDcdnStagingConfig
type RollbackDcdnStagingConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRollbackDcdnStagingConfigRequest creates a request to invoke RollbackDcdnStagingConfig API
func CreateRollbackDcdnStagingConfigRequest() (request *RollbackDcdnStagingConfigRequest) {
	request = &RollbackDcdnStagingConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "RollbackDcdnStagingConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateRollbackDcdnStagingConfigResponse creates a response to parse from RollbackDcdnStagingConfig response
func CreateRollbackDcdnStagingConfigResponse() (response *RollbackDcdnStagingConfigResponse) {
	response = &RollbackDcdnStagingConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
