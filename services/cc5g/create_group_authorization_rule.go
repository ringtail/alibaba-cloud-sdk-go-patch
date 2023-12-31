package cc5g

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

// CreateGroupAuthorizationRule invokes the cc5g.CreateGroupAuthorizationRule API synchronously
func (client *Client) CreateGroupAuthorizationRule(request *CreateGroupAuthorizationRuleRequest) (response *CreateGroupAuthorizationRuleResponse, err error) {
	response = CreateCreateGroupAuthorizationRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateGroupAuthorizationRuleWithChan invokes the cc5g.CreateGroupAuthorizationRule API asynchronously
func (client *Client) CreateGroupAuthorizationRuleWithChan(request *CreateGroupAuthorizationRuleRequest) (<-chan *CreateGroupAuthorizationRuleResponse, <-chan error) {
	responseChan := make(chan *CreateGroupAuthorizationRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateGroupAuthorizationRule(request)
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

// CreateGroupAuthorizationRuleWithCallback invokes the cc5g.CreateGroupAuthorizationRule API asynchronously
func (client *Client) CreateGroupAuthorizationRuleWithCallback(request *CreateGroupAuthorizationRuleRequest, callback func(response *CreateGroupAuthorizationRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateGroupAuthorizationRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateGroupAuthorizationRule(request)
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

// CreateGroupAuthorizationRuleRequest is the request struct for api CreateGroupAuthorizationRule
type CreateGroupAuthorizationRuleRequest struct {
	*requests.RpcRequest
	WirelessCloudConnectorGroupId string           `position:"Query" name:"WirelessCloudConnectorGroupId"`
	ClientToken                   string           `position:"Query" name:"ClientToken"`
	SourceCidr                    string           `position:"Query" name:"SourceCidr"`
	DestinationType               string           `position:"Query" name:"DestinationType"`
	Destination                   string           `position:"Query" name:"Destination"`
	Description                   string           `position:"Query" name:"Description"`
	Protocol                      string           `position:"Query" name:"Protocol"`
	Policy                        string           `position:"Query" name:"Policy"`
	DryRun                        requests.Boolean `position:"Query" name:"DryRun"`
	DestinationPort               string           `position:"Query" name:"DestinationPort"`
	Name                          string           `position:"Query" name:"Name"`
}

// CreateGroupAuthorizationRuleResponse is the response struct for api CreateGroupAuthorizationRule
type CreateGroupAuthorizationRuleResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	AuthorizationRuleId string `json:"AuthorizationRuleId" xml:"AuthorizationRuleId"`
}

// CreateCreateGroupAuthorizationRuleRequest creates a request to invoke CreateGroupAuthorizationRule API
func CreateCreateGroupAuthorizationRuleRequest() (request *CreateGroupAuthorizationRuleRequest) {
	request = &CreateGroupAuthorizationRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "CreateGroupAuthorizationRule", "fivegcc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateGroupAuthorizationRuleResponse creates a response to parse from CreateGroupAuthorizationRule response
func CreateCreateGroupAuthorizationRuleResponse() (response *CreateGroupAuthorizationRuleResponse) {
	response = &CreateGroupAuthorizationRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
