package ltl

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

// UploadMPCoSPhaseTextInfo invokes the ltl.UploadMPCoSPhaseTextInfo API synchronously
func (client *Client) UploadMPCoSPhaseTextInfo(request *UploadMPCoSPhaseTextInfoRequest) (response *UploadMPCoSPhaseTextInfoResponse, err error) {
	response = CreateUploadMPCoSPhaseTextInfoResponse()
	err = client.DoAction(request, response)
	return
}

// UploadMPCoSPhaseTextInfoWithChan invokes the ltl.UploadMPCoSPhaseTextInfo API asynchronously
func (client *Client) UploadMPCoSPhaseTextInfoWithChan(request *UploadMPCoSPhaseTextInfoRequest) (<-chan *UploadMPCoSPhaseTextInfoResponse, <-chan error) {
	responseChan := make(chan *UploadMPCoSPhaseTextInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadMPCoSPhaseTextInfo(request)
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

// UploadMPCoSPhaseTextInfoWithCallback invokes the ltl.UploadMPCoSPhaseTextInfo API asynchronously
func (client *Client) UploadMPCoSPhaseTextInfoWithCallback(request *UploadMPCoSPhaseTextInfoRequest, callback func(response *UploadMPCoSPhaseTextInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadMPCoSPhaseTextInfoResponse
		var err error
		defer close(result)
		response, err = client.UploadMPCoSPhaseTextInfo(request)
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

// UploadMPCoSPhaseTextInfoRequest is the request struct for api UploadMPCoSPhaseTextInfo
type UploadMPCoSPhaseTextInfoRequest struct {
	*requests.RpcRequest
	PhaseData       string                 `position:"Query" name:"PhaseData"`
	PhaseId         string                 `position:"Query" name:"PhaseId"`
	ApiVersion      string                 `position:"Query" name:"ApiVersion"`
	BizChainId      string                 `position:"Query" name:"BizChainId"`
	DataKey         string                 `position:"Query" name:"DataKey"`
	DataSeq         string                 `position:"Query" name:"DataSeq"`
	PhaseGroupId    string                 `position:"Query" name:"PhaseGroupId"`
	RelatedDataList map[string]interface{} `position:"Query" name:"RelatedDataList"`
}

// UploadMPCoSPhaseTextInfoResponse is the response struct for api UploadMPCoSPhaseTextInfo
type UploadMPCoSPhaseTextInfoResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUploadMPCoSPhaseTextInfoRequest creates a request to invoke UploadMPCoSPhaseTextInfo API
func CreateUploadMPCoSPhaseTextInfoRequest() (request *UploadMPCoSPhaseTextInfoRequest) {
	request = &UploadMPCoSPhaseTextInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ltl", "2019-05-10", "UploadMPCoSPhaseTextInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateUploadMPCoSPhaseTextInfoResponse creates a response to parse from UploadMPCoSPhaseTextInfo response
func CreateUploadMPCoSPhaseTextInfoResponse() (response *UploadMPCoSPhaseTextInfoResponse) {
	response = &UploadMPCoSPhaseTextInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
