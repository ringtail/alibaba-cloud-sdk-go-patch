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

// ListEmrHiveAuditLogs invokes the dataworks_public.ListEmrHiveAuditLogs API synchronously
func (client *Client) ListEmrHiveAuditLogs(request *ListEmrHiveAuditLogsRequest) (response *ListEmrHiveAuditLogsResponse, err error) {
	response = CreateListEmrHiveAuditLogsResponse()
	err = client.DoAction(request, response)
	return
}

// ListEmrHiveAuditLogsWithChan invokes the dataworks_public.ListEmrHiveAuditLogs API asynchronously
func (client *Client) ListEmrHiveAuditLogsWithChan(request *ListEmrHiveAuditLogsRequest) (<-chan *ListEmrHiveAuditLogsResponse, <-chan error) {
	responseChan := make(chan *ListEmrHiveAuditLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListEmrHiveAuditLogs(request)
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

// ListEmrHiveAuditLogsWithCallback invokes the dataworks_public.ListEmrHiveAuditLogs API asynchronously
func (client *Client) ListEmrHiveAuditLogsWithCallback(request *ListEmrHiveAuditLogsRequest, callback func(response *ListEmrHiveAuditLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListEmrHiveAuditLogsResponse
		var err error
		defer close(result)
		response, err = client.ListEmrHiveAuditLogs(request)
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

// ListEmrHiveAuditLogsRequest is the request struct for api ListEmrHiveAuditLogs
type ListEmrHiveAuditLogsRequest struct {
	*requests.RpcRequest
	DatabaseName string           `position:"Query" name:"DatabaseName"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	EndTime      requests.Integer `position:"Query" name:"EndTime"`
	ClusterId    string           `position:"Query" name:"ClusterId"`
	StartTime    requests.Integer `position:"Query" name:"StartTime"`
	TableName    string           `position:"Query" name:"TableName"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
}

// ListEmrHiveAuditLogsResponse is the response struct for api ListEmrHiveAuditLogs
type ListEmrHiveAuditLogsResponse struct {
	*responses.BaseResponse
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateListEmrHiveAuditLogsRequest creates a request to invoke ListEmrHiveAuditLogs API
func CreateListEmrHiveAuditLogsRequest() (request *ListEmrHiveAuditLogsRequest) {
	request = &ListEmrHiveAuditLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2018-06-01", "ListEmrHiveAuditLogs", "", "")
	request.Method = requests.POST
	return
}

// CreateListEmrHiveAuditLogsResponse creates a response to parse from ListEmrHiveAuditLogs response
func CreateListEmrHiveAuditLogsResponse() (response *ListEmrHiveAuditLogsResponse) {
	response = &ListEmrHiveAuditLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}