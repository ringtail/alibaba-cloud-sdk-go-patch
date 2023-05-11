package linkvisual

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

// SetDeviceRecordLifeCycle invokes the linkvisual.SetDeviceRecordLifeCycle API synchronously
func (client *Client) SetDeviceRecordLifeCycle(request *SetDeviceRecordLifeCycleRequest) (response *SetDeviceRecordLifeCycleResponse, err error) {
	response = CreateSetDeviceRecordLifeCycleResponse()
	err = client.DoAction(request, response)
	return
}

// SetDeviceRecordLifeCycleWithChan invokes the linkvisual.SetDeviceRecordLifeCycle API asynchronously
func (client *Client) SetDeviceRecordLifeCycleWithChan(request *SetDeviceRecordLifeCycleRequest) (<-chan *SetDeviceRecordLifeCycleResponse, <-chan error) {
	responseChan := make(chan *SetDeviceRecordLifeCycleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDeviceRecordLifeCycle(request)
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

// SetDeviceRecordLifeCycleWithCallback invokes the linkvisual.SetDeviceRecordLifeCycle API asynchronously
func (client *Client) SetDeviceRecordLifeCycleWithCallback(request *SetDeviceRecordLifeCycleRequest, callback func(response *SetDeviceRecordLifeCycleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDeviceRecordLifeCycleResponse
		var err error
		defer close(result)
		response, err = client.SetDeviceRecordLifeCycle(request)
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

// SetDeviceRecordLifeCycleRequest is the request struct for api SetDeviceRecordLifeCycle
type SetDeviceRecordLifeCycleRequest struct {
	*requests.RpcRequest
	IotId         string           `position:"Query" name:"IotId"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	Day           requests.Integer `position:"Query" name:"Day"`
	ProductKey    string           `position:"Query" name:"ProductKey"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	DeviceName    string           `position:"Query" name:"DeviceName"`
}

// SetDeviceRecordLifeCycleResponse is the response struct for api SetDeviceRecordLifeCycle
type SetDeviceRecordLifeCycleResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateSetDeviceRecordLifeCycleRequest creates a request to invoke SetDeviceRecordLifeCycle API
func CreateSetDeviceRecordLifeCycleRequest() (request *SetDeviceRecordLifeCycleRequest) {
	request = &SetDeviceRecordLifeCycleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "SetDeviceRecordLifeCycle", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetDeviceRecordLifeCycleResponse creates a response to parse from SetDeviceRecordLifeCycle response
func CreateSetDeviceRecordLifeCycleResponse() (response *SetDeviceRecordLifeCycleResponse) {
	response = &SetDeviceRecordLifeCycleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
