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

// SetDevicePictureLifeCycle invokes the linkvisual.SetDevicePictureLifeCycle API synchronously
func (client *Client) SetDevicePictureLifeCycle(request *SetDevicePictureLifeCycleRequest) (response *SetDevicePictureLifeCycleResponse, err error) {
	response = CreateSetDevicePictureLifeCycleResponse()
	err = client.DoAction(request, response)
	return
}

// SetDevicePictureLifeCycleWithChan invokes the linkvisual.SetDevicePictureLifeCycle API asynchronously
func (client *Client) SetDevicePictureLifeCycleWithChan(request *SetDevicePictureLifeCycleRequest) (<-chan *SetDevicePictureLifeCycleResponse, <-chan error) {
	responseChan := make(chan *SetDevicePictureLifeCycleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDevicePictureLifeCycle(request)
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

// SetDevicePictureLifeCycleWithCallback invokes the linkvisual.SetDevicePictureLifeCycle API asynchronously
func (client *Client) SetDevicePictureLifeCycleWithCallback(request *SetDevicePictureLifeCycleRequest, callback func(response *SetDevicePictureLifeCycleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDevicePictureLifeCycleResponse
		var err error
		defer close(result)
		response, err = client.SetDevicePictureLifeCycle(request)
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

// SetDevicePictureLifeCycleRequest is the request struct for api SetDevicePictureLifeCycle
type SetDevicePictureLifeCycleRequest struct {
	*requests.RpcRequest
	IotId         string           `position:"Query" name:"IotId"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	Day           requests.Integer `position:"Query" name:"Day"`
	ProductKey    string           `position:"Query" name:"ProductKey"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	DeviceName    string           `position:"Query" name:"DeviceName"`
}

// SetDevicePictureLifeCycleResponse is the response struct for api SetDevicePictureLifeCycle
type SetDevicePictureLifeCycleResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateSetDevicePictureLifeCycleRequest creates a request to invoke SetDevicePictureLifeCycle API
func CreateSetDevicePictureLifeCycleRequest() (request *SetDevicePictureLifeCycleRequest) {
	request = &SetDevicePictureLifeCycleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "SetDevicePictureLifeCycle", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetDevicePictureLifeCycleResponse creates a response to parse from SetDevicePictureLifeCycle response
func CreateSetDevicePictureLifeCycleResponse() (response *SetDevicePictureLifeCycleResponse) {
	response = &SetDevicePictureLifeCycleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
