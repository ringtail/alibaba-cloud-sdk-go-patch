package live

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

// CustomTranscodeParameters is a nested struct in live response
type CustomTranscodeParameters struct {
	VideoProfile    string `json:"VideoProfile" xml:"VideoProfile"`
	AudioBitrate    int    `json:"AudioBitrate" xml:"AudioBitrate"`
	RtsFlag         string `json:"RtsFlag" xml:"RtsFlag"`
	Height          int    `json:"Height" xml:"Height"`
	TemplateType    string `json:"TemplateType" xml:"TemplateType"`
	Bframes         string `json:"Bframes" xml:"Bframes"`
	AudioRate       int    `json:"AudioRate" xml:"AudioRate"`
	FPS             int    `json:"FPS" xml:"FPS"`
	AudioCodec      string `json:"AudioCodec" xml:"AudioCodec"`
	Gop             string `json:"Gop" xml:"Gop"`
	VideoBitrate    int    `json:"VideoBitrate" xml:"VideoBitrate"`
	Width           int    `json:"Width" xml:"Width"`
	AudioChannelNum int    `json:"AudioChannelNum" xml:"AudioChannelNum"`
	AudioProfile    string `json:"AudioProfile" xml:"AudioProfile"`
}
