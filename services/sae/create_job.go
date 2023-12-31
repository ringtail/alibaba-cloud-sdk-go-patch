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

// CreateJob invokes the sae.CreateJob API synchronously
func (client *Client) CreateJob(request *CreateJobRequest) (response *CreateJobResponse, err error) {
	response = CreateCreateJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateJobWithChan invokes the sae.CreateJob API asynchronously
func (client *Client) CreateJobWithChan(request *CreateJobRequest) (<-chan *CreateJobResponse, <-chan error) {
	responseChan := make(chan *CreateJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateJob(request)
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

// CreateJobWithCallback invokes the sae.CreateJob API asynchronously
func (client *Client) CreateJobWithCallback(request *CreateJobRequest, callback func(response *CreateJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateJobResponse
		var err error
		defer close(result)
		response, err = client.CreateJob(request)
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

// CreateJobRequest is the request struct for api CreateJob
type CreateJobRequest struct {
	*requests.RoaRequest
	NasId                         string           `position:"Query" name:"NasId"`
	JarStartArgs                  string           `position:"Query" name:"JarStartArgs"`
	ConcurrencyPolicy             string           `position:"Query" name:"ConcurrencyPolicy"`
	TriggerConfig                 string           `position:"Query" name:"TriggerConfig"`
	OssAkSecret                   string           `position:"Body" name:"OssAkSecret"`
	MountHost                     string           `position:"Query" name:"MountHost"`
	AutoConfig                    requests.Boolean `position:"Query" name:"AutoConfig"`
	Envs                          string           `position:"Query" name:"Envs"`
	PhpPECLExtensions             string           `position:"Body" name:"PhpPECLExtensions"`
	PhpArmsConfigLocation         string           `position:"Query" name:"PhpArmsConfigLocation"`
	ProgrammingLanguage           string           `position:"Query" name:"ProgrammingLanguage"`
	CustomHostAlias               string           `position:"Query" name:"CustomHostAlias"`
	Deploy                        requests.Boolean `position:"Query" name:"Deploy"`
	JarStartOptions               string           `position:"Query" name:"JarStartOptions"`
	AppName                       string           `position:"Query" name:"AppName"`
	NamespaceId                   string           `position:"Query" name:"NamespaceId"`
	Slice                         requests.Boolean `position:"Query" name:"Slice"`
	ConfigMapMountDesc            string           `position:"Body" name:"ConfigMapMountDesc"`
	OssMountDescs                 string           `position:"Body" name:"OssMountDescs"`
	ImagePullSecrets              string           `position:"Query" name:"ImagePullSecrets"`
	PreStop                       string           `position:"Query" name:"PreStop"`
	Python                        string           `position:"Query" name:"Python"`
	Cpu                           requests.Integer `position:"Query" name:"Cpu"`
	BackoffLimit                  requests.Integer `position:"Query" name:"BackoffLimit"`
	VSwitchId                     string           `position:"Query" name:"VSwitchId"`
	PackageType                   string           `position:"Query" name:"PackageType"`
	PostStart                     string           `position:"Query" name:"PostStart"`
	PhpExtensions                 string           `position:"Body" name:"PhpExtensions"`
	AssociateEip                  requests.Boolean `position:"Body" name:"AssociateEip"`
	WebContainer                  string           `position:"Query" name:"WebContainer"`
	Memory                        requests.Integer `position:"Query" name:"Memory"`
	SlsConfigs                    string           `position:"Query" name:"SlsConfigs"`
	CommandArgs                   string           `position:"Query" name:"CommandArgs"`
	AcrAssumeRoleArn              string           `position:"Query" name:"AcrAssumeRoleArn"`
	Readiness                     string           `position:"Query" name:"Readiness"`
	Timezone                      string           `position:"Query" name:"Timezone"`
	OssAkId                       string           `position:"Body" name:"OssAkId"`
	Liveness                      string           `position:"Query" name:"Liveness"`
	SecurityGroupId               string           `position:"Query" name:"SecurityGroupId"`
	PackageVersion                string           `position:"Query" name:"PackageVersion"`
	TomcatConfig                  string           `position:"Query" name:"TomcatConfig"`
	Timeout                       requests.Integer `position:"Query" name:"Timeout"`
	WarStartOptions               string           `position:"Query" name:"WarStartOptions"`
	PackageRuntimeCustomBuild     string           `position:"Body" name:"PackageRuntimeCustomBuild"`
	EdasContainerVersion          string           `position:"Query" name:"EdasContainerVersion"`
	PackageUrl                    string           `position:"Query" name:"PackageUrl"`
	TerminationGracePeriodSeconds requests.Integer `position:"Query" name:"TerminationGracePeriodSeconds"`
	PhpConfig                     string           `position:"Body" name:"PhpConfig"`
	SliceEnvs                     string           `position:"Query" name:"SliceEnvs"`
	EnableImageAccl               requests.Boolean `position:"Body" name:"EnableImageAccl"`
	Replicas                      requests.Integer `position:"Query" name:"Replicas"`
	Workload                      string           `position:"Query" name:"Workload"`
	Command                       string           `position:"Query" name:"Command"`
	MountDesc                     string           `position:"Query" name:"MountDesc"`
	Jdk                           string           `position:"Query" name:"Jdk"`
	AppDescription                string           `position:"Query" name:"AppDescription"`
	AcrInstanceId                 string           `position:"Body" name:"AcrInstanceId"`
	VpcId                         string           `position:"Query" name:"VpcId"`
	ImageUrl                      string           `position:"Query" name:"ImageUrl"`
	Php                           string           `position:"Body" name:"Php"`
	RefAppId                      string           `position:"Query" name:"RefAppId"`
	PythonModules                 string           `position:"Query" name:"PythonModules"`
	PhpConfigLocation             string           `position:"Query" name:"PhpConfigLocation"`
}

// CreateJobResponse is the response struct for api CreateJob
type CreateJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateJobRequest creates a request to invoke CreateJob API
func CreateCreateJobRequest() (request *CreateJobRequest) {
	request = &CreateJobRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "CreateJob", "/pop/v1/sam/job/createJob", "serverless", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateJobResponse creates a response to parse from CreateJob response
func CreateCreateJobResponse() (response *CreateJobResponse) {
	response = &CreateJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
