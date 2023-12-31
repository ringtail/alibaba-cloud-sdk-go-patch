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

// DescribeExportInfo invokes the sas.DescribeExportInfo API synchronously
func (client *Client) DescribeExportInfo(request *DescribeExportInfoRequest) (response *DescribeExportInfoResponse, err error) {
	response = CreateDescribeExportInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeExportInfoWithChan invokes the sas.DescribeExportInfo API asynchronously
func (client *Client) DescribeExportInfoWithChan(request *DescribeExportInfoRequest) (<-chan *DescribeExportInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeExportInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeExportInfo(request)
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

// DescribeExportInfoWithCallback invokes the sas.DescribeExportInfo API asynchronously
func (client *Client) DescribeExportInfoWithCallback(request *DescribeExportInfoRequest, callback func(response *DescribeExportInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeExportInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeExportInfo(request)
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

// DescribeExportInfoRequest is the request struct for api DescribeExportInfo
type DescribeExportInfoRequest struct {
	*requests.RpcRequest
	SourceIp string           `position:"Query" name:"SourceIp"`
	ExportId requests.Integer `position:"Query" name:"ExportId"`
}

// DescribeExportInfoResponse is the response struct for api DescribeExportInfo
type DescribeExportInfoResponse struct {
	*responses.BaseResponse
	Link         string `json:"Link" xml:"Link"`
	Progress     int    `json:"Progress" xml:"Progress"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	CurrentCount int    `json:"CurrentCount" xml:"CurrentCount"`
	Message      string `json:"Message" xml:"Message"`
	FileName     string `json:"FileName" xml:"FileName"`
	TotalCount   int    `json:"TotalCount" xml:"TotalCount"`
	ExportStatus string `json:"ExportStatus" xml:"ExportStatus"`
	Id           int64  `json:"Id" xml:"Id"`
}

// CreateDescribeExportInfoRequest creates a request to invoke DescribeExportInfo API
func CreateDescribeExportInfoRequest() (request *DescribeExportInfoRequest) {
	request = &DescribeExportInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeExportInfo", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeExportInfoResponse creates a response to parse from DescribeExportInfo response
func CreateDescribeExportInfoResponse() (response *DescribeExportInfoResponse) {
	response = &DescribeExportInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
