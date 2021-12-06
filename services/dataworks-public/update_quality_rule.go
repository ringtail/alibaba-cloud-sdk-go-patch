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

// UpdateQualityRule invokes the dataworks_public.UpdateQualityRule API synchronously
func (client *Client) UpdateQualityRule(request *UpdateQualityRuleRequest) (response *UpdateQualityRuleResponse, err error) {
	response = CreateUpdateQualityRuleResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateQualityRuleWithChan invokes the dataworks_public.UpdateQualityRule API asynchronously
func (client *Client) UpdateQualityRuleWithChan(request *UpdateQualityRuleRequest) (<-chan *UpdateQualityRuleResponse, <-chan error) {
	responseChan := make(chan *UpdateQualityRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateQualityRule(request)
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

// UpdateQualityRuleWithCallback invokes the dataworks_public.UpdateQualityRule API asynchronously
func (client *Client) UpdateQualityRuleWithCallback(request *UpdateQualityRuleRequest, callback func(response *UpdateQualityRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateQualityRuleResponse
		var err error
		defer close(result)
		response, err = client.UpdateQualityRule(request)
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

// UpdateQualityRuleRequest is the request struct for api UpdateQualityRule
type UpdateQualityRuleRequest struct {
	*requests.RpcRequest
	Trend             string           `position:"Body" name:"Trend"`
	BlockType         requests.Integer `position:"Body" name:"BlockType"`
	PropertyType      string           `position:"Body" name:"PropertyType"`
	EntityId          requests.Integer `position:"Body" name:"EntityId"`
	RuleName          string           `position:"Body" name:"RuleName"`
	Checker           requests.Integer `position:"Body" name:"Checker"`
	Operator          string           `position:"Body" name:"Operator"`
	Property          string           `position:"Body" name:"Property"`
	Id                requests.Integer `position:"Body" name:"Id"`
	WarningThreshold  string           `position:"Body" name:"WarningThreshold"`
	MethodName        string           `position:"Body" name:"MethodName"`
	ProjectName       string           `position:"Body" name:"ProjectName"`
	RuleType          requests.Integer `position:"Body" name:"RuleType"`
	TemplateId        requests.Integer `position:"Body" name:"TemplateId"`
	ExpectValue       string           `position:"Body" name:"ExpectValue"`
	WhereCondition    string           `position:"Body" name:"WhereCondition"`
	CriticalThreshold string           `position:"Body" name:"CriticalThreshold"`
	Comment           string           `position:"Body" name:"Comment"`
	PredictType       requests.Integer `position:"Body" name:"PredictType"`
}

// UpdateQualityRuleResponse is the response struct for api UpdateQualityRule
type UpdateQualityRuleResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           bool   `json:"Data" xml:"Data"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateUpdateQualityRuleRequest creates a request to invoke UpdateQualityRule API
func CreateUpdateQualityRuleRequest() (request *UpdateQualityRuleRequest) {
	request = &UpdateQualityRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "UpdateQualityRule", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateQualityRuleResponse creates a response to parse from UpdateQualityRule response
func CreateUpdateQualityRuleResponse() (response *UpdateQualityRuleResponse) {
	response = &UpdateQualityRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
