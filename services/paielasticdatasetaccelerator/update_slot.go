package paielasticdatasetaccelerator

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

// UpdateSlot invokes the paielasticdatasetaccelerator.UpdateSlot API synchronously
func (client *Client) UpdateSlot(request *UpdateSlotRequest) (response *UpdateSlotResponse, err error) {
	response = CreateUpdateSlotResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateSlotWithChan invokes the paielasticdatasetaccelerator.UpdateSlot API asynchronously
func (client *Client) UpdateSlotWithChan(request *UpdateSlotRequest) (<-chan *UpdateSlotResponse, <-chan error) {
	responseChan := make(chan *UpdateSlotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateSlot(request)
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

// UpdateSlotWithCallback invokes the paielasticdatasetaccelerator.UpdateSlot API asynchronously
func (client *Client) UpdateSlotWithCallback(request *UpdateSlotRequest, callback func(response *UpdateSlotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateSlotResponse
		var err error
		defer close(result)
		response, err = client.UpdateSlot(request)
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

// UpdateSlotRequest is the request struct for api UpdateSlot
type UpdateSlotRequest struct {
	*requests.RoaRequest
	SlotId string `position:"Path" name:"SlotId"`
	Body   string `position:"Body" name:"body"`
}

// UpdateSlotResponse is the response struct for api UpdateSlot
type UpdateSlotResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateSlotRequest creates a request to invoke UpdateSlot API
func CreateUpdateSlotRequest() (request *UpdateSlotRequest) {
	request = &UpdateSlotRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PAIElasticDatasetAccelerator", "2022-08-01", "UpdateSlot", "/api/v1/slots/[SlotId]", "datasetacc", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateSlotResponse creates a response to parse from UpdateSlot response
func CreateUpdateSlotResponse() (response *UpdateSlotResponse) {
	response = &UpdateSlotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
