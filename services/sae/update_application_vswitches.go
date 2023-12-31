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

// UpdateApplicationVswitches invokes the sae.UpdateApplicationVswitches API synchronously
func (client *Client) UpdateApplicationVswitches(request *UpdateApplicationVswitchesRequest) (response *UpdateApplicationVswitchesResponse, err error) {
	response = CreateUpdateApplicationVswitchesResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateApplicationVswitchesWithChan invokes the sae.UpdateApplicationVswitches API asynchronously
func (client *Client) UpdateApplicationVswitchesWithChan(request *UpdateApplicationVswitchesRequest) (<-chan *UpdateApplicationVswitchesResponse, <-chan error) {
	responseChan := make(chan *UpdateApplicationVswitchesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateApplicationVswitches(request)
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

// UpdateApplicationVswitchesWithCallback invokes the sae.UpdateApplicationVswitches API asynchronously
func (client *Client) UpdateApplicationVswitchesWithCallback(request *UpdateApplicationVswitchesRequest, callback func(response *UpdateApplicationVswitchesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateApplicationVswitchesResponse
		var err error
		defer close(result)
		response, err = client.UpdateApplicationVswitches(request)
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

// UpdateApplicationVswitchesRequest is the request struct for api UpdateApplicationVswitches
type UpdateApplicationVswitchesRequest struct {
	*requests.RoaRequest
	VSwitchId string `position:"Query" name:"VSwitchId"`
	AppId     string `position:"Query" name:"AppId"`
}

// UpdateApplicationVswitchesResponse is the response struct for api UpdateApplicationVswitches
type UpdateApplicationVswitchesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateApplicationVswitchesRequest creates a request to invoke UpdateApplicationVswitches API
func CreateUpdateApplicationVswitchesRequest() (request *UpdateApplicationVswitchesRequest) {
	request = &UpdateApplicationVswitchesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "UpdateApplicationVswitches", "/pop/v1/sam/app/updateAppVswitches", "serverless", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateApplicationVswitchesResponse creates a response to parse from UpdateApplicationVswitches response
func CreateUpdateApplicationVswitchesResponse() (response *UpdateApplicationVswitchesResponse) {
	response = &UpdateApplicationVswitchesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
