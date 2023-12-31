package appstream_center

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

// GetResourceRenewPrice invokes the appstream_center.GetResourceRenewPrice API synchronously
func (client *Client) GetResourceRenewPrice(request *GetResourceRenewPriceRequest) (response *GetResourceRenewPriceResponse, err error) {
	response = CreateGetResourceRenewPriceResponse()
	err = client.DoAction(request, response)
	return
}

// GetResourceRenewPriceWithChan invokes the appstream_center.GetResourceRenewPrice API asynchronously
func (client *Client) GetResourceRenewPriceWithChan(request *GetResourceRenewPriceRequest) (<-chan *GetResourceRenewPriceResponse, <-chan error) {
	responseChan := make(chan *GetResourceRenewPriceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetResourceRenewPrice(request)
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

// GetResourceRenewPriceWithCallback invokes the appstream_center.GetResourceRenewPrice API asynchronously
func (client *Client) GetResourceRenewPriceWithCallback(request *GetResourceRenewPriceRequest, callback func(response *GetResourceRenewPriceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetResourceRenewPriceResponse
		var err error
		defer close(result)
		response, err = client.GetResourceRenewPrice(request)
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

// GetResourceRenewPriceRequest is the request struct for api GetResourceRenewPrice
type GetResourceRenewPriceRequest struct {
	*requests.RpcRequest
	Period             requests.Integer `position:"Query" name:"Period"`
	ProductType        string           `position:"Query" name:"ProductType"`
	PeriodUnit         string           `position:"Query" name:"PeriodUnit"`
	AppInstanceGroupId string           `position:"Query" name:"AppInstanceGroupId"`
}

// GetResourceRenewPriceResponse is the response struct for api GetResourceRenewPrice
type GetResourceRenewPriceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetResourceRenewPriceRequest creates a request to invoke GetResourceRenewPrice API
func CreateGetResourceRenewPriceRequest() (request *GetResourceRenewPriceRequest) {
	request = &GetResourceRenewPriceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("appstream-center", "2021-09-01", "GetResourceRenewPrice", "", "")
	request.Method = requests.POST
	return
}

// CreateGetResourceRenewPriceResponse creates a response to parse from GetResourceRenewPrice response
func CreateGetResourceRenewPriceResponse() (response *GetResourceRenewPriceResponse) {
	response = &GetResourceRenewPriceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
