package linkvisual

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

// DataItemInDetectUserFaceByUrl is a nested struct in linkvisual response
type DataItemInDetectUserFaceByUrl struct {
	BlurScore          float64   `json:"BlurScore" xml:"BlurScore"`
	Gender             int       `json:"Gender" xml:"Gender"`
	OcclusionScore     float64   `json:"OcclusionScore" xml:"OcclusionScore"`
	GoodForLibrary     bool      `json:"GoodForLibrary" xml:"GoodForLibrary"`
	GoodForRecognition bool      `json:"GoodForRecognition" xml:"GoodForRecognition"`
	Age                int       `json:"Age" xml:"Age"`
	FaceProbability    float64   `json:"FaceProbability" xml:"FaceProbability"`
	PoseScore          float64   `json:"PoseScore" xml:"PoseScore"`
	FaceRects          FaceRects `json:"FaceRects" xml:"FaceRects"`
	Landmarks          Landmarks `json:"Landmarks" xml:"Landmarks"`
}
