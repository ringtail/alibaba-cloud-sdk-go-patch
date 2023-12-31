package outboundbot

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

// SaveBaseStrategyPeriod invokes the outboundbot.SaveBaseStrategyPeriod API synchronously
func (client *Client) SaveBaseStrategyPeriod(request *SaveBaseStrategyPeriodRequest) (response *SaveBaseStrategyPeriodResponse, err error) {
	response = CreateSaveBaseStrategyPeriodResponse()
	err = client.DoAction(request, response)
	return
}

// SaveBaseStrategyPeriodWithChan invokes the outboundbot.SaveBaseStrategyPeriod API asynchronously
func (client *Client) SaveBaseStrategyPeriodWithChan(request *SaveBaseStrategyPeriodRequest) (<-chan *SaveBaseStrategyPeriodResponse, <-chan error) {
	responseChan := make(chan *SaveBaseStrategyPeriodResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveBaseStrategyPeriod(request)
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

// SaveBaseStrategyPeriodWithCallback invokes the outboundbot.SaveBaseStrategyPeriod API asynchronously
func (client *Client) SaveBaseStrategyPeriodWithCallback(request *SaveBaseStrategyPeriodRequest, callback func(response *SaveBaseStrategyPeriodResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveBaseStrategyPeriodResponse
		var err error
		defer close(result)
		response, err = client.SaveBaseStrategyPeriod(request)
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

// SaveBaseStrategyPeriodRequest is the request struct for api SaveBaseStrategyPeriod
type SaveBaseStrategyPeriodRequest struct {
	*requests.RpcRequest
	StrategyLevel         requests.Integer `position:"Query" name:"StrategyLevel"`
	EntryId               string           `position:"Query" name:"EntryId"`
	OnlyWeekdays          requests.Boolean `position:"Query" name:"OnlyWeekdays"`
	WorkingTimeFramesJson string           `position:"Query" name:"WorkingTimeFramesJson"`
	WorkingTime           *[]string        `position:"Query" name:"WorkingTime"  type:"Repeated"`
}

// SaveBaseStrategyPeriodResponse is the response struct for api SaveBaseStrategyPeriod
type SaveBaseStrategyPeriodResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateSaveBaseStrategyPeriodRequest creates a request to invoke SaveBaseStrategyPeriod API
func CreateSaveBaseStrategyPeriodRequest() (request *SaveBaseStrategyPeriodRequest) {
	request = &SaveBaseStrategyPeriodRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "SaveBaseStrategyPeriod", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveBaseStrategyPeriodResponse creates a response to parse from SaveBaseStrategyPeriod response
func CreateSaveBaseStrategyPeriodResponse() (response *SaveBaseStrategyPeriodResponse) {
	response = &SaveBaseStrategyPeriodResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
