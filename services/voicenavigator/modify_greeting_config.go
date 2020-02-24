package voicenavigator

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

// ModifyGreetingConfig invokes the voicenavigator.ModifyGreetingConfig API synchronously
// api document: https://help.aliyun.com/api/voicenavigator/modifygreetingconfig.html
func (client *Client) ModifyGreetingConfig(request *ModifyGreetingConfigRequest) (response *ModifyGreetingConfigResponse, err error) {
	response = CreateModifyGreetingConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyGreetingConfigWithChan invokes the voicenavigator.ModifyGreetingConfig API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/modifygreetingconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyGreetingConfigWithChan(request *ModifyGreetingConfigRequest) (<-chan *ModifyGreetingConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyGreetingConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyGreetingConfig(request)
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

// ModifyGreetingConfigWithCallback invokes the voicenavigator.ModifyGreetingConfig API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/modifygreetingconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyGreetingConfigWithCallback(request *ModifyGreetingConfigRequest, callback func(response *ModifyGreetingConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyGreetingConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyGreetingConfig(request)
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

// ModifyGreetingConfigRequest is the request struct for api ModifyGreetingConfig
type ModifyGreetingConfigRequest struct {
	*requests.RpcRequest
	GreetingWords string `position:"Query" name:"GreetingWords"`
	IntentTrigger string `position:"Query" name:"IntentTrigger"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	SourceType    string `position:"Query" name:"SourceType"`
}

// ModifyGreetingConfigResponse is the response struct for api ModifyGreetingConfig
type ModifyGreetingConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyGreetingConfigRequest creates a request to invoke ModifyGreetingConfig API
func CreateModifyGreetingConfigRequest() (request *ModifyGreetingConfigRequest) {
	request = &ModifyGreetingConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VoiceNavigator", "2018-06-12", "ModifyGreetingConfig", "voicebot", "openAPI")
	return
}

// CreateModifyGreetingConfigResponse creates a response to parse from ModifyGreetingConfig response
func CreateModifyGreetingConfigResponse() (response *ModifyGreetingConfigResponse) {
	response = &ModifyGreetingConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
