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

// GetDatabaseExportOrderDetail invokes the dms_enterprise.GetDatabaseExportOrderDetail API synchronously
func (client *Client) GetDatabaseExportOrderDetail(request *GetDatabaseExportOrderDetailRequest) (response *GetDatabaseExportOrderDetailResponse, err error) {
	response = CreateGetDatabaseExportOrderDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetDatabaseExportOrderDetailWithChan invokes the dms_enterprise.GetDatabaseExportOrderDetail API asynchronously
func (client *Client) GetDatabaseExportOrderDetailWithChan(request *GetDatabaseExportOrderDetailRequest) (<-chan *GetDatabaseExportOrderDetailResponse, <-chan error) {
	responseChan := make(chan *GetDatabaseExportOrderDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDatabaseExportOrderDetail(request)
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

// GetDatabaseExportOrderDetailWithCallback invokes the dms_enterprise.GetDatabaseExportOrderDetail API asynchronously
func (client *Client) GetDatabaseExportOrderDetailWithCallback(request *GetDatabaseExportOrderDetailRequest, callback func(response *GetDatabaseExportOrderDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDatabaseExportOrderDetailResponse
		var err error
		defer close(result)
		response, err = client.GetDatabaseExportOrderDetail(request)
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

// GetDatabaseExportOrderDetailRequest is the request struct for api GetDatabaseExportOrderDetail
type GetDatabaseExportOrderDetailRequest struct {
	*requests.RpcRequest
	Tid     requests.Integer `position:"Query" name:"Tid"`
	OrderId requests.Integer `position:"Body" name:"OrderId"`
}

// GetDatabaseExportOrderDetailResponse is the response struct for api GetDatabaseExportOrderDetail
type GetDatabaseExportOrderDetailResponse struct {
	*responses.BaseResponse
	RequestId                 string                    `json:"RequestId" xml:"RequestId"`
	ErrorCode                 string                    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage              string                    `json:"ErrorMessage" xml:"ErrorMessage"`
	Success                   bool                      `json:"Success" xml:"Success"`
	DatabaseExportOrderDetail DatabaseExportOrderDetail `json:"DatabaseExportOrderDetail" xml:"DatabaseExportOrderDetail"`
}

// CreateGetDatabaseExportOrderDetailRequest creates a request to invoke GetDatabaseExportOrderDetail API
func CreateGetDatabaseExportOrderDetailRequest() (request *GetDatabaseExportOrderDetailRequest) {
	request = &GetDatabaseExportOrderDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetDatabaseExportOrderDetail", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetDatabaseExportOrderDetailResponse creates a response to parse from GetDatabaseExportOrderDetail response
func CreateGetDatabaseExportOrderDetailResponse() (response *GetDatabaseExportOrderDetailResponse) {
	response = &GetDatabaseExportOrderDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
