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

// EnableFirewallRule invokes the swas_open.EnableFirewallRule API synchronously
func (client *Client) EnableFirewallRule(request *EnableFirewallRuleRequest) (response *EnableFirewallRuleResponse, err error) {
	response = CreateEnableFirewallRuleResponse()
	err = client.DoAction(request, response)
	return
}

// EnableFirewallRuleWithChan invokes the swas_open.EnableFirewallRule API asynchronously
func (client *Client) EnableFirewallRuleWithChan(request *EnableFirewallRuleRequest) (<-chan *EnableFirewallRuleResponse, <-chan error) {
	responseChan := make(chan *EnableFirewallRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableFirewallRule(request)
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

// EnableFirewallRuleWithCallback invokes the swas_open.EnableFirewallRule API asynchronously
func (client *Client) EnableFirewallRuleWithCallback(request *EnableFirewallRuleRequest, callback func(response *EnableFirewallRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableFirewallRuleResponse
		var err error
		defer close(result)
		response, err = client.EnableFirewallRule(request)
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

// EnableFirewallRuleRequest is the request struct for api EnableFirewallRule
type EnableFirewallRuleRequest struct {
	*requests.RpcRequest
	ClientToken  string `position:"Query" name:"ClientToken"`
	SourceCidrIp string `position:"Query" name:"SourceCidrIp"`
	Remark       string `position:"Query" name:"Remark"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	RuleId       string `position:"Query" name:"RuleId"`
}

// EnableFirewallRuleResponse is the response struct for api EnableFirewallRule
type EnableFirewallRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableFirewallRuleRequest creates a request to invoke EnableFirewallRule API
func CreateEnableFirewallRuleRequest() (request *EnableFirewallRuleRequest) {
	request = &EnableFirewallRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "EnableFirewallRule", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableFirewallRuleResponse creates a response to parse from EnableFirewallRule response
func CreateEnableFirewallRuleResponse() (response *EnableFirewallRuleResponse) {
	response = &EnableFirewallRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
