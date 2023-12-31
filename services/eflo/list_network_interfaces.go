package eflo

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

// ListNetworkInterfaces invokes the eflo.ListNetworkInterfaces API synchronously
func (client *Client) ListNetworkInterfaces(request *ListNetworkInterfacesRequest) (response *ListNetworkInterfacesResponse, err error) {
	response = CreateListNetworkInterfacesResponse()
	err = client.DoAction(request, response)
	return
}

// ListNetworkInterfacesWithChan invokes the eflo.ListNetworkInterfaces API asynchronously
func (client *Client) ListNetworkInterfacesWithChan(request *ListNetworkInterfacesRequest) (<-chan *ListNetworkInterfacesResponse, <-chan error) {
	responseChan := make(chan *ListNetworkInterfacesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNetworkInterfaces(request)
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

// ListNetworkInterfacesWithCallback invokes the eflo.ListNetworkInterfaces API asynchronously
func (client *Client) ListNetworkInterfacesWithCallback(request *ListNetworkInterfacesRequest, callback func(response *ListNetworkInterfacesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNetworkInterfacesResponse
		var err error
		defer close(result)
		response, err = client.ListNetworkInterfaces(request)
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

// ListNetworkInterfacesRequest is the request struct for api ListNetworkInterfaces
type ListNetworkInterfacesRequest struct {
	*requests.RpcRequest
	PageNumber         requests.Integer `position:"Body" name:"PageNumber"`
	PageSize           requests.Integer `position:"Body" name:"PageSize"`
	NodeId             string           `position:"Body" name:"NodeId"`
	SubnetId           string           `position:"Body" name:"SubnetId"`
	Ip                 string           `position:"Body" name:"Ip"`
	VpdId              string           `position:"Body" name:"VpdId"`
	EnablePage         requests.Boolean `position:"Body" name:"EnablePage"`
	NetworkInterfaceId string           `position:"Body" name:"NetworkInterfaceId"`
}

// ListNetworkInterfacesResponse is the response struct for api ListNetworkInterfaces
type ListNetworkInterfacesResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateListNetworkInterfacesRequest creates a request to invoke ListNetworkInterfaces API
func CreateListNetworkInterfacesRequest() (request *ListNetworkInterfacesRequest) {
	request = &ListNetworkInterfacesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "ListNetworkInterfaces", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListNetworkInterfacesResponse creates a response to parse from ListNetworkInterfaces response
func CreateListNetworkInterfacesResponse() (response *ListNetworkInterfacesResponse) {
	response = &ListNetworkInterfacesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
