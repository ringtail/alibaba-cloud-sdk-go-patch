package resourcemanager

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

// DeclineHandshake invokes the resourcemanager.DeclineHandshake API synchronously
func (client *Client) DeclineHandshake(request *DeclineHandshakeRequest) (response *DeclineHandshakeResponse, err error) {
	response = CreateDeclineHandshakeResponse()
	err = client.DoAction(request, response)
	return
}

// DeclineHandshakeWithChan invokes the resourcemanager.DeclineHandshake API asynchronously
func (client *Client) DeclineHandshakeWithChan(request *DeclineHandshakeRequest) (<-chan *DeclineHandshakeResponse, <-chan error) {
	responseChan := make(chan *DeclineHandshakeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeclineHandshake(request)
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

// DeclineHandshakeWithCallback invokes the resourcemanager.DeclineHandshake API asynchronously
func (client *Client) DeclineHandshakeWithCallback(request *DeclineHandshakeRequest, callback func(response *DeclineHandshakeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeclineHandshakeResponse
		var err error
		defer close(result)
		response, err = client.DeclineHandshake(request)
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

// DeclineHandshakeRequest is the request struct for api DeclineHandshake
type DeclineHandshakeRequest struct {
	*requests.RpcRequest
	HandshakeId string `position:"Query" name:"HandshakeId"`
}

// DeclineHandshakeResponse is the response struct for api DeclineHandshake
type DeclineHandshakeResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Handshake Handshake `json:"Handshake" xml:"Handshake"`
}

// CreateDeclineHandshakeRequest creates a request to invoke DeclineHandshake API
func CreateDeclineHandshakeRequest() (request *DeclineHandshakeRequest) {
	request = &DeclineHandshakeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "DeclineHandshake", "resourcemanager", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeclineHandshakeResponse creates a response to parse from DeclineHandshake response
func CreateDeclineHandshakeResponse() (response *DeclineHandshakeResponse) {
	response = &DeclineHandshakeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
