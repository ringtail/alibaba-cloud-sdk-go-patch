package mse

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

// DataInGetGatewayAuthConsumerDetail is a nested struct in mse response
type DataInGetGatewayAuthConsumerDetail struct {
	Id              int64      `json:"Id" xml:"Id"`
	Name            string     `json:"Name" xml:"Name"`
	ConsumerStatus  bool       `json:"ConsumerStatus" xml:"ConsumerStatus"`
	PrimaryUser     string     `json:"PrimaryUser" xml:"PrimaryUser"`
	GatewayUniqueId string     `json:"GatewayUniqueId" xml:"GatewayUniqueId"`
	Type            string     `json:"Type" xml:"Type"`
	GmtCreate       string     `json:"GmtCreate" xml:"GmtCreate"`
	GmtModified     string     `json:"GmtModified" xml:"GmtModified"`
	Description     string     `json:"Description" xml:"Description"`
	EncodeType      string     `json:"EncodeType" xml:"EncodeType"`
	Jwks            string     `json:"Jwks" xml:"Jwks"`
	TokenName       string     `json:"TokenName" xml:"TokenName"`
	TokenPass       bool       `json:"TokenPass" xml:"TokenPass"`
	TokenPosition   string     `json:"TokenPosition" xml:"TokenPosition"`
	TokenPrefix     string     `json:"TokenPrefix" xml:"TokenPrefix"`
	KeyName         string     `json:"KeyName" xml:"KeyName"`
	KeyValue        string     `json:"KeyValue" xml:"KeyValue"`
	ResourceList    []Resource `json:"ResourceList" xml:"ResourceList"`
}
