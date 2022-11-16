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

// ListQualityRules invokes the dataworks_public.ListQualityRules API synchronously
func (client *Client) ListQualityRules(request *ListQualityRulesRequest) (response *ListQualityRulesResponse, err error) {
	response = CreateListQualityRulesResponse()
	err = client.DoAction(request, response)
	return
}

// ListQualityRulesWithChan invokes the dataworks_public.ListQualityRules API asynchronously
func (client *Client) ListQualityRulesWithChan(request *ListQualityRulesRequest) (<-chan *ListQualityRulesResponse, <-chan error) {
	responseChan := make(chan *ListQualityRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListQualityRules(request)
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

// ListQualityRulesWithCallback invokes the dataworks_public.ListQualityRules API asynchronously
func (client *Client) ListQualityRulesWithCallback(request *ListQualityRulesRequest, callback func(response *ListQualityRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListQualityRulesResponse
		var err error
		defer close(result)
		response, err = client.ListQualityRules(request)
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

// ListQualityRulesRequest is the request struct for api ListQualityRules
type ListQualityRulesRequest struct {
	*requests.RpcRequest
	ProjectName string           `position:"Body" name:"ProjectName"`
	EntityId    requests.Integer `position:"Body" name:"EntityId"`
	PageNumber  requests.Integer `position:"Body" name:"PageNumber"`
	PageSize    requests.Integer `position:"Body" name:"PageSize"`
}

// ListQualityRulesResponse is the response struct for api ListQualityRules
type ListQualityRulesResponse struct {
	*responses.BaseResponse
	HttpStatusCode int                    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string                 `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string                 `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool                   `json:"Success" xml:"Success"`
	ErrorCode      string                 `json:"ErrorCode" xml:"ErrorCode"`
	Data           DataInListQualityRules `json:"Data" xml:"Data"`
}

// CreateListQualityRulesRequest creates a request to invoke ListQualityRules API
func CreateListQualityRulesRequest() (request *ListQualityRulesRequest) {
	request = &ListQualityRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListQualityRules", "", "")
	request.Method = requests.POST
	return
}

// CreateListQualityRulesResponse creates a response to parse from ListQualityRules response
func CreateListQualityRulesResponse() (response *ListQualityRulesResponse) {
	response = &ListQualityRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
