package iot

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

// RerunJob invokes the iot.RerunJob API synchronously
func (client *Client) RerunJob(request *RerunJobRequest) (response *RerunJobResponse, err error) {
	response = CreateRerunJobResponse()
	err = client.DoAction(request, response)
	return
}

// RerunJobWithChan invokes the iot.RerunJob API asynchronously
func (client *Client) RerunJobWithChan(request *RerunJobRequest) (<-chan *RerunJobResponse, <-chan error) {
	responseChan := make(chan *RerunJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RerunJob(request)
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

// RerunJobWithCallback invokes the iot.RerunJob API asynchronously
func (client *Client) RerunJobWithCallback(request *RerunJobRequest, callback func(response *RerunJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RerunJobResponse
		var err error
		defer close(result)
		response, err = client.RerunJob(request)
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

// RerunJobRequest is the request struct for api RerunJob
type RerunJobRequest struct {
	*requests.RpcRequest
	JobId         string `position:"Query" name:"JobId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// RerunJobResponse is the response struct for api RerunJob
type RerunJobResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateRerunJobRequest creates a request to invoke RerunJob API
func CreateRerunJobRequest() (request *RerunJobRequest) {
	request = &RerunJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "RerunJob", "", "")
	request.Method = requests.POST
	return
}

// CreateRerunJobResponse creates a response to parse from RerunJob response
func CreateRerunJobResponse() (response *RerunJobResponse) {
	response = &RerunJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
