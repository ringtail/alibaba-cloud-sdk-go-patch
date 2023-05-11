package linkvisual

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

// QueryRecordUrl invokes the linkvisual.QueryRecordUrl API synchronously
func (client *Client) QueryRecordUrl(request *QueryRecordUrlRequest) (response *QueryRecordUrlResponse, err error) {
	response = CreateQueryRecordUrlResponse()
	err = client.DoAction(request, response)
	return
}

// QueryRecordUrlWithChan invokes the linkvisual.QueryRecordUrl API asynchronously
func (client *Client) QueryRecordUrlWithChan(request *QueryRecordUrlRequest) (<-chan *QueryRecordUrlResponse, <-chan error) {
	responseChan := make(chan *QueryRecordUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryRecordUrl(request)
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

// QueryRecordUrlWithCallback invokes the linkvisual.QueryRecordUrl API asynchronously
func (client *Client) QueryRecordUrlWithCallback(request *QueryRecordUrlRequest, callback func(response *QueryRecordUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryRecordUrlResponse
		var err error
		defer close(result)
		response, err = client.QueryRecordUrl(request)
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

// QueryRecordUrlRequest is the request struct for api QueryRecordUrl
type QueryRecordUrlRequest struct {
	*requests.RpcRequest
	IotId            string           `position:"Query" name:"IotId"`
	IotInstanceId    string           `position:"Query" name:"IotInstanceId"`
	ProductKey       string           `position:"Query" name:"ProductKey"`
	FileName         string           `position:"Query" name:"FileName"`
	SeekTime         requests.Integer `position:"Query" name:"SeekTime"`
	ApiProduct       string           `position:"Body" name:"ApiProduct"`
	ApiRevision      string           `position:"Body" name:"ApiRevision"`
	DeviceName       string           `position:"Query" name:"DeviceName"`
	UrlValidDuration requests.Integer `position:"Query" name:"UrlValidDuration"`
}

// QueryRecordUrlResponse is the response struct for api QueryRecordUrl
type QueryRecordUrlResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	Data         string `json:"Data" xml:"Data"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateQueryRecordUrlRequest creates a request to invoke QueryRecordUrl API
func CreateQueryRecordUrlRequest() (request *QueryRecordUrlRequest) {
	request = &QueryRecordUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "QueryRecordUrl", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryRecordUrlResponse creates a response to parse from QueryRecordUrl response
func CreateQueryRecordUrlResponse() (response *QueryRecordUrlResponse) {
	response = &QueryRecordUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
