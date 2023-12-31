package scdn

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

// DataModule is a nested struct in scdn response
type DataModule struct {
	HttpAccValue            string                  `json:"HttpAccValue" xml:"HttpAccValue"`
	BpsValue                string                  `json:"BpsValue" xml:"BpsValue"`
	HttpQpsValue            string                  `json:"HttpQpsValue" xml:"HttpQpsValue"`
	ByteHitRate             string                  `json:"ByteHitRate" xml:"ByteHitRate"`
	HttpsTrafficValue       string                  `json:"HttpsTrafficValue" xml:"HttpsTrafficValue"`
	AccValue                string                  `json:"AccValue" xml:"AccValue"`
	Value                   string                  `json:"Value" xml:"Value"`
	TrafficValue            string                  `json:"TrafficValue" xml:"TrafficValue"`
	TimeStamp               string                  `json:"TimeStamp" xml:"TimeStamp"`
	QpsValue                string                  `json:"QpsValue" xml:"QpsValue"`
	ReqHitRate              string                  `json:"ReqHitRate" xml:"ReqHitRate"`
	HttpBpsValue            string                  `json:"HttpBpsValue" xml:"HttpBpsValue"`
	HttpsBpsValue           string                  `json:"HttpsBpsValue" xml:"HttpsBpsValue"`
	HttpsQpsValue           string                  `json:"HttpsQpsValue" xml:"HttpsQpsValue"`
	HttpsOriginBpsValue     string                  `json:"HttpsOriginBpsValue" xml:"HttpsOriginBpsValue"`
	OriginBpsValue          string                  `json:"OriginBpsValue" xml:"OriginBpsValue"`
	HttpTrafficValue        string                  `json:"HttpTrafficValue" xml:"HttpTrafficValue"`
	HttpsAccValue           string                  `json:"HttpsAccValue" xml:"HttpsAccValue"`
	HttpOriginBpsValue      string                  `json:"HttpOriginBpsValue" xml:"HttpOriginBpsValue"`
	HttpCodeDataPerInterval HttpCodeDataPerInterval `json:"HttpCodeDataPerInterval" xml:"HttpCodeDataPerInterval"`
}
