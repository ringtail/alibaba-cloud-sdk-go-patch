package viapi_regen

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

// DownloadDataset invokes the viapi_regen.DownloadDataset API synchronously
func (client *Client) DownloadDataset(request *DownloadDatasetRequest) (response *DownloadDatasetResponse, err error) {
	response = CreateDownloadDatasetResponse()
	err = client.DoAction(request, response)
	return
}

// DownloadDatasetWithChan invokes the viapi_regen.DownloadDataset API asynchronously
func (client *Client) DownloadDatasetWithChan(request *DownloadDatasetRequest) (<-chan *DownloadDatasetResponse, <-chan error) {
	responseChan := make(chan *DownloadDatasetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DownloadDataset(request)
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

// DownloadDatasetWithCallback invokes the viapi_regen.DownloadDataset API asynchronously
func (client *Client) DownloadDatasetWithCallback(request *DownloadDatasetRequest, callback func(response *DownloadDatasetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DownloadDatasetResponse
		var err error
		defer close(result)
		response, err = client.DownloadDataset(request)
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

// DownloadDatasetRequest is the request struct for api DownloadDataset
type DownloadDatasetRequest struct {
	*requests.RpcRequest
	DatasetId requests.Integer `position:"Body" name:"DatasetId"`
}

// DownloadDatasetResponse is the response struct for api DownloadDataset
type DownloadDatasetResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDownloadDatasetRequest creates a request to invoke DownloadDataset API
func CreateDownloadDatasetRequest() (request *DownloadDatasetRequest) {
	request = &DownloadDatasetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi-regen", "2021-11-19", "DownloadDataset", "selflearning", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDownloadDatasetResponse creates a response to parse from DownloadDataset response
func CreateDownloadDatasetResponse() (response *DownloadDatasetResponse) {
	response = &DownloadDatasetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
