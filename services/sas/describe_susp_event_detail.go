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

// DescribeSuspEventDetail invokes the sas.DescribeSuspEventDetail API synchronously
func (client *Client) DescribeSuspEventDetail(request *DescribeSuspEventDetailRequest) (response *DescribeSuspEventDetailResponse, err error) {
	response = CreateDescribeSuspEventDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSuspEventDetailWithChan invokes the sas.DescribeSuspEventDetail API asynchronously
func (client *Client) DescribeSuspEventDetailWithChan(request *DescribeSuspEventDetailRequest) (<-chan *DescribeSuspEventDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeSuspEventDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSuspEventDetail(request)
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

// DescribeSuspEventDetailWithCallback invokes the sas.DescribeSuspEventDetail API asynchronously
func (client *Client) DescribeSuspEventDetailWithCallback(request *DescribeSuspEventDetailRequest, callback func(response *DescribeSuspEventDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSuspEventDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeSuspEventDetail(request)
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

// DescribeSuspEventDetailRequest is the request struct for api DescribeSuspEventDetail
type DescribeSuspEventDetailRequest struct {
	*requests.RpcRequest
	SourceIp          string           `position:"Query" name:"SourceIp"`
	From              string           `position:"Query" name:"From"`
	Lang              string           `position:"Query" name:"Lang"`
	SuspiciousEventId requests.Integer `position:"Query" name:"SuspiciousEventId"`
}

// DescribeSuspEventDetailResponse is the response struct for api DescribeSuspEventDetail
type DescribeSuspEventDetailResponse struct {
	*responses.BaseResponse
	Type             string      `json:"Type" xml:"Type"`
	DataSource       string      `json:"DataSource" xml:"DataSource"`
	EventName        string      `json:"EventName" xml:"EventName"`
	InternetIp       string      `json:"InternetIp" xml:"InternetIp"`
	IntranetIp       string      `json:"IntranetIp" xml:"IntranetIp"`
	LastTime         string      `json:"LastTime" xml:"LastTime"`
	OperateMsg       string      `json:"OperateMsg" xml:"OperateMsg"`
	Uuid             string      `json:"Uuid" xml:"Uuid"`
	CanBeDealOnLine  bool        `json:"CanBeDealOnLine" xml:"CanBeDealOnLine"`
	RequestId        string      `json:"RequestId" xml:"RequestId"`
	EventTypeDesc    string      `json:"EventTypeDesc" xml:"EventTypeDesc"`
	EventDesc        string      `json:"EventDesc" xml:"EventDesc"`
	InstanceName     string      `json:"InstanceName" xml:"InstanceName"`
	EventStatus      string      `json:"EventStatus" xml:"EventStatus"`
	SaleVersion      string      `json:"SaleVersion" xml:"SaleVersion"`
	OperateErrorCode string      `json:"OperateErrorCode" xml:"OperateErrorCode"`
	SasId            string      `json:"SasId" xml:"SasId"`
	Level            string      `json:"Level" xml:"Level"`
	Id               int         `json:"Id" xml:"Id"`
	Details          []QuaraFile `json:"Details" xml:"Details"`
}

// CreateDescribeSuspEventDetailRequest creates a request to invoke DescribeSuspEventDetail API
func CreateDescribeSuspEventDetailRequest() (request *DescribeSuspEventDetailRequest) {
	request = &DescribeSuspEventDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeSuspEventDetail", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSuspEventDetailResponse creates a response to parse from DescribeSuspEventDetail response
func CreateDescribeSuspEventDetailResponse() (response *DescribeSuspEventDetailResponse) {
	response = &DescribeSuspEventDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
