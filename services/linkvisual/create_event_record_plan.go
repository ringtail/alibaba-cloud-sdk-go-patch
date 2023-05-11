package linkvisual

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

// CreateEventRecordPlan invokes the linkvisual.CreateEventRecordPlan API synchronously
func (client *Client) CreateEventRecordPlan(request *CreateEventRecordPlanRequest) (response *CreateEventRecordPlanResponse, err error) {
	response = CreateCreateEventRecordPlanResponse()
	err = client.DoAction(request, response)
	return
}

// CreateEventRecordPlanWithChan invokes the linkvisual.CreateEventRecordPlan API asynchronously
func (client *Client) CreateEventRecordPlanWithChan(request *CreateEventRecordPlanRequest) (<-chan *CreateEventRecordPlanResponse, <-chan error) {
	responseChan := make(chan *CreateEventRecordPlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateEventRecordPlan(request)
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

// CreateEventRecordPlanWithCallback invokes the linkvisual.CreateEventRecordPlan API asynchronously
func (client *Client) CreateEventRecordPlanWithCallback(request *CreateEventRecordPlanRequest, callback func(response *CreateEventRecordPlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateEventRecordPlanResponse
		var err error
		defer close(result)
		response, err = client.CreateEventRecordPlan(request)
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

// CreateEventRecordPlanRequest is the request struct for api CreateEventRecordPlan
type CreateEventRecordPlanRequest struct {
	*requests.RpcRequest
	EventTypes        string           `position:"Query" name:"EventTypes"`
	PreRecordDuration requests.Integer `position:"Query" name:"PreRecordDuration"`
	RecordDuration    requests.Integer `position:"Query" name:"RecordDuration"`
	TemplateId        string           `position:"Query" name:"TemplateId"`
	ApiProduct        string           `position:"Body" name:"ApiProduct"`
	Name              string           `position:"Query" name:"Name"`
	ApiRevision       string           `position:"Body" name:"ApiRevision"`
}

// CreateEventRecordPlanResponse is the response struct for api CreateEventRecordPlan
type CreateEventRecordPlanResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	Data         string `json:"Data" xml:"Data"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateCreateEventRecordPlanRequest creates a request to invoke CreateEventRecordPlan API
func CreateCreateEventRecordPlanRequest() (request *CreateEventRecordPlanRequest) {
	request = &CreateEventRecordPlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "CreateEventRecordPlan", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateEventRecordPlanResponse creates a response to parse from CreateEventRecordPlan response
func CreateCreateEventRecordPlanResponse() (response *CreateEventRecordPlanResponse) {
	response = &CreateEventRecordPlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
