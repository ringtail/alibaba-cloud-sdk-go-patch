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

// CreateRecordPlan invokes the linkvisual.CreateRecordPlan API synchronously
func (client *Client) CreateRecordPlan(request *CreateRecordPlanRequest) (response *CreateRecordPlanResponse, err error) {
	response = CreateCreateRecordPlanResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRecordPlanWithChan invokes the linkvisual.CreateRecordPlan API asynchronously
func (client *Client) CreateRecordPlanWithChan(request *CreateRecordPlanRequest) (<-chan *CreateRecordPlanResponse, <-chan error) {
	responseChan := make(chan *CreateRecordPlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRecordPlan(request)
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

// CreateRecordPlanWithCallback invokes the linkvisual.CreateRecordPlan API asynchronously
func (client *Client) CreateRecordPlanWithCallback(request *CreateRecordPlanRequest, callback func(response *CreateRecordPlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRecordPlanResponse
		var err error
		defer close(result)
		response, err = client.CreateRecordPlan(request)
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

// CreateRecordPlanRequest is the request struct for api CreateRecordPlan
type CreateRecordPlanRequest struct {
	*requests.RpcRequest
	TemplateId  string `position:"Query" name:"TemplateId"`
	ApiProduct  string `position:"Body" name:"ApiProduct"`
	Name        string `position:"Query" name:"Name"`
	ApiRevision string `position:"Body" name:"ApiRevision"`
}

// CreateRecordPlanResponse is the response struct for api CreateRecordPlan
type CreateRecordPlanResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	Data         string `json:"Data" xml:"Data"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateCreateRecordPlanRequest creates a request to invoke CreateRecordPlan API
func CreateCreateRecordPlanRequest() (request *CreateRecordPlanRequest) {
	request = &CreateRecordPlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "CreateRecordPlan", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateRecordPlanResponse creates a response to parse from CreateRecordPlan response
func CreateCreateRecordPlanResponse() (response *CreateRecordPlanResponse) {
	response = &CreateRecordPlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
