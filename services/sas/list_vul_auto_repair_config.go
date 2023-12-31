package sas

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

// ListVulAutoRepairConfig invokes the sas.ListVulAutoRepairConfig API synchronously
func (client *Client) ListVulAutoRepairConfig(request *ListVulAutoRepairConfigRequest) (response *ListVulAutoRepairConfigResponse, err error) {
	response = CreateListVulAutoRepairConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ListVulAutoRepairConfigWithChan invokes the sas.ListVulAutoRepairConfig API asynchronously
func (client *Client) ListVulAutoRepairConfigWithChan(request *ListVulAutoRepairConfigRequest) (<-chan *ListVulAutoRepairConfigResponse, <-chan error) {
	responseChan := make(chan *ListVulAutoRepairConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVulAutoRepairConfig(request)
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

// ListVulAutoRepairConfigWithCallback invokes the sas.ListVulAutoRepairConfig API asynchronously
func (client *Client) ListVulAutoRepairConfigWithCallback(request *ListVulAutoRepairConfigRequest, callback func(response *ListVulAutoRepairConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVulAutoRepairConfigResponse
		var err error
		defer close(result)
		response, err = client.ListVulAutoRepairConfig(request)
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

// ListVulAutoRepairConfigRequest is the request struct for api ListVulAutoRepairConfig
type ListVulAutoRepairConfigRequest struct {
	*requests.RpcRequest
	Type        string           `position:"Query" name:"Type"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	AliasName   string           `position:"Query" name:"AliasName"`
}

// ListVulAutoRepairConfigResponse is the response struct for api ListVulAutoRepairConfig
type ListVulAutoRepairConfigResponse struct {
	*responses.BaseResponse
	RequestId               string                `json:"RequestId" xml:"RequestId"`
	Success                 bool                  `json:"Success" xml:"Success"`
	Code                    string                `json:"Code" xml:"Code"`
	Message                 string                `json:"Message" xml:"Message"`
	HttpStatusCode          int                   `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PageInfo                PageInfo              `json:"PageInfo" xml:"PageInfo"`
	VulAutoRepairConfigList []VulAutoRepairConfig `json:"VulAutoRepairConfigList" xml:"VulAutoRepairConfigList"`
}

// CreateListVulAutoRepairConfigRequest creates a request to invoke ListVulAutoRepairConfig API
func CreateListVulAutoRepairConfigRequest() (request *ListVulAutoRepairConfigRequest) {
	request = &ListVulAutoRepairConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ListVulAutoRepairConfig", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListVulAutoRepairConfigResponse creates a response to parse from ListVulAutoRepairConfig response
func CreateListVulAutoRepairConfigResponse() (response *ListVulAutoRepairConfigResponse) {
	response = &ListVulAutoRepairConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
