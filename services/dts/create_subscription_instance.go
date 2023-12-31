package dts

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

// CreateSubscriptionInstance invokes the dts.CreateSubscriptionInstance API synchronously
func (client *Client) CreateSubscriptionInstance(request *CreateSubscriptionInstanceRequest) (response *CreateSubscriptionInstanceResponse, err error) {
	response = CreateCreateSubscriptionInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSubscriptionInstanceWithChan invokes the dts.CreateSubscriptionInstance API asynchronously
func (client *Client) CreateSubscriptionInstanceWithChan(request *CreateSubscriptionInstanceRequest) (<-chan *CreateSubscriptionInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateSubscriptionInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSubscriptionInstance(request)
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

// CreateSubscriptionInstanceWithCallback invokes the dts.CreateSubscriptionInstance API asynchronously
func (client *Client) CreateSubscriptionInstanceWithCallback(request *CreateSubscriptionInstanceRequest, callback func(response *CreateSubscriptionInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSubscriptionInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateSubscriptionInstance(request)
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

// CreateSubscriptionInstanceRequest is the request struct for api CreateSubscriptionInstance
type CreateSubscriptionInstanceRequest struct {
	*requests.RpcRequest
	ClientToken                string           `position:"Query" name:"ClientToken"`
	SourceEndpointInstanceType string           `position:"Query" name:"SourceEndpoint.InstanceType"`
	AccountId                  string           `position:"Query" name:"AccountId"`
	Period                     string           `position:"Query" name:"Period"`
	OwnerId                    string           `position:"Query" name:"OwnerId"`
	UsedTime                   requests.Integer `position:"Query" name:"UsedTime"`
	Region                     string           `position:"Query" name:"Region"`
	PayType                    string           `position:"Query" name:"PayType"`
}

// CreateSubscriptionInstanceResponse is the response struct for api CreateSubscriptionInstance
type CreateSubscriptionInstanceResponse struct {
	*responses.BaseResponse
	RequestId              string `json:"RequestId" xml:"RequestId"`
	SubscriptionInstanceId string `json:"SubscriptionInstanceId" xml:"SubscriptionInstanceId"`
	ErrCode                string `json:"ErrCode" xml:"ErrCode"`
	Success                string `json:"Success" xml:"Success"`
	ErrMessage             string `json:"ErrMessage" xml:"ErrMessage"`
}

// CreateCreateSubscriptionInstanceRequest creates a request to invoke CreateSubscriptionInstance API
func CreateCreateSubscriptionInstanceRequest() (request *CreateSubscriptionInstanceRequest) {
	request = &CreateSubscriptionInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "CreateSubscriptionInstance", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSubscriptionInstanceResponse creates a response to parse from CreateSubscriptionInstance response
func CreateCreateSubscriptionInstanceResponse() (response *CreateSubscriptionInstanceResponse) {
	response = &CreateSubscriptionInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
