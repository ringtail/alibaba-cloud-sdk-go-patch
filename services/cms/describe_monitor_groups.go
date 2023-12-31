package cms

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

// DescribeMonitorGroups invokes the cms.DescribeMonitorGroups API synchronously
func (client *Client) DescribeMonitorGroups(request *DescribeMonitorGroupsRequest) (response *DescribeMonitorGroupsResponse, err error) {
	response = CreateDescribeMonitorGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMonitorGroupsWithChan invokes the cms.DescribeMonitorGroups API asynchronously
func (client *Client) DescribeMonitorGroupsWithChan(request *DescribeMonitorGroupsRequest) (<-chan *DescribeMonitorGroupsResponse, <-chan error) {
	responseChan := make(chan *DescribeMonitorGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMonitorGroups(request)
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

// DescribeMonitorGroupsWithCallback invokes the cms.DescribeMonitorGroups API asynchronously
func (client *Client) DescribeMonitorGroupsWithCallback(request *DescribeMonitorGroupsRequest, callback func(response *DescribeMonitorGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMonitorGroupsResponse
		var err error
		defer close(result)
		response, err = client.DescribeMonitorGroups(request)
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

// DescribeMonitorGroupsRequest is the request struct for api DescribeMonitorGroups
type DescribeMonitorGroupsRequest struct {
	*requests.RpcRequest
	SelectContactGroups    requests.Boolean            `position:"Query" name:"SelectContactGroups"`
	IncludeTemplateHistory requests.Boolean            `position:"Query" name:"IncludeTemplateHistory"`
	DynamicTagRuleId       string                      `position:"Query" name:"DynamicTagRuleId"`
	Type                   string                      `position:"Query" name:"Type"`
	PageNumber             requests.Integer            `position:"Query" name:"PageNumber"`
	ResourceGroupId        string                      `position:"Query" name:"ResourceGroupId"`
	GroupFounderTagKey     string                      `position:"Query" name:"GroupFounderTagKey"`
	PageSize               requests.Integer            `position:"Query" name:"PageSize"`
	GroupFounderTagValue   string                      `position:"Query" name:"GroupFounderTagValue"`
	Tag                    *[]DescribeMonitorGroupsTag `position:"Query" name:"Tag"  type:"Repeated"`
	Keyword                string                      `position:"Query" name:"Keyword"`
	Types                  string                      `position:"Query" name:"Types"`
	GroupId                string                      `position:"Query" name:"GroupId"`
	GroupName              string                      `position:"Query" name:"GroupName"`
	InstanceId             string                      `position:"Query" name:"InstanceId"`
	ServiceId              string                      `position:"Query" name:"ServiceId"`
}

// DescribeMonitorGroupsTag is a repeated param struct in DescribeMonitorGroupsRequest
type DescribeMonitorGroupsTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeMonitorGroupsResponse is the response struct for api DescribeMonitorGroups
type DescribeMonitorGroupsResponse struct {
	*responses.BaseResponse
	RequestId  string                           `json:"RequestId" xml:"RequestId"`
	Success    bool                             `json:"Success" xml:"Success"`
	Code       int                              `json:"Code" xml:"Code"`
	Message    string                           `json:"Message" xml:"Message"`
	PageNumber int                              `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                              `json:"PageSize" xml:"PageSize"`
	Total      int                              `json:"Total" xml:"Total"`
	Resources  ResourcesInDescribeMonitorGroups `json:"Resources" xml:"Resources"`
}

// CreateDescribeMonitorGroupsRequest creates a request to invoke DescribeMonitorGroups API
func CreateDescribeMonitorGroupsRequest() (request *DescribeMonitorGroupsRequest) {
	request = &DescribeMonitorGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeMonitorGroups", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeMonitorGroupsResponse creates a response to parse from DescribeMonitorGroups response
func CreateDescribeMonitorGroupsResponse() (response *DescribeMonitorGroupsResponse) {
	response = &DescribeMonitorGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
