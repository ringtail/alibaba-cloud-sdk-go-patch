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

// EnableSiteMonitors invokes the cms.EnableSiteMonitors API synchronously
func (client *Client) EnableSiteMonitors(request *EnableSiteMonitorsRequest) (response *EnableSiteMonitorsResponse, err error) {
	response = CreateEnableSiteMonitorsResponse()
	err = client.DoAction(request, response)
	return
}

// EnableSiteMonitorsWithChan invokes the cms.EnableSiteMonitors API asynchronously
func (client *Client) EnableSiteMonitorsWithChan(request *EnableSiteMonitorsRequest) (<-chan *EnableSiteMonitorsResponse, <-chan error) {
	responseChan := make(chan *EnableSiteMonitorsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableSiteMonitors(request)
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

// EnableSiteMonitorsWithCallback invokes the cms.EnableSiteMonitors API asynchronously
func (client *Client) EnableSiteMonitorsWithCallback(request *EnableSiteMonitorsRequest, callback func(response *EnableSiteMonitorsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableSiteMonitorsResponse
		var err error
		defer close(result)
		response, err = client.EnableSiteMonitors(request)
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

// EnableSiteMonitorsRequest is the request struct for api EnableSiteMonitors
type EnableSiteMonitorsRequest struct {
	*requests.RpcRequest
	TaskIds string `position:"Query" name:"TaskIds"`
}

// EnableSiteMonitorsResponse is the response struct for api EnableSiteMonitors
type EnableSiteMonitorsResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateEnableSiteMonitorsRequest creates a request to invoke EnableSiteMonitors API
func CreateEnableSiteMonitorsRequest() (request *EnableSiteMonitorsRequest) {
	request = &EnableSiteMonitorsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "EnableSiteMonitors", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableSiteMonitorsResponse creates a response to parse from EnableSiteMonitors response
func CreateEnableSiteMonitorsResponse() (response *EnableSiteMonitorsResponse) {
	response = &EnableSiteMonitorsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
