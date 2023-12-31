package cloudesl

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

// EventInfoVO is a nested struct in cloudesl response
type EventInfoVO struct {
	EventId     string `json:"EventId" xml:"EventId"`
	EventTime   string `json:"EventTime" xml:"EventTime"`
	Category    string `json:"Category" xml:"Category"`
	EslBarCode  string `json:"EslBarCode" xml:"EslBarCode"`
	ApMac       string `json:"ApMac" xml:"ApMac"`
	StoreId     string `json:"StoreId" xml:"StoreId"`
	ItemBarCode string `json:"ItemBarCode" xml:"ItemBarCode"`
	ItemId      string `json:"ItemId" xml:"ItemId"`
	ItemTitle   string `json:"ItemTitle" xml:"ItemTitle"`
	Content     string `json:"Content" xml:"Content"`
	Status      string `json:"Status" xml:"Status"`
	Staff       string `json:"Staff" xml:"Staff"`
	ProcessTime string `json:"ProcessTime" xml:"ProcessTime"`
	Reason      string `json:"Reason" xml:"Reason"`
}
