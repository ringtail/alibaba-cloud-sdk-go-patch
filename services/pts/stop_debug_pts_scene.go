package pts

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

// StopDebugPtsScene invokes the pts.StopDebugPtsScene API synchronously
func (client *Client) StopDebugPtsScene(request *StopDebugPtsSceneRequest) (response *StopDebugPtsSceneResponse, err error) {
	response = CreateStopDebugPtsSceneResponse()
	err = client.DoAction(request, response)
	return
}

// StopDebugPtsSceneWithChan invokes the pts.StopDebugPtsScene API asynchronously
func (client *Client) StopDebugPtsSceneWithChan(request *StopDebugPtsSceneRequest) (<-chan *StopDebugPtsSceneResponse, <-chan error) {
	responseChan := make(chan *StopDebugPtsSceneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopDebugPtsScene(request)
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

// StopDebugPtsSceneWithCallback invokes the pts.StopDebugPtsScene API asynchronously
func (client *Client) StopDebugPtsSceneWithCallback(request *StopDebugPtsSceneRequest, callback func(response *StopDebugPtsSceneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopDebugPtsSceneResponse
		var err error
		defer close(result)
		response, err = client.StopDebugPtsScene(request)
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

// StopDebugPtsSceneRequest is the request struct for api StopDebugPtsScene
type StopDebugPtsSceneRequest struct {
	*requests.RpcRequest
	SceneId string `position:"Query" name:"SceneId"`
	PlanId  string `position:"Query" name:"PlanId"`
}

// StopDebugPtsSceneResponse is the response struct for api StopDebugPtsScene
type StopDebugPtsSceneResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateStopDebugPtsSceneRequest creates a request to invoke StopDebugPtsScene API
func CreateStopDebugPtsSceneRequest() (request *StopDebugPtsSceneRequest) {
	request = &StopDebugPtsSceneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PTS", "2020-10-20", "StopDebugPtsScene", "", "")
	request.Method = requests.POST
	return
}

// CreateStopDebugPtsSceneResponse creates a response to parse from StopDebugPtsScene response
func CreateStopDebugPtsSceneResponse() (response *StopDebugPtsSceneResponse) {
	response = &StopDebugPtsSceneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
