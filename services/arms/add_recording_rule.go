package arms

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

// AddRecordingRule invokes the arms.AddRecordingRule API synchronously
func (client *Client) AddRecordingRule(request *AddRecordingRuleRequest) (response *AddRecordingRuleResponse, err error) {
	response = CreateAddRecordingRuleResponse()
	err = client.DoAction(request, response)
	return
}

// AddRecordingRuleWithChan invokes the arms.AddRecordingRule API asynchronously
func (client *Client) AddRecordingRuleWithChan(request *AddRecordingRuleRequest) (<-chan *AddRecordingRuleResponse, <-chan error) {
	responseChan := make(chan *AddRecordingRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddRecordingRule(request)
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

// AddRecordingRuleWithCallback invokes the arms.AddRecordingRule API asynchronously
func (client *Client) AddRecordingRuleWithCallback(request *AddRecordingRuleRequest, callback func(response *AddRecordingRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddRecordingRuleResponse
		var err error
		defer close(result)
		response, err = client.AddRecordingRule(request)
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

// AddRecordingRuleRequest is the request struct for api AddRecordingRule
type AddRecordingRuleRequest struct {
	*requests.RpcRequest
	RuleYaml  string `position:"Query" name:"RuleYaml"`
	ClusterId string `position:"Query" name:"ClusterId"`
}

// AddRecordingRuleResponse is the response struct for api AddRecordingRule
type AddRecordingRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateAddRecordingRuleRequest creates a request to invoke AddRecordingRule API
func CreateAddRecordingRuleRequest() (request *AddRecordingRuleRequest) {
	request = &AddRecordingRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "AddRecordingRule", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddRecordingRuleResponse creates a response to parse from AddRecordingRule response
func CreateAddRecordingRuleResponse() (response *AddRecordingRuleResponse) {
	response = &AddRecordingRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
