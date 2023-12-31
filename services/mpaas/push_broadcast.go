package mpaas

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

// PushBroadcast invokes the mpaas.PushBroadcast API synchronously
func (client *Client) PushBroadcast(request *PushBroadcastRequest) (response *PushBroadcastResponse, err error) {
	response = CreatePushBroadcastResponse()
	err = client.DoAction(request, response)
	return
}

// PushBroadcastWithChan invokes the mpaas.PushBroadcast API asynchronously
func (client *Client) PushBroadcastWithChan(request *PushBroadcastRequest) (<-chan *PushBroadcastResponse, <-chan error) {
	responseChan := make(chan *PushBroadcastResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PushBroadcast(request)
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

// PushBroadcastWithCallback invokes the mpaas.PushBroadcast API asynchronously
func (client *Client) PushBroadcastWithCallback(request *PushBroadcastRequest, callback func(response *PushBroadcastResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PushBroadcastResponse
		var err error
		defer close(result)
		response, err = client.PushBroadcast(request)
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

// PushBroadcastRequest is the request struct for api PushBroadcast
type PushBroadcastRequest struct {
	*requests.RpcRequest
	AndroidChannel   requests.Integer `position:"Body" name:"AndroidChannel"`
	TaskName         string           `position:"Body" name:"TaskName"`
	TemplateKeyValue string           `position:"Body" name:"TemplateKeyValue"`
	PushAction       requests.Integer `position:"Body" name:"PushAction"`
	DeliveryType     requests.Integer `position:"Body" name:"DeliveryType"`
	TemplateName     string           `position:"Body" name:"TemplateName"`
	NotifyType       string           `position:"Body" name:"NotifyType"`
	ExtendedParams   string           `position:"Body" name:"ExtendedParams"`
	Silent           requests.Integer `position:"Body" name:"Silent"`
	StrategyContent  string           `position:"Body" name:"StrategyContent"`
	PushStatus       requests.Integer `position:"Body" name:"PushStatus"`
	Classification   string           `position:"Body" name:"Classification"`
	UnBindPeriod     requests.Integer `position:"Body" name:"UnBindPeriod"`
	ExpiredSeconds   requests.Integer `position:"Body" name:"ExpiredSeconds"`
	AppId            string           `position:"Body" name:"AppId"`
	Msgkey           string           `position:"Body" name:"Msgkey"`
	StrategyType     requests.Integer `position:"Body" name:"StrategyType"`
	BindPeriod       requests.Integer `position:"Body" name:"BindPeriod"`
	WorkspaceId      string           `position:"Body" name:"WorkspaceId"`
}

// PushBroadcastResponse is the response struct for api PushBroadcast
type PushBroadcastResponse struct {
	*responses.BaseResponse
	ResultMessage string     `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string     `json:"ResultCode" xml:"ResultCode"`
	RequestId     string     `json:"RequestId" xml:"RequestId"`
	PushResult    PushResult `json:"PushResult" xml:"PushResult"`
}

// CreatePushBroadcastRequest creates a request to invoke PushBroadcast API
func CreatePushBroadcastRequest() (request *PushBroadcastRequest) {
	request = &PushBroadcastRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "PushBroadcast", "", "")
	request.Method = requests.POST
	return
}

// CreatePushBroadcastResponse creates a response to parse from PushBroadcast response
func CreatePushBroadcastResponse() (response *PushBroadcastResponse) {
	response = &PushBroadcastResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
