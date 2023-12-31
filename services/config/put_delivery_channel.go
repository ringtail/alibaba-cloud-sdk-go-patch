package config

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

// PutDeliveryChannel invokes the config.PutDeliveryChannel API synchronously
func (client *Client) PutDeliveryChannel(request *PutDeliveryChannelRequest) (response *PutDeliveryChannelResponse, err error) {
	response = CreatePutDeliveryChannelResponse()
	err = client.DoAction(request, response)
	return
}

// PutDeliveryChannelWithChan invokes the config.PutDeliveryChannel API asynchronously
func (client *Client) PutDeliveryChannelWithChan(request *PutDeliveryChannelRequest) (<-chan *PutDeliveryChannelResponse, <-chan error) {
	responseChan := make(chan *PutDeliveryChannelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutDeliveryChannel(request)
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

// PutDeliveryChannelWithCallback invokes the config.PutDeliveryChannel API asynchronously
func (client *Client) PutDeliveryChannelWithCallback(request *PutDeliveryChannelRequest, callback func(response *PutDeliveryChannelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutDeliveryChannelResponse
		var err error
		defer close(result)
		response, err = client.PutDeliveryChannel(request)
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

// PutDeliveryChannelRequest is the request struct for api PutDeliveryChannel
type PutDeliveryChannelRequest struct {
	*requests.RpcRequest
	ClientToken                  string           `position:"Body" name:"ClientToken"`
	Description                  string           `position:"Body" name:"Description"`
	DeliveryChannelTargetArn     string           `position:"Body" name:"DeliveryChannelTargetArn"`
	DeliveryChannelCondition     string           `position:"Body" name:"DeliveryChannelCondition"`
	DeliveryChannelAssumeRoleArn string           `position:"Body" name:"DeliveryChannelAssumeRoleArn"`
	DeliveryChannelName          string           `position:"Body" name:"DeliveryChannelName"`
	DeliveryChannelId            string           `position:"Body" name:"DeliveryChannelId"`
	DeliveryChannelType          string           `position:"Body" name:"DeliveryChannelType"`
	Status                       requests.Integer `position:"Body" name:"Status"`
}

// PutDeliveryChannelResponse is the response struct for api PutDeliveryChannel
type PutDeliveryChannelResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	DeliveryChannelId string `json:"DeliveryChannelId" xml:"DeliveryChannelId"`
}

// CreatePutDeliveryChannelRequest creates a request to invoke PutDeliveryChannel API
func CreatePutDeliveryChannelRequest() (request *PutDeliveryChannelRequest) {
	request = &PutDeliveryChannelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2019-01-08", "PutDeliveryChannel", "", "")
	request.Method = requests.POST
	return
}

// CreatePutDeliveryChannelResponse creates a response to parse from PutDeliveryChannel response
func CreatePutDeliveryChannelResponse() (response *PutDeliveryChannelResponse) {
	response = &PutDeliveryChannelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
