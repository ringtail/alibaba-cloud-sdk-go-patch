package dataworks_public

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

// GetMetaTableIntroWiki invokes the dataworks_public.GetMetaTableIntroWiki API synchronously
func (client *Client) GetMetaTableIntroWiki(request *GetMetaTableIntroWikiRequest) (response *GetMetaTableIntroWikiResponse, err error) {
	response = CreateGetMetaTableIntroWikiResponse()
	err = client.DoAction(request, response)
	return
}

// GetMetaTableIntroWikiWithChan invokes the dataworks_public.GetMetaTableIntroWiki API asynchronously
func (client *Client) GetMetaTableIntroWikiWithChan(request *GetMetaTableIntroWikiRequest) (<-chan *GetMetaTableIntroWikiResponse, <-chan error) {
	responseChan := make(chan *GetMetaTableIntroWikiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMetaTableIntroWiki(request)
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

// GetMetaTableIntroWikiWithCallback invokes the dataworks_public.GetMetaTableIntroWiki API asynchronously
func (client *Client) GetMetaTableIntroWikiWithCallback(request *GetMetaTableIntroWikiRequest, callback func(response *GetMetaTableIntroWikiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMetaTableIntroWikiResponse
		var err error
		defer close(result)
		response, err = client.GetMetaTableIntroWiki(request)
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

// GetMetaTableIntroWikiRequest is the request struct for api GetMetaTableIntroWiki
type GetMetaTableIntroWikiRequest struct {
	*requests.RpcRequest
	WikiVersion requests.Integer `position:"Query" name:"WikiVersion"`
	TableGuid   string           `position:"Query" name:"TableGuid"`
}

// GetMetaTableIntroWikiResponse is the response struct for api GetMetaTableIntroWiki
type GetMetaTableIntroWikiResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetMetaTableIntroWikiRequest creates a request to invoke GetMetaTableIntroWiki API
func CreateGetMetaTableIntroWikiRequest() (request *GetMetaTableIntroWikiRequest) {
	request = &GetMetaTableIntroWikiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetMetaTableIntroWiki", "", "")
	request.Method = requests.POST
	return
}

// CreateGetMetaTableIntroWikiResponse creates a response to parse from GetMetaTableIntroWiki response
func CreateGetMetaTableIntroWikiResponse() (response *GetMetaTableIntroWikiResponse) {
	response = &GetMetaTableIntroWikiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
