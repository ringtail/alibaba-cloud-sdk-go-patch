package opensearch

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

// DescribeABTestScene invokes the opensearch.DescribeABTestScene API synchronously
func (client *Client) DescribeABTestScene(request *DescribeABTestSceneRequest) (response *DescribeABTestSceneResponse, err error) {
	response = CreateDescribeABTestSceneResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeABTestSceneWithChan invokes the opensearch.DescribeABTestScene API asynchronously
func (client *Client) DescribeABTestSceneWithChan(request *DescribeABTestSceneRequest) (<-chan *DescribeABTestSceneResponse, <-chan error) {
	responseChan := make(chan *DescribeABTestSceneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeABTestScene(request)
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

// DescribeABTestSceneWithCallback invokes the opensearch.DescribeABTestScene API asynchronously
func (client *Client) DescribeABTestSceneWithCallback(request *DescribeABTestSceneRequest, callback func(response *DescribeABTestSceneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeABTestSceneResponse
		var err error
		defer close(result)
		response, err = client.DescribeABTestScene(request)
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

// DescribeABTestSceneRequest is the request struct for api DescribeABTestScene
type DescribeABTestSceneRequest struct {
	*requests.RoaRequest
	SceneId          requests.Integer `position:"Path" name:"sceneId"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
}

// DescribeABTestSceneResponse is the response struct for api DescribeABTestScene
type DescribeABTestSceneResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    Result `json:"result" xml:"result"`
}

// CreateDescribeABTestSceneRequest creates a request to invoke DescribeABTestScene API
func CreateDescribeABTestSceneRequest() (request *DescribeABTestSceneRequest) {
	request = &DescribeABTestSceneRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "DescribeABTestScene", "/v4/openapi/app-groups/[appGroupIdentity]/scenes/[sceneId]", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeABTestSceneResponse creates a response to parse from DescribeABTestScene response
func CreateDescribeABTestSceneResponse() (response *DescribeABTestSceneResponse) {
	response = &DescribeABTestSceneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
