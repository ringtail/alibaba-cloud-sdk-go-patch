package sae

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

// UpdateJob invokes the sae.UpdateJob API synchronously
func (client *Client) UpdateJob(request *UpdateJobRequest) (response *UpdateJobResponse, err error) {
	response = CreateUpdateJobResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateJobWithChan invokes the sae.UpdateJob API asynchronously
func (client *Client) UpdateJobWithChan(request *UpdateJobRequest) (<-chan *UpdateJobResponse, <-chan error) {
	responseChan := make(chan *UpdateJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateJob(request)
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

// UpdateJobWithCallback invokes the sae.UpdateJob API asynchronously
func (client *Client) UpdateJobWithCallback(request *UpdateJobRequest, callback func(response *UpdateJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateJobResponse
		var err error
		defer close(result)
		response, err = client.UpdateJob(request)
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

// UpdateJobRequest is the request struct for api UpdateJob
type UpdateJobRequest struct {
	*requests.RoaRequest
	NasId                            string           `position:"Query" name:"NasId"`
	JarStartArgs                     string           `position:"Query" name:"JarStartArgs"`
	ConcurrencyPolicy                string           `position:"Query" name:"ConcurrencyPolicy"`
	TriggerConfig                    string           `position:"Query" name:"TriggerConfig"`
	OssAkSecret                      string           `position:"Body" name:"OssAkSecret"`
	MountHost                        string           `position:"Query" name:"MountHost"`
	BatchWaitTime                    requests.Integer `position:"Query" name:"BatchWaitTime"`
	Envs                             string           `position:"Query" name:"Envs"`
	PhpPECLExtensions                string           `position:"Body" name:"PhpPECLExtensions"`
	PhpArmsConfigLocation            string           `position:"Query" name:"PhpArmsConfigLocation"`
	ProgrammingLanguage              string           `position:"Query" name:"ProgrammingLanguage"`
	CustomHostAlias                  string           `position:"Query" name:"CustomHostAlias"`
	JarStartOptions                  string           `position:"Query" name:"JarStartOptions"`
	Slice                            requests.Boolean `position:"Query" name:"Slice"`
	ConfigMapMountDesc               string           `position:"Body" name:"ConfigMapMountDesc"`
	OssMountDescs                    string           `position:"Body" name:"OssMountDescs"`
	ImagePullSecrets                 string           `position:"Query" name:"ImagePullSecrets"`
	PreStop                          string           `position:"Query" name:"PreStop"`
	Python                           string           `position:"Query" name:"Python"`
	BackoffLimit                     requests.Integer `position:"Query" name:"BackoffLimit"`
	UpdateStrategy                   string           `position:"Query" name:"UpdateStrategy"`
	ChangeOrderDesc                  string           `position:"Query" name:"ChangeOrderDesc"`
	AutoEnableApplicationScalingRule requests.Boolean `position:"Query" name:"AutoEnableApplicationScalingRule"`
	PostStart                        string           `position:"Query" name:"PostStart"`
	PhpExtensions                    string           `position:"Body" name:"PhpExtensions"`
	AssociateEip                     requests.Boolean `position:"Body" name:"AssociateEip"`
	WebContainer                     string           `position:"Query" name:"WebContainer"`
	EnableAhas                       string           `position:"Query" name:"EnableAhas"`
	SlsConfigs                       string           `position:"Query" name:"SlsConfigs"`
	CommandArgs                      string           `position:"Query" name:"CommandArgs"`
	AcrAssumeRoleArn                 string           `position:"Query" name:"AcrAssumeRoleArn"`
	Readiness                        string           `position:"Query" name:"Readiness"`
	Timezone                         string           `position:"Query" name:"Timezone"`
	OssAkId                          string           `position:"Body" name:"OssAkId"`
	Liveness                         string           `position:"Query" name:"Liveness"`
	PackageVersion                   string           `position:"Query" name:"PackageVersion"`
	TomcatConfig                     string           `position:"Query" name:"TomcatConfig"`
	Timeout                          requests.Integer `position:"Query" name:"Timeout"`
	WarStartOptions                  string           `position:"Query" name:"WarStartOptions"`
	PackageRuntimeCustomBuild        string           `position:"Body" name:"PackageRuntimeCustomBuild"`
	EdasContainerVersion             string           `position:"Query" name:"EdasContainerVersion"`
	PackageUrl                       string           `position:"Query" name:"PackageUrl"`
	TerminationGracePeriodSeconds    requests.Integer `position:"Query" name:"TerminationGracePeriodSeconds"`
	PhpConfig                        string           `position:"Body" name:"PhpConfig"`
	SliceEnvs                        string           `position:"Query" name:"SliceEnvs"`
	EnableImageAccl                  requests.Boolean `position:"Body" name:"EnableImageAccl"`
	EnableGreyTagRoute               requests.Boolean `position:"Query" name:"EnableGreyTagRoute"`
	Replicas                         string           `position:"Query" name:"Replicas"`
	Command                          string           `position:"Query" name:"Command"`
	MountDesc                        string           `position:"Query" name:"MountDesc"`
	Jdk                              string           `position:"Query" name:"Jdk"`
	MinReadyInstances                requests.Integer `position:"Query" name:"MinReadyInstances"`
	AcrInstanceId                    string           `position:"Body" name:"AcrInstanceId"`
	AppId                            string           `position:"Query" name:"AppId"`
	ImageUrl                         string           `position:"Query" name:"ImageUrl"`
	Php                              string           `position:"Body" name:"Php"`
	RefAppId                         string           `position:"Query" name:"RefAppId"`
	PythonModules                    string           `position:"Query" name:"PythonModules"`
	PhpConfigLocation                string           `position:"Query" name:"PhpConfigLocation"`
}

// UpdateJobResponse is the response struct for api UpdateJob
type UpdateJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateUpdateJobRequest creates a request to invoke UpdateJob API
func CreateUpdateJobRequest() (request *UpdateJobRequest) {
	request = &UpdateJobRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "UpdateJob", "/pop/v1/sam/job/updateJob", "serverless", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateJobResponse creates a response to parse from UpdateJob response
func CreateUpdateJobResponse() (response *UpdateJobResponse) {
	response = &UpdateJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
