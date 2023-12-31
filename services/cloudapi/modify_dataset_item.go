package cloudapi

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

// ModifyDatasetItem invokes the cloudapi.ModifyDatasetItem API synchronously
func (client *Client) ModifyDatasetItem(request *ModifyDatasetItemRequest) (response *ModifyDatasetItemResponse, err error) {
	response = CreateModifyDatasetItemResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDatasetItemWithChan invokes the cloudapi.ModifyDatasetItem API asynchronously
func (client *Client) ModifyDatasetItemWithChan(request *ModifyDatasetItemRequest) (<-chan *ModifyDatasetItemResponse, <-chan error) {
	responseChan := make(chan *ModifyDatasetItemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDatasetItem(request)
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

// ModifyDatasetItemWithCallback invokes the cloudapi.ModifyDatasetItem API asynchronously
func (client *Client) ModifyDatasetItemWithCallback(request *ModifyDatasetItemRequest, callback func(response *ModifyDatasetItemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDatasetItemResponse
		var err error
		defer close(result)
		response, err = client.ModifyDatasetItem(request)
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

// ModifyDatasetItemRequest is the request struct for api ModifyDatasetItem
type ModifyDatasetItemRequest struct {
	*requests.RpcRequest
	Description   string `position:"Query" name:"Description"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ExpiredTime   string `position:"Query" name:"ExpiredTime"`
	DatasetId     string `position:"Query" name:"DatasetId"`
	DatasetItemId string `position:"Query" name:"DatasetItemId"`
}

// ModifyDatasetItemResponse is the response struct for api ModifyDatasetItem
type ModifyDatasetItemResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDatasetItemRequest creates a request to invoke ModifyDatasetItem API
func CreateModifyDatasetItemRequest() (request *ModifyDatasetItemRequest) {
	request = &ModifyDatasetItemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyDatasetItem", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDatasetItemResponse creates a response to parse from ModifyDatasetItem response
func CreateModifyDatasetItemResponse() (response *ModifyDatasetItemResponse) {
	response = &ModifyDatasetItemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
