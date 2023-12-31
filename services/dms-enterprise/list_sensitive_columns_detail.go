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

// ListSensitiveColumnsDetail invokes the dms_enterprise.ListSensitiveColumnsDetail API synchronously
func (client *Client) ListSensitiveColumnsDetail(request *ListSensitiveColumnsDetailRequest) (response *ListSensitiveColumnsDetailResponse, err error) {
	response = CreateListSensitiveColumnsDetailResponse()
	err = client.DoAction(request, response)
	return
}

// ListSensitiveColumnsDetailWithChan invokes the dms_enterprise.ListSensitiveColumnsDetail API asynchronously
func (client *Client) ListSensitiveColumnsDetailWithChan(request *ListSensitiveColumnsDetailRequest) (<-chan *ListSensitiveColumnsDetailResponse, <-chan error) {
	responseChan := make(chan *ListSensitiveColumnsDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSensitiveColumnsDetail(request)
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

// ListSensitiveColumnsDetailWithCallback invokes the dms_enterprise.ListSensitiveColumnsDetail API asynchronously
func (client *Client) ListSensitiveColumnsDetailWithCallback(request *ListSensitiveColumnsDetailRequest, callback func(response *ListSensitiveColumnsDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSensitiveColumnsDetailResponse
		var err error
		defer close(result)
		response, err = client.ListSensitiveColumnsDetail(request)
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

// ListSensitiveColumnsDetailRequest is the request struct for api ListSensitiveColumnsDetail
type ListSensitiveColumnsDetailRequest struct {
	*requests.RpcRequest
	Tid        requests.Integer `position:"Query" name:"Tid"`
	TableName  string           `position:"Query" name:"TableName"`
	SchemaName string           `position:"Query" name:"SchemaName"`
	ColumnName string           `position:"Query" name:"ColumnName"`
	DbId       requests.Integer `position:"Query" name:"DbId"`
	Logic      requests.Boolean `position:"Query" name:"Logic"`
}

// ListSensitiveColumnsDetailResponse is the response struct for api ListSensitiveColumnsDetail
type ListSensitiveColumnsDetailResponse struct {
	*responses.BaseResponse
	RequestId                  string                     `json:"RequestId" xml:"RequestId"`
	ErrorCode                  string                     `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage               string                     `json:"ErrorMessage" xml:"ErrorMessage"`
	Success                    bool                       `json:"Success" xml:"Success"`
	SensitiveColumnsDetailList SensitiveColumnsDetailList `json:"SensitiveColumnsDetailList" xml:"SensitiveColumnsDetailList"`
}

// CreateListSensitiveColumnsDetailRequest creates a request to invoke ListSensitiveColumnsDetail API
func CreateListSensitiveColumnsDetailRequest() (request *ListSensitiveColumnsDetailRequest) {
	request = &ListSensitiveColumnsDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListSensitiveColumnsDetail", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListSensitiveColumnsDetailResponse creates a response to parse from ListSensitiveColumnsDetail response
func CreateListSensitiveColumnsDetailResponse() (response *ListSensitiveColumnsDetailResponse) {
	response = &ListSensitiveColumnsDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
