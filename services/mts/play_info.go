package mts

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

func (client *Client) PlayInfo(request *PlayInfoRequest) (response *PlayInfoResponse, err error) {
	response = CreatePlayInfoResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) PlayInfoWithChan(request *PlayInfoRequest) (<-chan *PlayInfoResponse, <-chan error) {
	responseChan := make(chan *PlayInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PlayInfo(request)
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

func (client *Client) PlayInfoWithCallback(request *PlayInfoRequest, callback func(response *PlayInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PlayInfoResponse
		var err error
		defer close(result)
		response, err = client.PlayInfo(request)
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

type PlayInfoRequest struct {
	*requests.RpcRequest
	PlayDomain           string           `position:"Query" name:"PlayDomain"`
	ResourceOwnerId      string           `position:"Query" name:"ResourceOwnerId"`
	Formats              string           `position:"Query" name:"Formats"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	HlsUriToken          string           `position:"Query" name:"HlsUriToken"`
	OwnerId              string           `position:"Query" name:"OwnerId"`
	MediaId              string           `position:"Query" name:"MediaId"`
	Rand                 string           `position:"Query" name:"Rand"`
	AuthTimeout          requests.Integer `position:"Query" name:"AuthTimeout"`
	AuthInfo             string           `position:"Query" name:"AuthInfo"`
}

type PlayInfoResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	NotFoundCDNDomain struct {
		String []string `json:"String" xml:"String"`
	} `json:"NotFoundCDNDomain" xml:"NotFoundCDNDomain"`
	PlayInfoList struct {
		PlayInfo []struct {
			Url            string `json:"Url" xml:"Url"`
			Duration       string `json:"duration" xml:"duration"`
			Size           string `json:"size" xml:"size"`
			Width          string `json:"width" xml:"width"`
			Height         string `json:"height" xml:"height"`
			Bitrate        string `json:"bitrate" xml:"bitrate"`
			Fps            string `json:"fps" xml:"fps"`
			Format         string `json:"format" xml:"format"`
			Definition     string `json:"definition" xml:"definition"`
			Encryption     string `json:"encryption" xml:"encryption"`
			Rand           string `json:"rand" xml:"rand"`
			Plaintext      string `json:"plaintext" xml:"plaintext"`
			Complexity     string `json:"complexity" xml:"complexity"`
			ActivityName   string `json:"activityName" xml:"activityName"`
			EncryptionType string `json:"encryptionType" xml:"encryptionType"`
			DownloadType   string `json:"downloadType" xml:"downloadType"`
		} `json:"PlayInfo" xml:"PlayInfo"`
	} `json:"PlayInfoList" xml:"PlayInfoList"`
}

func CreatePlayInfoRequest() (request *PlayInfoRequest) {
	request = &PlayInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "PlayInfo", "mts", "openAPI")
	return
}

func CreatePlayInfoResponse() (response *PlayInfoResponse) {
	response = &PlayInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}