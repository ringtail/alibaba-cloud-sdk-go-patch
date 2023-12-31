package ltl

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

// ListDeviceGroup invokes the ltl.ListDeviceGroup API synchronously
func (client *Client) ListDeviceGroup(request *ListDeviceGroupRequest) (response *ListDeviceGroupResponse, err error) {
	response = CreateListDeviceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ListDeviceGroupWithChan invokes the ltl.ListDeviceGroup API asynchronously
func (client *Client) ListDeviceGroupWithChan(request *ListDeviceGroupRequest) (<-chan *ListDeviceGroupResponse, <-chan error) {
	responseChan := make(chan *ListDeviceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDeviceGroup(request)
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

// ListDeviceGroupWithCallback invokes the ltl.ListDeviceGroup API asynchronously
func (client *Client) ListDeviceGroupWithCallback(request *ListDeviceGroupRequest, callback func(response *ListDeviceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDeviceGroupResponse
		var err error
		defer close(result)
		response, err = client.ListDeviceGroup(request)
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

// ListDeviceGroupRequest is the request struct for api ListDeviceGroup
type ListDeviceGroupRequest struct {
	*requests.RpcRequest
	Size       requests.Integer `position:"Query" name:"Size"`
	Num        requests.Integer `position:"Query" name:"Num"`
	ApiVersion string           `position:"Query" name:"ApiVersion"`
	ProductKey string           `position:"Query" name:"ProductKey"`
	BizChainId string           `position:"Query" name:"BizChainId"`
}

// ListDeviceGroupResponse is the response struct for api ListDeviceGroup
type ListDeviceGroupResponse struct {
	*responses.BaseResponse
	Code      int                   `json:"Code" xml:"Code"`
	Message   string                `json:"Message" xml:"Message"`
	RequestId string                `json:"RequestId" xml:"RequestId"`
	Success   bool                  `json:"Success" xml:"Success"`
	Data      DataInListDeviceGroup `json:"Data" xml:"Data"`
}

// CreateListDeviceGroupRequest creates a request to invoke ListDeviceGroup API
func CreateListDeviceGroupRequest() (request *ListDeviceGroupRequest) {
	request = &ListDeviceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ltl", "2019-05-10", "ListDeviceGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateListDeviceGroupResponse creates a response to parse from ListDeviceGroup response
func CreateListDeviceGroupResponse() (response *ListDeviceGroupResponse) {
	response = &ListDeviceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
