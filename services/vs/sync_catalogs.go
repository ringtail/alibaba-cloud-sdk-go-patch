package vs

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

// SyncCatalogs invokes the vs.SyncCatalogs API synchronously
func (client *Client) SyncCatalogs(request *SyncCatalogsRequest) (response *SyncCatalogsResponse, err error) {
	response = CreateSyncCatalogsResponse()
	err = client.DoAction(request, response)
	return
}

// SyncCatalogsWithChan invokes the vs.SyncCatalogs API asynchronously
func (client *Client) SyncCatalogsWithChan(request *SyncCatalogsRequest) (<-chan *SyncCatalogsResponse, <-chan error) {
	responseChan := make(chan *SyncCatalogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SyncCatalogs(request)
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

// SyncCatalogsWithCallback invokes the vs.SyncCatalogs API asynchronously
func (client *Client) SyncCatalogsWithCallback(request *SyncCatalogsRequest, callback func(response *SyncCatalogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SyncCatalogsResponse
		var err error
		defer close(result)
		response, err = client.SyncCatalogs(request)
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

// SyncCatalogsRequest is the request struct for api SyncCatalogs
type SyncCatalogsRequest struct {
	*requests.RpcRequest
	Id      string           `position:"Query" name:"Id"`
	ShowLog string           `position:"Query" name:"ShowLog"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// SyncCatalogsResponse is the response struct for api SyncCatalogs
type SyncCatalogsResponse struct {
	*responses.BaseResponse
	Id        string `json:"Id" xml:"Id"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSyncCatalogsRequest creates a request to invoke SyncCatalogs API
func CreateSyncCatalogsRequest() (request *SyncCatalogsRequest) {
	request = &SyncCatalogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "SyncCatalogs", "", "")
	request.Method = requests.POST
	return
}

// CreateSyncCatalogsResponse creates a response to parse from SyncCatalogs response
func CreateSyncCatalogsResponse() (response *SyncCatalogsResponse) {
	response = &SyncCatalogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
