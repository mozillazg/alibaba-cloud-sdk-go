package jarvis

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

// DescribePunishList invokes the jarvis.DescribePunishList API synchronously
// api document: https://help.aliyun.com/api/jarvis/describepunishlist.html
func (client *Client) DescribePunishList(request *DescribePunishListRequest) (response *DescribePunishListResponse, err error) {
	response = CreateDescribePunishListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePunishListWithChan invokes the jarvis.DescribePunishList API asynchronously
// api document: https://help.aliyun.com/api/jarvis/describepunishlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePunishListWithChan(request *DescribePunishListRequest) (<-chan *DescribePunishListResponse, <-chan error) {
	responseChan := make(chan *DescribePunishListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePunishList(request)
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

// DescribePunishListWithCallback invokes the jarvis.DescribePunishList API asynchronously
// api document: https://help.aliyun.com/api/jarvis/describepunishlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePunishListWithCallback(request *DescribePunishListRequest, callback func(response *DescribePunishListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePunishListResponse
		var err error
		defer close(result)
		response, err = client.DescribePunishList(request)
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

// DescribePunishListRequest is the request struct for api DescribePunishList
type DescribePunishListRequest struct {
	*requests.RpcRequest
	SrcIP        string           `position:"Query" name:"SrcIP"`
	SourceIp     string           `position:"Query" name:"SourceIp"`
	PageSize     requests.Integer `position:"Query" name:"pageSize"`
	CurrentPage  requests.Integer `position:"Query" name:"currentPage"`
	PunishStatus string           `position:"Query" name:"PunishStatus"`
	Lang         string           `position:"Query" name:"Lang"`
	SrcUid       requests.Integer `position:"Query" name:"srcUid"`
	SourceCode   string           `position:"Query" name:"sourceCode"`
}

// DescribePunishListResponse is the response struct for api DescribePunishList
type DescribePunishListResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Module    string   `json:"Module" xml:"Module"`
	PageInfo  PageInfo `json:"PageInfo" xml:"PageInfo"`
	DataList  []Data   `json:"DataList" xml:"DataList"`
}

// CreateDescribePunishListRequest creates a request to invoke DescribePunishList API
func CreateDescribePunishListRequest() (request *DescribePunishListRequest) {
	request = &DescribePunishListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jarvis", "2018-02-06", "DescribePunishList", "", "")
	return
}

// CreateDescribePunishListResponse creates a response to parse from DescribePunishList response
func CreateDescribePunishListResponse() (response *DescribePunishListResponse) {
	response = &DescribePunishListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
