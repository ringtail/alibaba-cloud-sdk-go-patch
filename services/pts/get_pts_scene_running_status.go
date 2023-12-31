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

// GetPtsSceneRunningStatus invokes the pts.GetPtsSceneRunningStatus API synchronously
func (client *Client) GetPtsSceneRunningStatus(request *GetPtsSceneRunningStatusRequest) (response *GetPtsSceneRunningStatusResponse, err error) {
	response = CreateGetPtsSceneRunningStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetPtsSceneRunningStatusWithChan invokes the pts.GetPtsSceneRunningStatus API asynchronously
func (client *Client) GetPtsSceneRunningStatusWithChan(request *GetPtsSceneRunningStatusRequest) (<-chan *GetPtsSceneRunningStatusResponse, <-chan error) {
	responseChan := make(chan *GetPtsSceneRunningStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPtsSceneRunningStatus(request)
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

// GetPtsSceneRunningStatusWithCallback invokes the pts.GetPtsSceneRunningStatus API asynchronously
func (client *Client) GetPtsSceneRunningStatusWithCallback(request *GetPtsSceneRunningStatusRequest, callback func(response *GetPtsSceneRunningStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPtsSceneRunningStatusResponse
		var err error
		defer close(result)
		response, err = client.GetPtsSceneRunningStatus(request)
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

// GetPtsSceneRunningStatusRequest is the request struct for api GetPtsSceneRunningStatus
type GetPtsSceneRunningStatusRequest struct {
	*requests.RpcRequest
	SceneId string `position:"Query" name:"SceneId"`
}

// GetPtsSceneRunningStatusResponse is the response struct for api GetPtsSceneRunningStatus
type GetPtsSceneRunningStatusResponse struct {
	*responses.BaseResponse
	Status         string `json:"Status" xml:"Status"`
	ModifiedTime   string `json:"ModifiedTime" xml:"ModifiedTime"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Message        string `json:"Message" xml:"Message"`
	SceneName      string `json:"SceneName" xml:"SceneName"`
	CreateTime     string `json:"CreateTime" xml:"CreateTime"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateGetPtsSceneRunningStatusRequest creates a request to invoke GetPtsSceneRunningStatus API
func CreateGetPtsSceneRunningStatusRequest() (request *GetPtsSceneRunningStatusRequest) {
	request = &GetPtsSceneRunningStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PTS", "2020-10-20", "GetPtsSceneRunningStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateGetPtsSceneRunningStatusResponse creates a response to parse from GetPtsSceneRunningStatus response
func CreateGetPtsSceneRunningStatusResponse() (response *GetPtsSceneRunningStatusResponse) {
	response = &GetPtsSceneRunningStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
