package vod

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

// GetImageInfos invokes the vod.GetImageInfos API synchronously
func (client *Client) GetImageInfos(request *GetImageInfosRequest) (response *GetImageInfosResponse, err error) {
	response = CreateGetImageInfosResponse()
	err = client.DoAction(request, response)
	return
}

// GetImageInfosWithChan invokes the vod.GetImageInfos API asynchronously
func (client *Client) GetImageInfosWithChan(request *GetImageInfosRequest) (<-chan *GetImageInfosResponse, <-chan error) {
	responseChan := make(chan *GetImageInfosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetImageInfos(request)
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

// GetImageInfosWithCallback invokes the vod.GetImageInfos API asynchronously
func (client *Client) GetImageInfosWithCallback(request *GetImageInfosRequest, callback func(response *GetImageInfosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetImageInfosResponse
		var err error
		defer close(result)
		response, err = client.GetImageInfos(request)
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

// GetImageInfosRequest is the request struct for api GetImageInfos
type GetImageInfosRequest struct {
	*requests.RpcRequest
	OutputType  string           `position:"Query" name:"OutputType"`
	AuthTimeout requests.Integer `position:"Query" name:"AuthTimeout"`
	ImageIds    string           `position:"Query" name:"ImageIds"`
}

// GetImageInfosResponse is the response struct for api GetImageInfos
type GetImageInfosResponse struct {
	*responses.BaseResponse
	RequestId        string   `json:"RequestId" xml:"RequestId"`
	NonExistImageIds []string `json:"NonExistImageIds" xml:"NonExistImageIds"`
	ImageInfo        []Image  `json:"ImageInfo" xml:"ImageInfo"`
}

// CreateGetImageInfosRequest creates a request to invoke GetImageInfos API
func CreateGetImageInfosRequest() (request *GetImageInfosRequest) {
	request = &GetImageInfosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetImageInfos", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetImageInfosResponse creates a response to parse from GetImageInfos response
func CreateGetImageInfosResponse() (response *GetImageInfosResponse) {
	response = &GetImageInfosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
