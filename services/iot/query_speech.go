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

// QuerySpeech invokes the iot.QuerySpeech API synchronously
func (client *Client) QuerySpeech(request *QuerySpeechRequest) (response *QuerySpeechResponse, err error) {
	response = CreateQuerySpeechResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySpeechWithChan invokes the iot.QuerySpeech API asynchronously
func (client *Client) QuerySpeechWithChan(request *QuerySpeechRequest) (<-chan *QuerySpeechResponse, <-chan error) {
	responseChan := make(chan *QuerySpeechResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySpeech(request)
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

// QuerySpeechWithCallback invokes the iot.QuerySpeech API asynchronously
func (client *Client) QuerySpeechWithCallback(request *QuerySpeechRequest, callback func(response *QuerySpeechResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySpeechResponse
		var err error
		defer close(result)
		response, err = client.QuerySpeech(request)
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

// QuerySpeechRequest is the request struct for api QuerySpeech
type QuerySpeechRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Body" name:"IotInstanceId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	SpeechCode    string `position:"Body" name:"SpeechCode"`
}

// QuerySpeechResponse is the response struct for api QuerySpeech
type QuerySpeechResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateQuerySpeechRequest creates a request to invoke QuerySpeech API
func CreateQuerySpeechRequest() (request *QuerySpeechRequest) {
	request = &QuerySpeechRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QuerySpeech", "", "")
	request.Method = requests.POST
	return
}

// CreateQuerySpeechResponse creates a response to parse from QuerySpeech response
func CreateQuerySpeechResponse() (response *QuerySpeechResponse) {
	response = &QuerySpeechResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
