package ccc

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

// ListCampaigns invokes the ccc.ListCampaigns API synchronously
func (client *Client) ListCampaigns(request *ListCampaignsRequest) (response *ListCampaignsResponse, err error) {
	response = CreateListCampaignsResponse()
	err = client.DoAction(request, response)
	return
}

// ListCampaignsWithChan invokes the ccc.ListCampaigns API asynchronously
func (client *Client) ListCampaignsWithChan(request *ListCampaignsRequest) (<-chan *ListCampaignsResponse, <-chan error) {
	responseChan := make(chan *ListCampaignsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCampaigns(request)
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

// ListCampaignsWithCallback invokes the ccc.ListCampaigns API asynchronously
func (client *Client) ListCampaignsWithCallback(request *ListCampaignsRequest, callback func(response *ListCampaignsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCampaignsResponse
		var err error
		defer close(result)
		response, err = client.ListCampaigns(request)
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

// ListCampaignsRequest is the request struct for api ListCampaigns
type ListCampaignsRequest struct {
	*requests.RpcRequest
	ActualStartTimeTo   string           `position:"Query" name:"ActualStartTimeTo"`
	QueueId             string           `position:"Query" name:"QueueId"`
	ActualStartTimeFrom string           `position:"Query" name:"ActualStartTimeFrom"`
	PageNumber          requests.Integer `position:"Query" name:"PageNumber"`
	PlanedStartTimeFrom string           `position:"Query" name:"PlanedStartTimeFrom"`
	InstanceId          string           `position:"Query" name:"InstanceId"`
	Name                string           `position:"Query" name:"Name"`
	PageSize            requests.Integer `position:"Query" name:"PageSize"`
	PlanedStartTimeTo   string           `position:"Query" name:"PlanedStartTimeTo"`
	State               string           `position:"Query" name:"State"`
}

// ListCampaignsResponse is the response struct for api ListCampaigns
type ListCampaignsResponse struct {
	*responses.BaseResponse
	RequestId      string              `json:"RequestId" xml:"RequestId"`
	Message        string              `json:"Message" xml:"Message"`
	HttpStatusCode int64               `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string              `json:"Code" xml:"Code"`
	Success        bool                `json:"Success" xml:"Success"`
	Data           DataInListCampaigns `json:"Data" xml:"Data"`
}

// CreateListCampaignsRequest creates a request to invoke ListCampaigns API
func CreateListCampaignsRequest() (request *ListCampaignsRequest) {
	request = &ListCampaignsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListCampaigns", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListCampaignsResponse creates a response to parse from ListCampaigns response
func CreateListCampaignsResponse() (response *ListCampaignsResponse) {
	response = &ListCampaignsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
