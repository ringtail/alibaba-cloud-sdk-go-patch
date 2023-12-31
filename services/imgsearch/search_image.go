package imgsearch

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

// SearchImage invokes the imgsearch.SearchImage API synchronously
func (client *Client) SearchImage(request *SearchImageRequest) (response *SearchImageResponse, err error) {
	response = CreateSearchImageResponse()
	err = client.DoAction(request, response)
	return
}

// SearchImageWithChan invokes the imgsearch.SearchImage API asynchronously
func (client *Client) SearchImageWithChan(request *SearchImageRequest) (<-chan *SearchImageResponse, <-chan error) {
	responseChan := make(chan *SearchImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchImage(request)
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

// SearchImageWithCallback invokes the imgsearch.SearchImage API asynchronously
func (client *Client) SearchImageWithCallback(request *SearchImageRequest, callback func(response *SearchImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchImageResponse
		var err error
		defer close(result)
		response, err = client.SearchImage(request)
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

// SearchImageRequest is the request struct for api SearchImage
type SearchImageRequest struct {
	*requests.RpcRequest
	DbName   string           `position:"Body" name:"DbName"`
	ImageUrl string           `position:"Body" name:"ImageUrl"`
	Limit    requests.Integer `position:"Body" name:"Limit"`
}

// SearchImageResponse is the response struct for api SearchImage
type SearchImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSearchImageRequest creates a request to invoke SearchImage API
func CreateSearchImageRequest() (request *SearchImageRequest) {
	request = &SearchImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imgsearch", "2020-03-20", "SearchImage", "", "")
	request.Method = requests.POST
	return
}

// CreateSearchImageResponse creates a response to parse from SearchImage response
func CreateSearchImageResponse() (response *SearchImageResponse) {
	response = &SearchImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
