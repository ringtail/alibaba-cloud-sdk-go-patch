package slb

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

// DescribeDomainExtensionAttribute invokes the slb.DescribeDomainExtensionAttribute API synchronously
func (client *Client) DescribeDomainExtensionAttribute(request *DescribeDomainExtensionAttributeRequest) (response *DescribeDomainExtensionAttributeResponse, err error) {
	response = CreateDescribeDomainExtensionAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainExtensionAttributeWithChan invokes the slb.DescribeDomainExtensionAttribute API asynchronously
func (client *Client) DescribeDomainExtensionAttributeWithChan(request *DescribeDomainExtensionAttributeRequest) (<-chan *DescribeDomainExtensionAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainExtensionAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainExtensionAttribute(request)
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

// DescribeDomainExtensionAttributeWithCallback invokes the slb.DescribeDomainExtensionAttribute API asynchronously
func (client *Client) DescribeDomainExtensionAttributeWithCallback(request *DescribeDomainExtensionAttributeRequest, callback func(response *DescribeDomainExtensionAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainExtensionAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainExtensionAttribute(request)
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

// DescribeDomainExtensionAttributeRequest is the request struct for api DescribeDomainExtensionAttribute
type DescribeDomainExtensionAttributeRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DomainExtensionId    string           `position:"Query" name:"DomainExtensionId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tags                 string           `position:"Query" name:"Tags"`
}

// DescribeDomainExtensionAttributeResponse is the response struct for api DescribeDomainExtensionAttribute
type DescribeDomainExtensionAttributeResponse struct {
	*responses.BaseResponse
	Domain              string                                               `json:"Domain" xml:"Domain"`
	RequestId           string                                               `json:"RequestId" xml:"RequestId"`
	LoadBalancerId      string                                               `json:"LoadBalancerId" xml:"LoadBalancerId"`
	ListenerPort        int                                                  `json:"ListenerPort" xml:"ListenerPort"`
	ServerCertificateId string                                               `json:"ServerCertificateId" xml:"ServerCertificateId"`
	DomainExtensionId   string                                               `json:"DomainExtensionId" xml:"DomainExtensionId"`
	Certificates        CertificatesInDescribeDomainExtensionAttribute       `json:"Certificates" xml:"Certificates"`
	ServerCertificates  ServerCertificatesInDescribeDomainExtensionAttribute `json:"ServerCertificates" xml:"ServerCertificates"`
}

// CreateDescribeDomainExtensionAttributeRequest creates a request to invoke DescribeDomainExtensionAttribute API
func CreateDescribeDomainExtensionAttributeRequest() (request *DescribeDomainExtensionAttributeRequest) {
	request = &DescribeDomainExtensionAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeDomainExtensionAttribute", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainExtensionAttributeResponse creates a response to parse from DescribeDomainExtensionAttribute response
func CreateDescribeDomainExtensionAttributeResponse() (response *DescribeDomainExtensionAttributeResponse) {
	response = &DescribeDomainExtensionAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
