package dcdn

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

// DescribeDcdnFullDomainsBlockIPConfig invokes the dcdn.DescribeDcdnFullDomainsBlockIPConfig API synchronously
func (client *Client) DescribeDcdnFullDomainsBlockIPConfig(request *DescribeDcdnFullDomainsBlockIPConfigRequest) (response *DescribeDcdnFullDomainsBlockIPConfigResponse, err error) {
	response = CreateDescribeDcdnFullDomainsBlockIPConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnFullDomainsBlockIPConfigWithChan invokes the dcdn.DescribeDcdnFullDomainsBlockIPConfig API asynchronously
func (client *Client) DescribeDcdnFullDomainsBlockIPConfigWithChan(request *DescribeDcdnFullDomainsBlockIPConfigRequest) (<-chan *DescribeDcdnFullDomainsBlockIPConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnFullDomainsBlockIPConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnFullDomainsBlockIPConfig(request)
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

// DescribeDcdnFullDomainsBlockIPConfigWithCallback invokes the dcdn.DescribeDcdnFullDomainsBlockIPConfig API asynchronously
func (client *Client) DescribeDcdnFullDomainsBlockIPConfigWithCallback(request *DescribeDcdnFullDomainsBlockIPConfigRequest, callback func(response *DescribeDcdnFullDomainsBlockIPConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnFullDomainsBlockIPConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnFullDomainsBlockIPConfig(request)
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

// DescribeDcdnFullDomainsBlockIPConfigRequest is the request struct for api DescribeDcdnFullDomainsBlockIPConfig
type DescribeDcdnFullDomainsBlockIPConfigRequest struct {
	*requests.RpcRequest
	IPList string `position:"Query" name:"IPList"`
}

// DescribeDcdnFullDomainsBlockIPConfigResponse is the response struct for api DescribeDcdnFullDomainsBlockIPConfig
type DescribeDcdnFullDomainsBlockIPConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Code      int    `json:"Code" xml:"Code"`
}

// CreateDescribeDcdnFullDomainsBlockIPConfigRequest creates a request to invoke DescribeDcdnFullDomainsBlockIPConfig API
func CreateDescribeDcdnFullDomainsBlockIPConfigRequest() (request *DescribeDcdnFullDomainsBlockIPConfigRequest) {
	request = &DescribeDcdnFullDomainsBlockIPConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnFullDomainsBlockIPConfig", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeDcdnFullDomainsBlockIPConfigResponse creates a response to parse from DescribeDcdnFullDomainsBlockIPConfig response
func CreateDescribeDcdnFullDomainsBlockIPConfigResponse() (response *DescribeDcdnFullDomainsBlockIPConfigResponse) {
	response = &DescribeDcdnFullDomainsBlockIPConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
