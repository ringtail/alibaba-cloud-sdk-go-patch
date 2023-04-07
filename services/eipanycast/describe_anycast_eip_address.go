package eipanycast

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

// DescribeAnycastEipAddress invokes the eipanycast.DescribeAnycastEipAddress API synchronously
func (client *Client) DescribeAnycastEipAddress(request *DescribeAnycastEipAddressRequest) (response *DescribeAnycastEipAddressResponse, err error) {
	response = CreateDescribeAnycastEipAddressResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAnycastEipAddressWithChan invokes the eipanycast.DescribeAnycastEipAddress API asynchronously
func (client *Client) DescribeAnycastEipAddressWithChan(request *DescribeAnycastEipAddressRequest) (<-chan *DescribeAnycastEipAddressResponse, <-chan error) {
	responseChan := make(chan *DescribeAnycastEipAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAnycastEipAddress(request)
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

// DescribeAnycastEipAddressWithCallback invokes the eipanycast.DescribeAnycastEipAddress API asynchronously
func (client *Client) DescribeAnycastEipAddressWithCallback(request *DescribeAnycastEipAddressRequest, callback func(response *DescribeAnycastEipAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAnycastEipAddressResponse
		var err error
		defer close(result)
		response, err = client.DescribeAnycastEipAddress(request)
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

// DescribeAnycastEipAddressRequest is the request struct for api DescribeAnycastEipAddress
type DescribeAnycastEipAddressRequest struct {
	*requests.RpcRequest
	Ip             string `position:"Query" name:"Ip"`
	AnycastId      string `position:"Query" name:"AnycastId"`
	BindInstanceId string `position:"Query" name:"BindInstanceId"`
}

// DescribeAnycastEipAddressResponse is the response struct for api DescribeAnycastEipAddress
type DescribeAnycastEipAddressResponse struct {
	*responses.BaseResponse
	Status                 string               `json:"Status" xml:"Status"`
	Description            string               `json:"Description" xml:"Description"`
	RequestId              string               `json:"RequestId" xml:"RequestId"`
	InstanceChargeType     string               `json:"InstanceChargeType" xml:"InstanceChargeType"`
	CreateTime             string               `json:"CreateTime" xml:"CreateTime"`
	BusinessStatus         string               `json:"BusinessStatus" xml:"BusinessStatus"`
	InternetChargeType     string               `json:"InternetChargeType" xml:"InternetChargeType"`
	Name                   string               `json:"Name" xml:"Name"`
	AnycastId              string               `json:"AnycastId" xml:"AnycastId"`
	ServiceLocation        string               `json:"ServiceLocation" xml:"ServiceLocation"`
	Bandwidth              int                  `json:"Bandwidth" xml:"Bandwidth"`
	IpAddress              string               `json:"IpAddress" xml:"IpAddress"`
	Bid                    string               `json:"Bid" xml:"Bid"`
	AliUid                 int64                `json:"AliUid" xml:"AliUid"`
	AnycastEipBindInfoList []AnycastEipBindInfo `json:"AnycastEipBindInfoList" xml:"AnycastEipBindInfoList"`
	Tags                   []Tag                `json:"Tags" xml:"Tags"`
}

// CreateDescribeAnycastEipAddressRequest creates a request to invoke DescribeAnycastEipAddress API
func CreateDescribeAnycastEipAddressRequest() (request *DescribeAnycastEipAddressRequest) {
	request = &DescribeAnycastEipAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Eipanycast", "2020-03-09", "DescribeAnycastEipAddress", "eipanycast", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAnycastEipAddressResponse creates a response to parse from DescribeAnycastEipAddress response
func CreateDescribeAnycastEipAddressResponse() (response *DescribeAnycastEipAddressResponse) {
	response = &DescribeAnycastEipAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
