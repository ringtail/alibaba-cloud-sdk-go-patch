package dts

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

// SuspendDtsJobs invokes the dts.SuspendDtsJobs API synchronously
func (client *Client) SuspendDtsJobs(request *SuspendDtsJobsRequest) (response *SuspendDtsJobsResponse, err error) {
	response = CreateSuspendDtsJobsResponse()
	err = client.DoAction(request, response)
	return
}

// SuspendDtsJobsWithChan invokes the dts.SuspendDtsJobs API asynchronously
func (client *Client) SuspendDtsJobsWithChan(request *SuspendDtsJobsRequest) (<-chan *SuspendDtsJobsResponse, <-chan error) {
	responseChan := make(chan *SuspendDtsJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SuspendDtsJobs(request)
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

// SuspendDtsJobsWithCallback invokes the dts.SuspendDtsJobs API asynchronously
func (client *Client) SuspendDtsJobsWithCallback(request *SuspendDtsJobsRequest, callback func(response *SuspendDtsJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SuspendDtsJobsResponse
		var err error
		defer close(result)
		response, err = client.SuspendDtsJobs(request)
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

// SuspendDtsJobsRequest is the request struct for api SuspendDtsJobs
type SuspendDtsJobsRequest struct {
	*requests.RpcRequest
	DtsJobIds string `position:"Query" name:"DtsJobIds"`
}

// SuspendDtsJobsResponse is the response struct for api SuspendDtsJobs
type SuspendDtsJobsResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
}

// CreateSuspendDtsJobsRequest creates a request to invoke SuspendDtsJobs API
func CreateSuspendDtsJobsRequest() (request *SuspendDtsJobsRequest) {
	request = &SuspendDtsJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "SuspendDtsJobs", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSuspendDtsJobsResponse creates a response to parse from SuspendDtsJobs response
func CreateSuspendDtsJobsResponse() (response *SuspendDtsJobsResponse) {
	response = &SuspendDtsJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
