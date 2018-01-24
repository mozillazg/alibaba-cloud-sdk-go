package ccc

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

func (client *Client) ListPhoneNumbers(request *ListPhoneNumbersRequest) (response *ListPhoneNumbersResponse, err error) {
	response = CreateListPhoneNumbersResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListPhoneNumbersWithChan(request *ListPhoneNumbersRequest) (<-chan *ListPhoneNumbersResponse, <-chan error) {
	responseChan := make(chan *ListPhoneNumbersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPhoneNumbers(request)
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

func (client *Client) ListPhoneNumbersWithCallback(request *ListPhoneNumbersRequest, callback func(response *ListPhoneNumbersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPhoneNumbersResponse
		var err error
		defer close(result)
		response, err = client.ListPhoneNumbers(request)
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

type ListPhoneNumbersRequest struct {
	*requests.RpcRequest
	OutboundOnly requests.Boolean `position:"Query" name:"OutboundOnly"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
}

type ListPhoneNumbersResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PhoneNumbers   struct {
		PhoneNumber []struct {
			PhoneNumberId          string `json:"PhoneNumberId" xml:"PhoneNumberId"`
			InstanceId             string `json:"InstanceId" xml:"InstanceId"`
			Number                 string `json:"Number" xml:"Number"`
			PhoneNumberDescription string `json:"PhoneNumberDescription" xml:"PhoneNumberDescription"`
			TestOnly               bool   `json:"TestOnly" xml:"TestOnly"`
			RemainingTime          int    `json:"RemainingTime" xml:"RemainingTime"`
			AllowOutbound          bool   `json:"AllowOutbound" xml:"AllowOutbound"`
			Usage                  string `json:"Usage" xml:"Usage"`
			Trunks                 int    `json:"Trunks" xml:"Trunks"`
			ContactFlow            struct {
				ContactFlowId          string `json:"ContactFlowId" xml:"ContactFlowId"`
				InstanceId             string `json:"InstanceId" xml:"InstanceId"`
				ContactFlowName        string `json:"ContactFlowName" xml:"ContactFlowName"`
				ContactFlowDescription string `json:"ContactFlowDescription" xml:"ContactFlowDescription"`
				Type                   string `json:"Type" xml:"Type"`
			} `json:"ContactFlow" xml:"ContactFlow"`
		} `json:"PhoneNumber" xml:"PhoneNumber"`
	} `json:"PhoneNumbers" xml:"PhoneNumbers"`
}

func CreateListPhoneNumbersRequest() (request *ListPhoneNumbersRequest) {
	request = &ListPhoneNumbersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListPhoneNumbers", "", "")
	return
}

func CreateListPhoneNumbersResponse() (response *ListPhoneNumbersResponse) {
	response = &ListPhoneNumbersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}