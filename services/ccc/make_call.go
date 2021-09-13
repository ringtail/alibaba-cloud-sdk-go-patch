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

// MakeCall invokes the ccc.MakeCall API synchronously
func (client *Client) MakeCall(request *MakeCallRequest) (response *MakeCallResponse, err error) {
	response = CreateMakeCallResponse()
	err = client.DoAction(request, response)
	return
}

// MakeCallWithChan invokes the ccc.MakeCall API asynchronously
func (client *Client) MakeCallWithChan(request *MakeCallRequest) (<-chan *MakeCallResponse, <-chan error) {
	responseChan := make(chan *MakeCallResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MakeCall(request)
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

// MakeCallWithCallback invokes the ccc.MakeCall API asynchronously
func (client *Client) MakeCallWithCallback(request *MakeCallRequest, callback func(response *MakeCallResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MakeCallResponse
		var err error
		defer close(result)
		response, err = client.MakeCall(request)
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

// MakeCallRequest is the request struct for api MakeCall
type MakeCallRequest struct {
	*requests.RpcRequest
	Callee         string           `position:"Query" name:"Callee"`
	UserId         string           `position:"Query" name:"UserId"`
	DeviceId       string           `position:"Query" name:"DeviceId"`
	Tags           string           `position:"Query" name:"Tags"`
	TimeoutSeconds requests.Integer `position:"Query" name:"TimeoutSeconds"`
	Caller         string           `position:"Query" name:"Caller"`
	InstanceId     string           `position:"Query" name:"InstanceId"`
}

// MakeCallResponse is the response struct for api MakeCall
type MakeCallResponse struct {
	*responses.BaseResponse
	Code           string   `json:"Code" xml:"Code"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Params         []string `json:"Params" xml:"Params"`
	Data           Data     `json:"Data" xml:"Data"`
}

// CreateMakeCallRequest creates a request to invoke MakeCall API
func CreateMakeCallRequest() (request *MakeCallRequest) {
	request = &MakeCallRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "MakeCall", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateMakeCallResponse creates a response to parse from MakeCall response
func CreateMakeCallResponse() (response *MakeCallResponse) {
	response = &MakeCallResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
