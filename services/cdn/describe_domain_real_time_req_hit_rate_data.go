package cdn

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

// DescribeDomainRealTimeReqHitRateData invokes the cdn.DescribeDomainRealTimeReqHitRateData API synchronously
func (client *Client) DescribeDomainRealTimeReqHitRateData(request *DescribeDomainRealTimeReqHitRateDataRequest) (response *DescribeDomainRealTimeReqHitRateDataResponse, err error) {
	response = CreateDescribeDomainRealTimeReqHitRateDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainRealTimeReqHitRateDataWithChan invokes the cdn.DescribeDomainRealTimeReqHitRateData API asynchronously
func (client *Client) DescribeDomainRealTimeReqHitRateDataWithChan(request *DescribeDomainRealTimeReqHitRateDataRequest) (<-chan *DescribeDomainRealTimeReqHitRateDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainRealTimeReqHitRateDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainRealTimeReqHitRateData(request)
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

// DescribeDomainRealTimeReqHitRateDataWithCallback invokes the cdn.DescribeDomainRealTimeReqHitRateData API asynchronously
func (client *Client) DescribeDomainRealTimeReqHitRateDataWithCallback(request *DescribeDomainRealTimeReqHitRateDataRequest, callback func(response *DescribeDomainRealTimeReqHitRateDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainRealTimeReqHitRateDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainRealTimeReqHitRateData(request)
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

// DescribeDomainRealTimeReqHitRateDataRequest is the request struct for api DescribeDomainRealTimeReqHitRateData
type DescribeDomainRealTimeReqHitRateDataRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
}

// DescribeDomainRealTimeReqHitRateDataResponse is the response struct for api DescribeDomainRealTimeReqHitRateData
type DescribeDomainRealTimeReqHitRateDataResponse struct {
	*responses.BaseResponse
	RequestId string                                     `json:"RequestId" xml:"RequestId"`
	Data      DataInDescribeDomainRealTimeReqHitRateData `json:"Data" xml:"Data"`
}

// CreateDescribeDomainRealTimeReqHitRateDataRequest creates a request to invoke DescribeDomainRealTimeReqHitRateData API
func CreateDescribeDomainRealTimeReqHitRateDataRequest() (request *DescribeDomainRealTimeReqHitRateDataRequest) {
	request = &DescribeDomainRealTimeReqHitRateDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainRealTimeReqHitRateData", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeDomainRealTimeReqHitRateDataResponse creates a response to parse from DescribeDomainRealTimeReqHitRateData response
func CreateDescribeDomainRealTimeReqHitRateDataResponse() (response *DescribeDomainRealTimeReqHitRateDataResponse) {
	response = &DescribeDomainRealTimeReqHitRateDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
