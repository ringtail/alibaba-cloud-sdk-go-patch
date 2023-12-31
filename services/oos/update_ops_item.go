package oos

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

// UpdateOpsItem invokes the oos.UpdateOpsItem API synchronously
func (client *Client) UpdateOpsItem(request *UpdateOpsItemRequest) (response *UpdateOpsItemResponse, err error) {
	response = CreateUpdateOpsItemResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateOpsItemWithChan invokes the oos.UpdateOpsItem API asynchronously
func (client *Client) UpdateOpsItemWithChan(request *UpdateOpsItemRequest) (<-chan *UpdateOpsItemResponse, <-chan error) {
	responseChan := make(chan *UpdateOpsItemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateOpsItem(request)
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

// UpdateOpsItemWithCallback invokes the oos.UpdateOpsItem API asynchronously
func (client *Client) UpdateOpsItemWithCallback(request *UpdateOpsItemRequest, callback func(response *UpdateOpsItemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateOpsItemResponse
		var err error
		defer close(result)
		response, err = client.UpdateOpsItem(request)
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

// UpdateOpsItemRequest is the request struct for api UpdateOpsItem
type UpdateOpsItemRequest struct {
	*requests.RpcRequest
	ClientToken     string           `position:"Query" name:"ClientToken"`
	Description     string           `position:"Query" name:"Description"`
	Source          string           `position:"Query" name:"Source"`
	Title           string           `position:"Query" name:"Title"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	Severity        string           `position:"Query" name:"Severity"`
	Solutions       string           `position:"Query" name:"Solutions"`
	Resources       string           `position:"Query" name:"Resources"`
	Priority        requests.Integer `position:"Query" name:"Priority"`
	DedupString     string           `position:"Query" name:"DedupString"`
	Tags            string           `position:"Query" name:"Tags"`
	OpsItemId       string           `position:"Query" name:"OpsItemId"`
	Category        string           `position:"Query" name:"Category"`
	Status          string           `position:"Query" name:"Status"`
}

// UpdateOpsItemResponse is the response struct for api UpdateOpsItem
type UpdateOpsItemResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	OpsItem   OpsItemInUpdateOpsItem `json:"OpsItem" xml:"OpsItem"`
}

// CreateUpdateOpsItemRequest creates a request to invoke UpdateOpsItem API
func CreateUpdateOpsItemRequest() (request *UpdateOpsItemRequest) {
	request = &UpdateOpsItemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "UpdateOpsItem", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateOpsItemResponse creates a response to parse from UpdateOpsItem response
func CreateUpdateOpsItemResponse() (response *UpdateOpsItemResponse) {
	response = &UpdateOpsItemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
