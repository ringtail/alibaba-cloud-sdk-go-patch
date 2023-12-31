package cloudesl

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

// SaveTemplateGroup invokes the cloudesl.SaveTemplateGroup API synchronously
func (client *Client) SaveTemplateGroup(request *SaveTemplateGroupRequest) (response *SaveTemplateGroupResponse, err error) {
	response = CreateSaveTemplateGroupResponse()
	err = client.DoAction(request, response)
	return
}

// SaveTemplateGroupWithChan invokes the cloudesl.SaveTemplateGroup API asynchronously
func (client *Client) SaveTemplateGroupWithChan(request *SaveTemplateGroupRequest) (<-chan *SaveTemplateGroupResponse, <-chan error) {
	responseChan := make(chan *SaveTemplateGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveTemplateGroup(request)
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

// SaveTemplateGroupWithCallback invokes the cloudesl.SaveTemplateGroup API asynchronously
func (client *Client) SaveTemplateGroupWithCallback(request *SaveTemplateGroupRequest, callback func(response *SaveTemplateGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveTemplateGroupResponse
		var err error
		defer close(result)
		response, err = client.SaveTemplateGroup(request)
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

// SaveTemplateGroupRequest is the request struct for api SaveTemplateGroup
type SaveTemplateGroupRequest struct {
	*requests.RpcRequest
	TemplateVersion string           `position:"Body" name:"TemplateVersion"`
	EslModelId      string           `position:"Body" name:"EslModelId"`
	GroupId         string           `position:"Body" name:"GroupId"`
	WidthPx         requests.Integer `position:"Body" name:"WidthPx"`
	GroupName       string           `position:"Body" name:"GroupName"`
	HeightPx        requests.Integer `position:"Body" name:"HeightPx"`
}

// SaveTemplateGroupResponse is the response struct for api SaveTemplateGroup
type SaveTemplateGroupResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Message        string `json:"Message" xml:"Message"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	Code           string `json:"Code" xml:"Code"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
}

// CreateSaveTemplateGroupRequest creates a request to invoke SaveTemplateGroup API
func CreateSaveTemplateGroupRequest() (request *SaveTemplateGroupRequest) {
	request = &SaveTemplateGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "SaveTemplateGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveTemplateGroupResponse creates a response to parse from SaveTemplateGroup response
func CreateSaveTemplateGroupResponse() (response *SaveTemplateGroupResponse) {
	response = &SaveTemplateGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
