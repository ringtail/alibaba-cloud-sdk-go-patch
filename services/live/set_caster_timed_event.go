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

// SetCasterTimedEvent invokes the live.SetCasterTimedEvent API synchronously
func (client *Client) SetCasterTimedEvent(request *SetCasterTimedEventRequest) (response *SetCasterTimedEventResponse, err error) {
	response = CreateSetCasterTimedEventResponse()
	err = client.DoAction(request, response)
	return
}

// SetCasterTimedEventWithChan invokes the live.SetCasterTimedEvent API asynchronously
func (client *Client) SetCasterTimedEventWithChan(request *SetCasterTimedEventRequest) (<-chan *SetCasterTimedEventResponse, <-chan error) {
	responseChan := make(chan *SetCasterTimedEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetCasterTimedEvent(request)
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

// SetCasterTimedEventWithCallback invokes the live.SetCasterTimedEvent API asynchronously
func (client *Client) SetCasterTimedEventWithCallback(request *SetCasterTimedEventRequest, callback func(response *SetCasterTimedEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetCasterTimedEventResponse
		var err error
		defer close(result)
		response, err = client.SetCasterTimedEvent(request)
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

// SetCasterTimedEventRequest is the request struct for api SetCasterTimedEvent
type SetCasterTimedEventRequest struct {
	*requests.RpcRequest
	EventName    string           `position:"Query" name:"EventName"`
	StartTimeUTC string           `position:"Query" name:"StartTimeUTC"`
	CasterId     string           `position:"Query" name:"CasterId"`
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
}

// SetCasterTimedEventResponse is the response struct for api SetCasterTimedEvent
type SetCasterTimedEventResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetCasterTimedEventRequest creates a request to invoke SetCasterTimedEvent API
func CreateSetCasterTimedEventRequest() (request *SetCasterTimedEventRequest) {
	request = &SetCasterTimedEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "SetCasterTimedEvent", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetCasterTimedEventResponse creates a response to parse from SetCasterTimedEvent response
func CreateSetCasterTimedEventResponse() (response *SetCasterTimedEventResponse) {
	response = &SetCasterTimedEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
