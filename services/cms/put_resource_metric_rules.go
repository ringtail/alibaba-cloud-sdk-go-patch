package cms

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

// PutResourceMetricRules invokes the cms.PutResourceMetricRules API synchronously
func (client *Client) PutResourceMetricRules(request *PutResourceMetricRulesRequest) (response *PutResourceMetricRulesResponse, err error) {
	response = CreatePutResourceMetricRulesResponse()
	err = client.DoAction(request, response)
	return
}

// PutResourceMetricRulesWithChan invokes the cms.PutResourceMetricRules API asynchronously
func (client *Client) PutResourceMetricRulesWithChan(request *PutResourceMetricRulesRequest) (<-chan *PutResourceMetricRulesResponse, <-chan error) {
	responseChan := make(chan *PutResourceMetricRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutResourceMetricRules(request)
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

// PutResourceMetricRulesWithCallback invokes the cms.PutResourceMetricRules API asynchronously
func (client *Client) PutResourceMetricRulesWithCallback(request *PutResourceMetricRulesRequest, callback func(response *PutResourceMetricRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutResourceMetricRulesResponse
		var err error
		defer close(result)
		response, err = client.PutResourceMetricRules(request)
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

// PutResourceMetricRulesRequest is the request struct for api PutResourceMetricRules
type PutResourceMetricRulesRequest struct {
	*requests.RpcRequest
	Rules *[]PutResourceMetricRulesRules `position:"Query" name:"Rules"  type:"Repeated"`
}

// PutResourceMetricRulesRules is a repeated param struct in PutResourceMetricRulesRequest
type PutResourceMetricRulesRules struct {
	EscalationsInfoN                      string                               `name:"Escalations.Info.N"`
	Webhook                               string                               `name:"Webhook"`
	EscalationsWarnComparisonOperator     string                               `name:"Escalations.Warn.ComparisonOperator"`
	DynamicAlertSensitivity               string                               `name:"DynamicAlertSensitivity"`
	RuleName                              string                               `name:"RuleName"`
	EscalationsInfoStatistics             string                               `name:"Escalations.Info.Statistics"`
	EffectiveInterval                     string                               `name:"EffectiveInterval"`
	DynamicAlertHistoryDataRange          string                               `name:"DynamicAlertHistoryDataRange"`
	EscalationsWarnPreCondition           string                               `name:"Escalations.Warn.PreCondition"`
	EscalationsInfoComparisonOperator     string                               `name:"Escalations.Info.ComparisonOperator"`
	NoDataPolicy                          string                               `name:"NoDataPolicy"`
	NoEffectiveInterval                   string                               `name:"NoEffectiveInterval"`
	EmailSubject                          string                               `name:"EmailSubject"`
	Options                               string                               `name:"Options"`
	EscalationsCriticalN                  string                               `name:"Escalations.Critical.N"`
	SilenceTime                           string                               `name:"SilenceTime"`
	Prometheus                            string                               `name:"Prometheus"`
	EscalationsInfoPreCondition           string                               `name:"Escalations.Info.PreCondition"`
	MetricName                            string                               `name:"MetricName"`
	EscalationsWarnTimes                  string                               `name:"Escalations.Warn.Times"`
	CompositeExpression                   string                               `name:"CompositeExpression"`
	EscalationsWarnThreshold              string                               `name:"Escalations.Warn.Threshold"`
	Period                                string                               `name:"Period"`
	ContactGroups                         string                               `name:"ContactGroups"`
	EscalationsCriticalStatistics         string                               `name:"Escalations.Critical.Statistics"`
	RuleType                              string                               `name:"RuleType"`
	GroupId                               string                               `name:"GroupId"`
	EscalationsInfoTimes                  string                               `name:"Escalations.Info.Times"`
	Resources                             string                               `name:"Resources"`
	Labels                                *[]PutResourceMetricRulesRulesLabels `name:"Labels" type:"Repeated"`
	EscalationsCriticalTimes              string                               `name:"Escalations.Critical.Times"`
	EscalationsInfoThreshold              string                               `name:"Escalations.Info.Threshold"`
	EscalationsWarnStatistics             string                               `name:"Escalations.Warn.Statistics"`
	Namespace                             string                               `name:"Namespace"`
	EscalationsWarnN                      string                               `name:"Escalations.Warn.N"`
	Interval                              string                               `name:"Interval"`
	RuleId                                string                               `name:"RuleId"`
	EscalationsCriticalComparisonOperator string                               `name:"Escalations.Critical.ComparisonOperator"`
	EscalationsCriticalPreCondition       string                               `name:"Escalations.Critical.PreCondition"`
	EscalationsCriticalThreshold          string                               `name:"Escalations.Critical.Threshold"`
}

// PutResourceMetricRulesRulesLabels is a repeated param struct in PutResourceMetricRulesRequest
type PutResourceMetricRulesRulesLabels struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// PutResourceMetricRulesResponse is the response struct for api PutResourceMetricRules
type PutResourceMetricRulesResponse struct {
	*responses.BaseResponse
	Code             string           `json:"Code" xml:"Code"`
	Message          string           `json:"Message" xml:"Message"`
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	Success          bool             `json:"Success" xml:"Success"`
	FailedListResult FailedListResult `json:"FailedListResult" xml:"FailedListResult"`
}

// CreatePutResourceMetricRulesRequest creates a request to invoke PutResourceMetricRules API
func CreatePutResourceMetricRulesRequest() (request *PutResourceMetricRulesRequest) {
	request = &PutResourceMetricRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "PutResourceMetricRules", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePutResourceMetricRulesResponse creates a response to parse from PutResourceMetricRules response
func CreatePutResourceMetricRulesResponse() (response *PutResourceMetricRulesResponse) {
	response = &PutResourceMetricRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
