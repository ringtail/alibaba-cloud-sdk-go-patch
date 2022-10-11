package aiworkspace

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

// RemoveImageLabels invokes the aiworkspace.RemoveImageLabels API synchronously
func (client *Client) RemoveImageLabels(request *RemoveImageLabelsRequest) (response *RemoveImageLabelsResponse, err error) {
	response = CreateRemoveImageLabelsResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveImageLabelsWithChan invokes the aiworkspace.RemoveImageLabels API asynchronously
func (client *Client) RemoveImageLabelsWithChan(request *RemoveImageLabelsRequest) (<-chan *RemoveImageLabelsResponse, <-chan error) {
	responseChan := make(chan *RemoveImageLabelsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveImageLabels(request)
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

// RemoveImageLabelsWithCallback invokes the aiworkspace.RemoveImageLabels API asynchronously
func (client *Client) RemoveImageLabelsWithCallback(request *RemoveImageLabelsRequest, callback func(response *RemoveImageLabelsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveImageLabelsResponse
		var err error
		defer close(result)
		response, err = client.RemoveImageLabels(request)
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

// RemoveImageLabelsRequest is the request struct for api RemoveImageLabels
type RemoveImageLabelsRequest struct {
	*requests.RoaRequest
	ImageId   string `position:"Path" name:"ImageId"`
	LabelKeys string `position:"Path" name:"LabelKeys"`
}

// RemoveImageLabelsResponse is the response struct for api RemoveImageLabels
type RemoveImageLabelsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRemoveImageLabelsRequest creates a request to invoke RemoveImageLabels API
func CreateRemoveImageLabelsRequest() (request *RemoveImageLabelsRequest) {
	request = &RemoveImageLabelsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("AIWorkSpace", "2021-02-04", "RemoveImageLabels", "/api/v1/images/[ImageId]/labels/[LabelKeys]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateRemoveImageLabelsResponse creates a response to parse from RemoveImageLabels response
func CreateRemoveImageLabelsResponse() (response *RemoveImageLabelsResponse) {
	response = &RemoveImageLabelsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
