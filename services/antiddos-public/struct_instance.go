package antiddos_public

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

// Instance is a nested struct in antiddos_public response
type Instance struct {
	Region              string                `json:"Region" xml:"Region"`
	InternetIp          string                `json:"InternetIp" xml:"InternetIp"`
	ElasticThreshold    int                   `json:"ElasticThreshold" xml:"ElasticThreshold"`
	InstanceName        string                `json:"InstanceName" xml:"InstanceName"`
	IsBgppack           bool                  `json:"IsBgppack" xml:"IsBgppack"`
	InstanceType        string                `json:"InstanceType" xml:"InstanceType"`
	DefensePpsThreshold int                   `json:"DefensePpsThreshold" xml:"DefensePpsThreshold"`
	IpVersion           string                `json:"IpVersion" xml:"IpVersion"`
	BlackholeThreshold  int                   `json:"BlackholeThreshold" xml:"BlackholeThreshold"`
	InstanceId          string                `json:"InstanceId" xml:"InstanceId"`
	InstanceStatus      string                `json:"InstanceStatus" xml:"InstanceStatus"`
	InstanceIp          string                `json:"InstanceIp" xml:"InstanceIp"`
	DefenseBpsThreshold int                   `json:"DefenseBpsThreshold" xml:"DefenseBpsThreshold"`
	IpAddressConfig     []IpAddressConfigItem `json:"IpAddressConfig" xml:"IpAddressConfig"`
}
