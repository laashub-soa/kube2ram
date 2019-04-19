package scdn

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

// DescribeScdnDomainTrafficData invokes the scdn.DescribeScdnDomainTrafficData API synchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomaintrafficdata.html
func (client *Client) DescribeScdnDomainTrafficData(request *DescribeScdnDomainTrafficDataRequest) (response *DescribeScdnDomainTrafficDataResponse, err error) {
	response = CreateDescribeScdnDomainTrafficDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnDomainTrafficDataWithChan invokes the scdn.DescribeScdnDomainTrafficData API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomaintrafficdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnDomainTrafficDataWithChan(request *DescribeScdnDomainTrafficDataRequest) (<-chan *DescribeScdnDomainTrafficDataResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnDomainTrafficDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnDomainTrafficData(request)
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

// DescribeScdnDomainTrafficDataWithCallback invokes the scdn.DescribeScdnDomainTrafficData API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomaintrafficdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnDomainTrafficDataWithCallback(request *DescribeScdnDomainTrafficDataRequest, callback func(response *DescribeScdnDomainTrafficDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnDomainTrafficDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnDomainTrafficData(request)
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

// DescribeScdnDomainTrafficDataRequest is the request struct for api DescribeScdnDomainTrafficData
type DescribeScdnDomainTrafficDataRequest struct {
	*requests.RpcRequest
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	Interval       string           `position:"Query" name:"Interval"`
}

// DescribeScdnDomainTrafficDataResponse is the response struct for api DescribeScdnDomainTrafficData
type DescribeScdnDomainTrafficDataResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	DomainName             string                 `json:"DomainName" xml:"DomainName"`
	StartTime              string                 `json:"StartTime" xml:"StartTime"`
	EndTime                string                 `json:"EndTime" xml:"EndTime"`
	DataInterval           string                 `json:"DataInterval" xml:"DataInterval"`
	TrafficDataPerInterval TrafficDataPerInterval `json:"TrafficDataPerInterval" xml:"TrafficDataPerInterval"`
}

// CreateDescribeScdnDomainTrafficDataRequest creates a request to invoke DescribeScdnDomainTrafficData API
func CreateDescribeScdnDomainTrafficDataRequest() (request *DescribeScdnDomainTrafficDataRequest) {
	request = &DescribeScdnDomainTrafficDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnDomainTrafficData", "scdn", "openAPI")
	return
}

// CreateDescribeScdnDomainTrafficDataResponse creates a response to parse from DescribeScdnDomainTrafficData response
func CreateDescribeScdnDomainTrafficDataResponse() (response *DescribeScdnDomainTrafficDataResponse) {
	response = &DescribeScdnDomainTrafficDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}