package linkwan

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

// PacketInListNodeGroupTransferPackets is a nested struct in linkwan response
type PacketInListNodeGroupTransferPackets struct {
	GwEui                   string  `json:"GwEui" xml:"GwEui"`
	DevEui                  string  `json:"DevEui" xml:"DevEui"`
	DevAddr                 string  `json:"DevAddr" xml:"DevAddr"`
	Freq                    float64 `json:"Freq" xml:"Freq"`
	Datr                    string  `json:"Datr" xml:"Datr"`
	ClassMode               string  `json:"ClassMode" xml:"ClassMode"`
	Rssi                    int     `json:"Rssi" xml:"Rssi"`
	Lsnr                    float64 `json:"Lsnr" xml:"Lsnr"`
	FPort                   int     `json:"FPort" xml:"FPort"`
	GwOwnerAliyunId         string  `json:"GwOwnerAliyunId" xml:"GwOwnerAliyunId"`
	FreqBandPlanGroupId     int64   `json:"FreqBandPlanGroupId" xml:"FreqBandPlanGroupId"`
	LogMillis               int64   `json:"LogMillis" xml:"LogMillis"`
	HasMacCommand           bool    `json:"HasMacCommand" xml:"HasMacCommand"`
	HasData                 bool    `json:"HasData" xml:"HasData"`
	Base64EncodedMacPayload string  `json:"Base64EncodedMacPayload" xml:"Base64EncodedMacPayload"`
	MacPayloadSize          int64   `json:"MacPayloadSize" xml:"MacPayloadSize"`
	ProcessEvent            string  `json:"ProcessEvent" xml:"ProcessEvent"`
	MessageType             string  `json:"MessageType" xml:"MessageType"`
	MacCommandCIDs          string  `json:"MacCommandCIDs" xml:"MacCommandCIDs"`
	FcntUp                  int64   `json:"FcntUp" xml:"FcntUp"`
	FcntDown                int64   `json:"FcntDown" xml:"FcntDown"`
}