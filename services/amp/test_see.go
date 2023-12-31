package amp

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

// TestSee invokes the amp.TestSee API synchronously
func (client *Client) TestSee(request *TestSeeRequest) (response *TestSeeResponse, err error) {
	response = CreateTestSeeResponse()
	err = client.DoAction(request, response)
	return
}

// TestSeeWithChan invokes the amp.TestSee API asynchronously
func (client *Client) TestSeeWithChan(request *TestSeeRequest) (<-chan *TestSeeResponse, <-chan error) {
	responseChan := make(chan *TestSeeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TestSee(request)
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

// TestSeeWithCallback invokes the amp.TestSee API asynchronously
func (client *Client) TestSeeWithCallback(request *TestSeeRequest, callback func(response *TestSeeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TestSeeResponse
		var err error
		defer close(result)
		response, err = client.TestSee(request)
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

// TestSeeRequest is the request struct for api TestSee
type TestSeeRequest struct {
	*requests.RoaRequest
}

// TestSeeResponse is the response struct for api TestSee
type TestSeeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateTestSeeRequest creates a request to invoke TestSee API
func CreateTestSeeRequest() (request *TestSeeRequest) {
	request = &TestSeeRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("amp", "2020-07-08", "TestSee", "/wo/eii//ee", "ServiceCode", "openAPI")
	request.Method = requests.POST
	return
}

// CreateTestSeeResponse creates a response to parse from TestSee response
func CreateTestSeeResponse() (response *TestSeeResponse) {
	response = &TestSeeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
