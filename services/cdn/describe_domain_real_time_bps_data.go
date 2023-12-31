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

// DescribeDomainRealTimeBpsData invokes the cdn.DescribeDomainRealTimeBpsData API synchronously
func (client *Client) DescribeDomainRealTimeBpsData(request *DescribeDomainRealTimeBpsDataRequest) (response *DescribeDomainRealTimeBpsDataResponse, err error) {
	response = CreateDescribeDomainRealTimeBpsDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainRealTimeBpsDataWithChan invokes the cdn.DescribeDomainRealTimeBpsData API asynchronously
func (client *Client) DescribeDomainRealTimeBpsDataWithChan(request *DescribeDomainRealTimeBpsDataRequest) (<-chan *DescribeDomainRealTimeBpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainRealTimeBpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainRealTimeBpsData(request)
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

// DescribeDomainRealTimeBpsDataWithCallback invokes the cdn.DescribeDomainRealTimeBpsData API asynchronously
func (client *Client) DescribeDomainRealTimeBpsDataWithCallback(request *DescribeDomainRealTimeBpsDataRequest, callback func(response *DescribeDomainRealTimeBpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainRealTimeBpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainRealTimeBpsData(request)
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

// DescribeDomainRealTimeBpsDataRequest is the request struct for api DescribeDomainRealTimeBpsData
type DescribeDomainRealTimeBpsDataRequest struct {
	*requests.RpcRequest
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
	StartTime      string `position:"Query" name:"StartTime"`
}

// DescribeDomainRealTimeBpsDataResponse is the response struct for api DescribeDomainRealTimeBpsData
type DescribeDomainRealTimeBpsDataResponse struct {
	*responses.BaseResponse
	RequestId string                              `json:"RequestId" xml:"RequestId"`
	Data      DataInDescribeDomainRealTimeBpsData `json:"Data" xml:"Data"`
}

// CreateDescribeDomainRealTimeBpsDataRequest creates a request to invoke DescribeDomainRealTimeBpsData API
func CreateDescribeDomainRealTimeBpsDataRequest() (request *DescribeDomainRealTimeBpsDataRequest) {
	request = &DescribeDomainRealTimeBpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainRealTimeBpsData", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeDomainRealTimeBpsDataResponse creates a response to parse from DescribeDomainRealTimeBpsData response
func CreateDescribeDomainRealTimeBpsDataResponse() (response *DescribeDomainRealTimeBpsDataResponse) {
	response = &DescribeDomainRealTimeBpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
