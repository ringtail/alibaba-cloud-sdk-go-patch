package computenestsupplier

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

// ListServices invokes the computenestsupplier.ListServices API synchronously
func (client *Client) ListServices(request *ListServicesRequest) (response *ListServicesResponse, err error) {
	response = CreateListServicesResponse()
	err = client.DoAction(request, response)
	return
}

// ListServicesWithChan invokes the computenestsupplier.ListServices API asynchronously
func (client *Client) ListServicesWithChan(request *ListServicesRequest) (<-chan *ListServicesResponse, <-chan error) {
	responseChan := make(chan *ListServicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListServices(request)
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

// ListServicesWithCallback invokes the computenestsupplier.ListServices API asynchronously
func (client *Client) ListServicesWithCallback(request *ListServicesRequest, callback func(response *ListServicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListServicesResponse
		var err error
		defer close(result)
		response, err = client.ListServices(request)
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

// ListServicesRequest is the request struct for api ListServices
type ListServicesRequest struct {
	*requests.RpcRequest
	AllVersions     requests.Boolean      `position:"Query" name:"AllVersions"`
	ResourceGroupId string                `position:"Query" name:"ResourceGroupId"`
	NextToken       string                `position:"Query" name:"NextToken"`
	Tag             *[]ListServicesTag    `position:"Query" name:"Tag"  type:"Repeated"`
	Filter          *[]ListServicesFilter `position:"Query" name:"Filter"  type:"Repeated"`
	MaxResults      string                `position:"Query" name:"MaxResults"`
}

// ListServicesTag is a repeated param struct in ListServicesRequest
type ListServicesTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// ListServicesFilter is a repeated param struct in ListServicesRequest
type ListServicesFilter struct {
	Name  string    `name:"Name"`
	Value *[]string `name:"Value" type:"Repeated"`
}

// ListServicesResponse is the response struct for api ListServices
type ListServicesResponse struct {
	*responses.BaseResponse
	NextToken  string    `json:"NextToken" xml:"NextToken"`
	RequestId  string    `json:"RequestId" xml:"RequestId"`
	TotalCount string    `json:"TotalCount" xml:"TotalCount"`
	MaxResults int       `json:"MaxResults" xml:"MaxResults"`
	Services   []Service `json:"Services" xml:"Services"`
}

// CreateListServicesRequest creates a request to invoke ListServices API
func CreateListServicesRequest() (request *ListServicesRequest) {
	request = &ListServicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ComputeNestSupplier", "2021-05-21", "ListServices", "", "")
	request.Method = requests.POST
	return
}

// CreateListServicesResponse creates a response to parse from ListServices response
func CreateListServicesResponse() (response *ListServicesResponse) {
	response = &ListServicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
