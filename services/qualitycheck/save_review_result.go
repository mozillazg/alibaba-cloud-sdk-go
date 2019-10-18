package qualitycheck

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

// SaveReviewResult invokes the qualitycheck.SaveReviewResult API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/savereviewresult.html
func (client *Client) SaveReviewResult(request *SaveReviewResultRequest) (response *SaveReviewResultResponse, err error) {
	response = CreateSaveReviewResultResponse()
	err = client.DoAction(request, response)
	return
}

// SaveReviewResultWithChan invokes the qualitycheck.SaveReviewResult API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/savereviewresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveReviewResultWithChan(request *SaveReviewResultRequest) (<-chan *SaveReviewResultResponse, <-chan error) {
	responseChan := make(chan *SaveReviewResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveReviewResult(request)
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

// SaveReviewResultWithCallback invokes the qualitycheck.SaveReviewResult API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/savereviewresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveReviewResultWithCallback(request *SaveReviewResultRequest, callback func(response *SaveReviewResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveReviewResultResponse
		var err error
		defer close(result)
		response, err = client.SaveReviewResult(request)
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

// SaveReviewResultRequest is the request struct for api SaveReviewResult
type SaveReviewResultRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// SaveReviewResultResponse is the response struct for api SaveReviewResult
type SaveReviewResultResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateSaveReviewResultRequest creates a request to invoke SaveReviewResult API
func CreateSaveReviewResultRequest() (request *SaveReviewResultRequest) {
	request = &SaveReviewResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "SaveReviewResult", "", "")
	return
}

// CreateSaveReviewResultResponse creates a response to parse from SaveReviewResult response
func CreateSaveReviewResultResponse() (response *SaveReviewResultResponse) {
	response = &SaveReviewResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}