package polardb

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

// CreateDBClusterEndpoint invokes the polardb.CreateDBClusterEndpoint API synchronously
func (client *Client) CreateDBClusterEndpoint(request *CreateDBClusterEndpointRequest) (response *CreateDBClusterEndpointResponse, err error) {
	response = CreateCreateDBClusterEndpointResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDBClusterEndpointWithChan invokes the polardb.CreateDBClusterEndpoint API asynchronously
func (client *Client) CreateDBClusterEndpointWithChan(request *CreateDBClusterEndpointRequest) (<-chan *CreateDBClusterEndpointResponse, <-chan error) {
	responseChan := make(chan *CreateDBClusterEndpointResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDBClusterEndpoint(request)
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

// CreateDBClusterEndpointWithCallback invokes the polardb.CreateDBClusterEndpoint API asynchronously
func (client *Client) CreateDBClusterEndpointWithCallback(request *CreateDBClusterEndpointRequest, callback func(response *CreateDBClusterEndpointResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDBClusterEndpointResponse
		var err error
		defer close(result)
		response, err = client.CreateDBClusterEndpoint(request)
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

// CreateDBClusterEndpointRequest is the request struct for api CreateDBClusterEndpoint
type CreateDBClusterEndpointRequest struct {
	*requests.RpcRequest
	AutoAddNewNodes       string           `position:"Query" name:"AutoAddNewNodes"`
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken           string           `position:"Query" name:"ClientToken"`
	ReadWriteMode         string           `position:"Query" name:"ReadWriteMode"`
	EndpointType          string           `position:"Query" name:"EndpointType"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId           string           `position:"Query" name:"DBClusterId"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	EndpointConfig        string           `position:"Query" name:"EndpointConfig"`
	DBEndpointDescription string           `position:"Query" name:"DBEndpointDescription"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	Nodes                 string           `position:"Query" name:"Nodes"`
}

// CreateDBClusterEndpointResponse is the response struct for api CreateDBClusterEndpoint
type CreateDBClusterEndpointResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateDBClusterEndpointRequest creates a request to invoke CreateDBClusterEndpoint API
func CreateCreateDBClusterEndpointRequest() (request *CreateDBClusterEndpointRequest) {
	request = &CreateDBClusterEndpointRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "CreateDBClusterEndpoint", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDBClusterEndpointResponse creates a response to parse from CreateDBClusterEndpoint response
func CreateCreateDBClusterEndpointResponse() (response *CreateDBClusterEndpointResponse) {
	response = &CreateDBClusterEndpointResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
