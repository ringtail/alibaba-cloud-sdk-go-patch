package alimt

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

// Data is a nested struct in alimt response
type Data struct {
	Translated              string                            `json:"Translated" xml:"Translated"`
	WordSpelledCorrectError string                            `json:"WordSpelledCorrectError" xml:"WordSpelledCorrectError"`
	ContainCoreClasses      string                            `json:"ContainCoreClasses" xml:"ContainCoreClasses"`
	OverLengthLimit         string                            `json:"OverLengthLimit" xml:"OverLengthLimit"`
	Titles                  string                            `json:"Titles" xml:"Titles"`
	Language                string                            `json:"Language" xml:"Language"`
	DuplicateWords          string                            `json:"DuplicateWords" xml:"DuplicateWords"`
	NoFirstUppercaseList    string                            `json:"NoFirstUppercaseList" xml:"NoFirstUppercaseList"`
	ImageData               string                            `json:"ImageData" xml:"ImageData"`
	AllUppercaseWords       string                            `json:"AllUppercaseWords" xml:"AllUppercaseWords"`
	Url                     string                            `json:"Url" xml:"Url"`
	TemplateJson            string                            `json:"TemplateJson" xml:"TemplateJson"`
	TaskId                  string                            `json:"TaskId" xml:"TaskId"`
	LanguageQualityScore    string                            `json:"LanguageQualityScore" xml:"LanguageQualityScore"`
	InPaintingUrl           string                            `json:"InPaintingUrl" xml:"InPaintingUrl"`
	FinalImageUrl           string                            `json:"FinalImageUrl" xml:"FinalImageUrl"`
	PictureEditor           string                            `json:"PictureEditor" xml:"PictureEditor"`
	Status                  string                            `json:"Status" xml:"Status"`
	Orc                     string                            `json:"Orc" xml:"Orc"`
	WordCount               string                            `json:"WordCount" xml:"WordCount"`
	TotalScore              string                            `json:"TotalScore" xml:"TotalScore"`
	DisableWords            string                            `json:"DisableWords" xml:"DisableWords"`
	Result                  []CertificateImageTranslateResult `json:"Result" xml:"Result"`
	TranslatedValues        []CertificateTranslateItemDTO     `json:"TranslatedValues" xml:"TranslatedValues"`
}
