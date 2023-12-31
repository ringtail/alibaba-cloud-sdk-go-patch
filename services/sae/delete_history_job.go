package sae

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

// DeleteHistoryJob invokes the sae.DeleteHistoryJob API synchronously
func (client *Client) DeleteHistoryJob(request *DeleteHistoryJobRequest) (response *DeleteHistoryJobResponse, err error) {
	response = CreateDeleteHistoryJobResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteHistoryJobWithChan invokes the sae.DeleteHistoryJob API asynchronously
func (client *Client) DeleteHistoryJobWithChan(request *DeleteHistoryJobRequest) (<-chan *DeleteHistoryJobResponse, <-chan error) {
	responseChan := make(chan *DeleteHistoryJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteHistoryJob(request)
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

// DeleteHistoryJobWithCallback invokes the sae.DeleteHistoryJob API asynchronously
func (client *Client) DeleteHistoryJobWithCallback(request *DeleteHistoryJobRequest, callback func(response *DeleteHistoryJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteHistoryJobResponse
		var err error
		defer close(result)
		response, err = client.DeleteHistoryJob(request)
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

// DeleteHistoryJobRequest is the request struct for api DeleteHistoryJob
type DeleteHistoryJobRequest struct {
	*requests.RoaRequest
	JobId string `position:"Query" name:"JobId"`
	AppId string `position:"Query" name:"AppId"`
}

// DeleteHistoryJobResponse is the response struct for api DeleteHistoryJob
type DeleteHistoryJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteHistoryJobRequest creates a request to invoke DeleteHistoryJob API
func CreateDeleteHistoryJobRequest() (request *DeleteHistoryJobRequest) {
	request = &DeleteHistoryJobRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DeleteHistoryJob", "/pop/v1/sam/job/deleteHistoryJob", "serverless", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteHistoryJobResponse creates a response to parse from DeleteHistoryJob response
func CreateDeleteHistoryJobResponse() (response *DeleteHistoryJobResponse) {
	response = &DeleteHistoryJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
