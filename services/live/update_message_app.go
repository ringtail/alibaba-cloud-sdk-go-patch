package live

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

// UpdateMessageApp invokes the live.UpdateMessageApp API synchronously
func (client *Client) UpdateMessageApp(request *UpdateMessageAppRequest) (response *UpdateMessageAppResponse, err error) {
	response = CreateUpdateMessageAppResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateMessageAppWithChan invokes the live.UpdateMessageApp API asynchronously
func (client *Client) UpdateMessageAppWithChan(request *UpdateMessageAppRequest) (<-chan *UpdateMessageAppResponse, <-chan error) {
	responseChan := make(chan *UpdateMessageAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateMessageApp(request)
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

// UpdateMessageAppWithCallback invokes the live.UpdateMessageApp API asynchronously
func (client *Client) UpdateMessageAppWithCallback(request *UpdateMessageAppRequest, callback func(response *UpdateMessageAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateMessageAppResponse
		var err error
		defer close(result)
		response, err = client.UpdateMessageApp(request)
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

// UpdateMessageAppRequest is the request struct for api UpdateMessageApp
type UpdateMessageAppRequest struct {
	*requests.RpcRequest
	Extension map[string]string `position:"Body" name:"Extension"  type:"Map"`
	AppConfig map[string]string `position:"Body" name:"AppConfig"  type:"Map"`
	AppName   string            `position:"Body" name:"AppName"`
	AppId     string            `position:"Body" name:"AppId"`
}

// UpdateMessageAppResponse is the response struct for api UpdateMessageApp
type UpdateMessageAppResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateUpdateMessageAppRequest creates a request to invoke UpdateMessageApp API
func CreateUpdateMessageAppRequest() (request *UpdateMessageAppRequest) {
	request = &UpdateMessageAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "UpdateMessageApp", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateMessageAppResponse creates a response to parse from UpdateMessageApp response
func CreateUpdateMessageAppResponse() (response *UpdateMessageAppResponse) {
	response = &UpdateMessageAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
