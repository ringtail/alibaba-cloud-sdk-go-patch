package polardb

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

// DeleteMaskingRules invokes the polardb.DeleteMaskingRules API synchronously
func (client *Client) DeleteMaskingRules(request *DeleteMaskingRulesRequest) (response *DeleteMaskingRulesResponse, err error) {
	response = CreateDeleteMaskingRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMaskingRulesWithChan invokes the polardb.DeleteMaskingRules API asynchronously
func (client *Client) DeleteMaskingRulesWithChan(request *DeleteMaskingRulesRequest) (<-chan *DeleteMaskingRulesResponse, <-chan error) {
	responseChan := make(chan *DeleteMaskingRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMaskingRules(request)
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

// DeleteMaskingRulesWithCallback invokes the polardb.DeleteMaskingRules API asynchronously
func (client *Client) DeleteMaskingRulesWithCallback(request *DeleteMaskingRulesRequest, callback func(response *DeleteMaskingRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMaskingRulesResponse
		var err error
		defer close(result)
		response, err = client.DeleteMaskingRules(request)
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

// DeleteMaskingRulesRequest is the request struct for api DeleteMaskingRules
type DeleteMaskingRulesRequest struct {
	*requests.RpcRequest
	DBClusterId  string `position:"Query" name:"DBClusterId"`
	RuleNameList string `position:"Query" name:"RuleNameList"`
}

// DeleteMaskingRulesResponse is the response struct for api DeleteMaskingRules
type DeleteMaskingRulesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteMaskingRulesRequest creates a request to invoke DeleteMaskingRules API
func CreateDeleteMaskingRulesRequest() (request *DeleteMaskingRulesRequest) {
	request = &DeleteMaskingRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DeleteMaskingRules", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteMaskingRulesResponse creates a response to parse from DeleteMaskingRules response
func CreateDeleteMaskingRulesResponse() (response *DeleteMaskingRulesResponse) {
	response = &DeleteMaskingRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
