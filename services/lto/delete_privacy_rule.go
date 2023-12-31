package lto

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

// DeletePrivacyRule invokes the lto.DeletePrivacyRule API synchronously
func (client *Client) DeletePrivacyRule(request *DeletePrivacyRuleRequest) (response *DeletePrivacyRuleResponse, err error) {
	response = CreateDeletePrivacyRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePrivacyRuleWithChan invokes the lto.DeletePrivacyRule API asynchronously
func (client *Client) DeletePrivacyRuleWithChan(request *DeletePrivacyRuleRequest) (<-chan *DeletePrivacyRuleResponse, <-chan error) {
	responseChan := make(chan *DeletePrivacyRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePrivacyRule(request)
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

// DeletePrivacyRuleWithCallback invokes the lto.DeletePrivacyRule API asynchronously
func (client *Client) DeletePrivacyRuleWithCallback(request *DeletePrivacyRuleRequest, callback func(response *DeletePrivacyRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePrivacyRuleResponse
		var err error
		defer close(result)
		response, err = client.DeletePrivacyRule(request)
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

// DeletePrivacyRuleRequest is the request struct for api DeletePrivacyRule
type DeletePrivacyRuleRequest struct {
	*requests.RpcRequest
	PrivacyRuleId string `position:"Query" name:"PrivacyRuleId"`
}

// DeletePrivacyRuleResponse is the response struct for api DeletePrivacyRule
type DeletePrivacyRuleResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateDeletePrivacyRuleRequest creates a request to invoke DeletePrivacyRule API
func CreateDeletePrivacyRuleRequest() (request *DeletePrivacyRuleRequest) {
	request = &DeletePrivacyRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("lto", "2021-07-07", "DeletePrivacyRule", "", "")
	request.Method = requests.POST
	return
}

// CreateDeletePrivacyRuleResponse creates a response to parse from DeletePrivacyRule response
func CreateDeletePrivacyRuleResponse() (response *DeletePrivacyRuleResponse) {
	response = &DeletePrivacyRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
