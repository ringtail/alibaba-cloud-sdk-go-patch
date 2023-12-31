package mse

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

// UpdateGatewaySpec invokes the mse.UpdateGatewaySpec API synchronously
func (client *Client) UpdateGatewaySpec(request *UpdateGatewaySpecRequest) (response *UpdateGatewaySpecResponse, err error) {
	response = CreateUpdateGatewaySpecResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateGatewaySpecWithChan invokes the mse.UpdateGatewaySpec API asynchronously
func (client *Client) UpdateGatewaySpecWithChan(request *UpdateGatewaySpecRequest) (<-chan *UpdateGatewaySpecResponse, <-chan error) {
	responseChan := make(chan *UpdateGatewaySpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateGatewaySpec(request)
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

// UpdateGatewaySpecWithCallback invokes the mse.UpdateGatewaySpec API asynchronously
func (client *Client) UpdateGatewaySpecWithCallback(request *UpdateGatewaySpecRequest, callback func(response *UpdateGatewaySpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateGatewaySpecResponse
		var err error
		defer close(result)
		response, err = client.UpdateGatewaySpec(request)
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

// UpdateGatewaySpecRequest is the request struct for api UpdateGatewaySpec
type UpdateGatewaySpecRequest struct {
	*requests.RpcRequest
	MseSessionId    string           `position:"Query" name:"MseSessionId"`
	GatewayUniqueId string           `position:"Query" name:"GatewayUniqueId"`
	Replica         requests.Integer `position:"Query" name:"Replica"`
	Spec            string           `position:"Query" name:"Spec"`
	AcceptLanguage  string           `position:"Query" name:"AcceptLanguage"`
}

// UpdateGatewaySpecResponse is the response struct for api UpdateGatewaySpec
type UpdateGatewaySpecResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           string `json:"Data" xml:"Data"`
}

// CreateUpdateGatewaySpecRequest creates a request to invoke UpdateGatewaySpec API
func CreateUpdateGatewaySpecRequest() (request *UpdateGatewaySpecRequest) {
	request = &UpdateGatewaySpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateGatewaySpec", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateGatewaySpecResponse creates a response to parse from UpdateGatewaySpec response
func CreateUpdateGatewaySpecResponse() (response *UpdateGatewaySpecResponse) {
	response = &UpdateGatewaySpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
