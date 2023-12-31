package dts

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

// DescribeDataCheckReportUrl invokes the dts.DescribeDataCheckReportUrl API synchronously
func (client *Client) DescribeDataCheckReportUrl(request *DescribeDataCheckReportUrlRequest) (response *DescribeDataCheckReportUrlResponse, err error) {
	response = CreateDescribeDataCheckReportUrlResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataCheckReportUrlWithChan invokes the dts.DescribeDataCheckReportUrl API asynchronously
func (client *Client) DescribeDataCheckReportUrlWithChan(request *DescribeDataCheckReportUrlRequest) (<-chan *DescribeDataCheckReportUrlResponse, <-chan error) {
	responseChan := make(chan *DescribeDataCheckReportUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataCheckReportUrl(request)
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

// DescribeDataCheckReportUrlWithCallback invokes the dts.DescribeDataCheckReportUrl API asynchronously
func (client *Client) DescribeDataCheckReportUrlWithCallback(request *DescribeDataCheckReportUrlRequest, callback func(response *DescribeDataCheckReportUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataCheckReportUrlResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataCheckReportUrl(request)
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

// DescribeDataCheckReportUrlRequest is the request struct for api DescribeDataCheckReportUrl
type DescribeDataCheckReportUrlRequest struct {
	*requests.RpcRequest
	JobStepId string `position:"Query" name:"JobStepId"`
	TbName    string `position:"Query" name:"TbName"`
	DbName    string `position:"Query" name:"DbName"`
}

// DescribeDataCheckReportUrlResponse is the response struct for api DescribeDataCheckReportUrl
type DescribeDataCheckReportUrlResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	Success        string `json:"Success" xml:"Success"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateDescribeDataCheckReportUrlRequest creates a request to invoke DescribeDataCheckReportUrl API
func CreateDescribeDataCheckReportUrlRequest() (request *DescribeDataCheckReportUrlRequest) {
	request = &DescribeDataCheckReportUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "DescribeDataCheckReportUrl", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDataCheckReportUrlResponse creates a response to parse from DescribeDataCheckReportUrl response
func CreateDescribeDataCheckReportUrlResponse() (response *DescribeDataCheckReportUrlResponse) {
	response = &DescribeDataCheckReportUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
