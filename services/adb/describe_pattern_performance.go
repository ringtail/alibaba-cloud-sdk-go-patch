package adb

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

// DescribePatternPerformance invokes the adb.DescribePatternPerformance API synchronously
func (client *Client) DescribePatternPerformance(request *DescribePatternPerformanceRequest) (response *DescribePatternPerformanceResponse, err error) {
	response = CreateDescribePatternPerformanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePatternPerformanceWithChan invokes the adb.DescribePatternPerformance API asynchronously
func (client *Client) DescribePatternPerformanceWithChan(request *DescribePatternPerformanceRequest) (<-chan *DescribePatternPerformanceResponse, <-chan error) {
	responseChan := make(chan *DescribePatternPerformanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePatternPerformance(request)
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

// DescribePatternPerformanceWithCallback invokes the adb.DescribePatternPerformance API asynchronously
func (client *Client) DescribePatternPerformanceWithCallback(request *DescribePatternPerformanceRequest, callback func(response *DescribePatternPerformanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePatternPerformanceResponse
		var err error
		defer close(result)
		response, err = client.DescribePatternPerformance(request)
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

// DescribePatternPerformanceRequest is the request struct for api DescribePatternPerformance
type DescribePatternPerformanceRequest struct {
	*requests.RpcRequest
	DBClusterId string `position:"Query" name:"DBClusterId"`
	PatternId   string `position:"Query" name:"PatternId"`
	EndTime     string `position:"Query" name:"EndTime"`
	StartTime   string `position:"Query" name:"StartTime"`
}

// DescribePatternPerformanceResponse is the response struct for api DescribePatternPerformance
type DescribePatternPerformanceResponse struct {
	*responses.BaseResponse
	EndTime      string            `json:"EndTime" xml:"EndTime"`
	RequestId    string            `json:"RequestId" xml:"RequestId"`
	StartTime    string            `json:"StartTime" xml:"StartTime"`
	Performances []PerformanceItem `json:"Performances" xml:"Performances"`
}

// CreateDescribePatternPerformanceRequest creates a request to invoke DescribePatternPerformance API
func CreateDescribePatternPerformanceRequest() (request *DescribePatternPerformanceRequest) {
	request = &DescribePatternPerformanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "DescribePatternPerformance", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePatternPerformanceResponse creates a response to parse from DescribePatternPerformance response
func CreateDescribePatternPerformanceResponse() (response *DescribePatternPerformanceResponse) {
	response = &DescribePatternPerformanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
