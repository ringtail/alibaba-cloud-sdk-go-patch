package iot

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

// CreateRuleAction invokes the iot.CreateRuleAction API synchronously
func (client *Client) CreateRuleAction(request *CreateRuleActionRequest) (response *CreateRuleActionResponse, err error) {
	response = CreateCreateRuleActionResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRuleActionWithChan invokes the iot.CreateRuleAction API asynchronously
func (client *Client) CreateRuleActionWithChan(request *CreateRuleActionRequest) (<-chan *CreateRuleActionResponse, <-chan error) {
	responseChan := make(chan *CreateRuleActionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRuleAction(request)
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

// CreateRuleActionWithCallback invokes the iot.CreateRuleAction API asynchronously
func (client *Client) CreateRuleActionWithCallback(request *CreateRuleActionRequest, callback func(response *CreateRuleActionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRuleActionResponse
		var err error
		defer close(result)
		response, err = client.CreateRuleAction(request)
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

// CreateRuleActionRequest is the request struct for api CreateRuleAction
type CreateRuleActionRequest struct {
	*requests.RpcRequest
	Configuration   string           `position:"Query" name:"Configuration"`
	Type            string           `position:"Query" name:"Type"`
	IotInstanceId   string           `position:"Query" name:"IotInstanceId"`
	ErrorActionFlag requests.Boolean `position:"Query" name:"ErrorActionFlag"`
	ApiProduct      string           `position:"Body" name:"ApiProduct"`
	ApiRevision     string           `position:"Body" name:"ApiRevision"`
	RuleId          requests.Integer `position:"Query" name:"RuleId"`
}

// CreateRuleActionResponse is the response struct for api CreateRuleAction
type CreateRuleActionResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Code         string `json:"Code" xml:"Code"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	ActionId     int64  `json:"ActionId" xml:"ActionId"`
}

// CreateCreateRuleActionRequest creates a request to invoke CreateRuleAction API
func CreateCreateRuleActionRequest() (request *CreateRuleActionRequest) {
	request = &CreateRuleActionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateRuleAction", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateRuleActionResponse creates a response to parse from CreateRuleAction response
func CreateCreateRuleActionResponse() (response *CreateRuleActionResponse) {
	response = &CreateRuleActionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
