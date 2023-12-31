package dms_enterprise

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

// GetDataTrackOrderDetail invokes the dms_enterprise.GetDataTrackOrderDetail API synchronously
func (client *Client) GetDataTrackOrderDetail(request *GetDataTrackOrderDetailRequest) (response *GetDataTrackOrderDetailResponse, err error) {
	response = CreateGetDataTrackOrderDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetDataTrackOrderDetailWithChan invokes the dms_enterprise.GetDataTrackOrderDetail API asynchronously
func (client *Client) GetDataTrackOrderDetailWithChan(request *GetDataTrackOrderDetailRequest) (<-chan *GetDataTrackOrderDetailResponse, <-chan error) {
	responseChan := make(chan *GetDataTrackOrderDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDataTrackOrderDetail(request)
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

// GetDataTrackOrderDetailWithCallback invokes the dms_enterprise.GetDataTrackOrderDetail API asynchronously
func (client *Client) GetDataTrackOrderDetailWithCallback(request *GetDataTrackOrderDetailRequest, callback func(response *GetDataTrackOrderDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDataTrackOrderDetailResponse
		var err error
		defer close(result)
		response, err = client.GetDataTrackOrderDetail(request)
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

// GetDataTrackOrderDetailRequest is the request struct for api GetDataTrackOrderDetail
type GetDataTrackOrderDetailRequest struct {
	*requests.RpcRequest
	Tid     requests.Integer `position:"Query" name:"Tid"`
	OrderId requests.Integer `position:"Query" name:"OrderId"`
}

// GetDataTrackOrderDetailResponse is the response struct for api GetDataTrackOrderDetail
type GetDataTrackOrderDetailResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	Success              bool                 `json:"Success" xml:"Success"`
	ErrorMessage         string               `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode            string               `json:"ErrorCode" xml:"ErrorCode"`
	DataTrackOrderDetail DataTrackOrderDetail `json:"DataTrackOrderDetail" xml:"DataTrackOrderDetail"`
}

// CreateGetDataTrackOrderDetailRequest creates a request to invoke GetDataTrackOrderDetail API
func CreateGetDataTrackOrderDetailRequest() (request *GetDataTrackOrderDetailRequest) {
	request = &GetDataTrackOrderDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetDataTrackOrderDetail", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetDataTrackOrderDetailResponse creates a response to parse from GetDataTrackOrderDetail response
func CreateGetDataTrackOrderDetailResponse() (response *GetDataTrackOrderDetailResponse) {
	response = &GetDataTrackOrderDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
