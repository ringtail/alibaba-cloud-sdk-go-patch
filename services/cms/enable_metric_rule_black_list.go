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

// EnableMetricRuleBlackList invokes the cms.EnableMetricRuleBlackList API synchronously
func (client *Client) EnableMetricRuleBlackList(request *EnableMetricRuleBlackListRequest) (response *EnableMetricRuleBlackListResponse, err error) {
	response = CreateEnableMetricRuleBlackListResponse()
	err = client.DoAction(request, response)
	return
}

// EnableMetricRuleBlackListWithChan invokes the cms.EnableMetricRuleBlackList API asynchronously
func (client *Client) EnableMetricRuleBlackListWithChan(request *EnableMetricRuleBlackListRequest) (<-chan *EnableMetricRuleBlackListResponse, <-chan error) {
	responseChan := make(chan *EnableMetricRuleBlackListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableMetricRuleBlackList(request)
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

// EnableMetricRuleBlackListWithCallback invokes the cms.EnableMetricRuleBlackList API asynchronously
func (client *Client) EnableMetricRuleBlackListWithCallback(request *EnableMetricRuleBlackListRequest, callback func(response *EnableMetricRuleBlackListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableMetricRuleBlackListResponse
		var err error
		defer close(result)
		response, err = client.EnableMetricRuleBlackList(request)
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

// EnableMetricRuleBlackListRequest is the request struct for api EnableMetricRuleBlackList
type EnableMetricRuleBlackListRequest struct {
	*requests.RpcRequest
	IsEnable requests.Boolean `position:"Query" name:"IsEnable"`
	Id       string           `position:"Query" name:"Id"`
}

// EnableMetricRuleBlackListResponse is the response struct for api EnableMetricRuleBlackList
type EnableMetricRuleBlackListResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Count     int    `json:"Count" xml:"Count"`
}

// CreateEnableMetricRuleBlackListRequest creates a request to invoke EnableMetricRuleBlackList API
func CreateEnableMetricRuleBlackListRequest() (request *EnableMetricRuleBlackListRequest) {
	request = &EnableMetricRuleBlackListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "EnableMetricRuleBlackList", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableMetricRuleBlackListResponse creates a response to parse from EnableMetricRuleBlackList response
func CreateEnableMetricRuleBlackListResponse() (response *EnableMetricRuleBlackListResponse) {
	response = &EnableMetricRuleBlackListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
