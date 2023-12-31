package imageprocess

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

// DetectKneeKeypointXRay invokes the imageprocess.DetectKneeKeypointXRay API synchronously
func (client *Client) DetectKneeKeypointXRay(request *DetectKneeKeypointXRayRequest) (response *DetectKneeKeypointXRayResponse, err error) {
	response = CreateDetectKneeKeypointXRayResponse()
	err = client.DoAction(request, response)
	return
}

// DetectKneeKeypointXRayWithChan invokes the imageprocess.DetectKneeKeypointXRay API asynchronously
func (client *Client) DetectKneeKeypointXRayWithChan(request *DetectKneeKeypointXRayRequest) (<-chan *DetectKneeKeypointXRayResponse, <-chan error) {
	responseChan := make(chan *DetectKneeKeypointXRayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectKneeKeypointXRay(request)
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

// DetectKneeKeypointXRayWithCallback invokes the imageprocess.DetectKneeKeypointXRay API asynchronously
func (client *Client) DetectKneeKeypointXRayWithCallback(request *DetectKneeKeypointXRayRequest, callback func(response *DetectKneeKeypointXRayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectKneeKeypointXRayResponse
		var err error
		defer close(result)
		response, err = client.DetectKneeKeypointXRay(request)
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

// DetectKneeKeypointXRayRequest is the request struct for api DetectKneeKeypointXRay
type DetectKneeKeypointXRayRequest struct {
	*requests.RpcRequest
	OrgName    string `position:"Body" name:"OrgName"`
	TracerId   string `position:"Body" name:"TracerId"`
	DataFormat string `position:"Body" name:"DataFormat"`
	OrgId      string `position:"Body" name:"OrgId"`
	ImageUrl   string `position:"Body" name:"ImageUrl"`
}

// DetectKneeKeypointXRayResponse is the response struct for api DetectKneeKeypointXRay
type DetectKneeKeypointXRayResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDetectKneeKeypointXRayRequest creates a request to invoke DetectKneeKeypointXRay API
func CreateDetectKneeKeypointXRayRequest() (request *DetectKneeKeypointXRayRequest) {
	request = &DetectKneeKeypointXRayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imageprocess", "2020-03-20", "DetectKneeKeypointXRay", "imageprocess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetectKneeKeypointXRayResponse creates a response to parse from DetectKneeKeypointXRay response
func CreateDetectKneeKeypointXRayResponse() (response *DetectKneeKeypointXRayResponse) {
	response = &DetectKneeKeypointXRayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
