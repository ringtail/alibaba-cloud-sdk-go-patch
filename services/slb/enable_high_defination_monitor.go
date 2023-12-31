package slb

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

// EnableHighDefinationMonitor invokes the slb.EnableHighDefinationMonitor API synchronously
func (client *Client) EnableHighDefinationMonitor(request *EnableHighDefinationMonitorRequest) (response *EnableHighDefinationMonitorResponse, err error) {
	response = CreateEnableHighDefinationMonitorResponse()
	err = client.DoAction(request, response)
	return
}

// EnableHighDefinationMonitorWithChan invokes the slb.EnableHighDefinationMonitor API asynchronously
func (client *Client) EnableHighDefinationMonitorWithChan(request *EnableHighDefinationMonitorRequest) (<-chan *EnableHighDefinationMonitorResponse, <-chan error) {
	responseChan := make(chan *EnableHighDefinationMonitorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableHighDefinationMonitor(request)
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

// EnableHighDefinationMonitorWithCallback invokes the slb.EnableHighDefinationMonitor API asynchronously
func (client *Client) EnableHighDefinationMonitorWithCallback(request *EnableHighDefinationMonitorRequest, callback func(response *EnableHighDefinationMonitorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableHighDefinationMonitorResponse
		var err error
		defer close(result)
		response, err = client.EnableHighDefinationMonitor(request)
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

// EnableHighDefinationMonitorRequest is the request struct for api EnableHighDefinationMonitor
type EnableHighDefinationMonitorRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	LogProject           string           `position:"Query" name:"LogProject"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tags                 string           `position:"Query" name:"Tags"`
	LogStore             string           `position:"Query" name:"LogStore"`
}

// EnableHighDefinationMonitorResponse is the response struct for api EnableHighDefinationMonitor
type EnableHighDefinationMonitorResponse struct {
	*responses.BaseResponse
	Success   string `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableHighDefinationMonitorRequest creates a request to invoke EnableHighDefinationMonitor API
func CreateEnableHighDefinationMonitorRequest() (request *EnableHighDefinationMonitorRequest) {
	request = &EnableHighDefinationMonitorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "EnableHighDefinationMonitor", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableHighDefinationMonitorResponse creates a response to parse from EnableHighDefinationMonitor response
func CreateEnableHighDefinationMonitorResponse() (response *EnableHighDefinationMonitorResponse) {
	response = &EnableHighDefinationMonitorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
