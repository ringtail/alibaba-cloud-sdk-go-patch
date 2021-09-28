package ccc

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

// ListAgentSummaryReports invokes the ccc.ListAgentSummaryReports API synchronously
func (client *Client) ListAgentSummaryReports(request *ListAgentSummaryReportsRequest) (response *ListAgentSummaryReportsResponse, err error) {
	response = CreateListAgentSummaryReportsResponse()
	err = client.DoAction(request, response)
	return
}

// ListAgentSummaryReportsWithChan invokes the ccc.ListAgentSummaryReports API asynchronously
func (client *Client) ListAgentSummaryReportsWithChan(request *ListAgentSummaryReportsRequest) (<-chan *ListAgentSummaryReportsResponse, <-chan error) {
	responseChan := make(chan *ListAgentSummaryReportsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAgentSummaryReports(request)
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

// ListAgentSummaryReportsWithCallback invokes the ccc.ListAgentSummaryReports API asynchronously
func (client *Client) ListAgentSummaryReportsWithCallback(request *ListAgentSummaryReportsRequest, callback func(response *ListAgentSummaryReportsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAgentSummaryReportsResponse
		var err error
		defer close(result)
		response, err = client.ListAgentSummaryReports(request)
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

// ListAgentSummaryReportsRequest is the request struct for api ListAgentSummaryReports
type ListAgentSummaryReportsRequest struct {
	*requests.RpcRequest
	AgentIds     string           `position:"Query" name:"AgentIds"`
	EndTime      string           `position:"Query" name:"EndTime"`
	StartTime    string           `position:"Query" name:"StartTime"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	SkillGroupId string           `position:"Query" name:"SkillGroupId"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
}

// ListAgentSummaryReportsResponse is the response struct for api ListAgentSummaryReports
type ListAgentSummaryReportsResponse struct {
	*responses.BaseResponse
	RequestId               string                  `json:"RequestId" xml:"RequestId"`
	Success                 bool                    `json:"Success" xml:"Success"`
	Code                    string                  `json:"Code" xml:"Code"`
	Message                 string                  `json:"Message" xml:"Message"`
	HttpStatusCode          int                     `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PagedAgentSummaryReport PagedAgentSummaryReport `json:"PagedAgentSummaryReport" xml:"PagedAgentSummaryReport"`
}

// CreateListAgentSummaryReportsRequest creates a request to invoke ListAgentSummaryReports API
func CreateListAgentSummaryReportsRequest() (request *ListAgentSummaryReportsRequest) {
	request = &ListAgentSummaryReportsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListAgentSummaryReports", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAgentSummaryReportsResponse creates a response to parse from ListAgentSummaryReports response
func CreateListAgentSummaryReportsResponse() (response *ListAgentSummaryReportsResponse) {
	response = &ListAgentSummaryReportsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
