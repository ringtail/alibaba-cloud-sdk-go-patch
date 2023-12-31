package privatelink

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

// Endpoint is a nested struct in privatelink response
type Endpoint struct {
	VpcId                  string     `json:"VpcId" xml:"VpcId"`
	EndpointName           string     `json:"EndpointName" xml:"EndpointName"`
	EndpointType           string     `json:"EndpointType" xml:"EndpointType"`
	CreateTime             string     `json:"CreateTime" xml:"CreateTime"`
	ServiceId              string     `json:"ServiceId" xml:"ServiceId"`
	ZoneAffinityEnabled    bool       `json:"ZoneAffinityEnabled" xml:"ZoneAffinityEnabled"`
	EndpointDomain         string     `json:"EndpointDomain" xml:"EndpointDomain"`
	EndpointStatus         string     `json:"EndpointStatus" xml:"EndpointStatus"`
	RegionId               string     `json:"RegionId" xml:"RegionId"`
	ResourceOwner          bool       `json:"ResourceOwner" xml:"ResourceOwner"`
	Bandwidth              int64      `json:"Bandwidth" xml:"Bandwidth"`
	ConnectionStatus       string     `json:"ConnectionStatus" xml:"ConnectionStatus"`
	EndpointDescription    string     `json:"EndpointDescription" xml:"EndpointDescription"`
	EndpointId             string     `json:"EndpointId" xml:"EndpointId"`
	EndpointBusinessStatus string     `json:"EndpointBusinessStatus" xml:"EndpointBusinessStatus"`
	ServiceName            string     `json:"ServiceName" xml:"ServiceName"`
	ResourceGroupId        string     `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Tags                   []TagModel `json:"Tags" xml:"Tags"`
}
