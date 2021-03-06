package cloudapi

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

// DeletePlugin invokes the cloudapi.DeletePlugin API synchronously
// api document: https://help.aliyun.com/api/cloudapi/deleteplugin.html
func (client *Client) DeletePlugin(request *DeletePluginRequest) (response *DeletePluginResponse, err error) {
	response = CreateDeletePluginResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePluginWithChan invokes the cloudapi.DeletePlugin API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/deleteplugin.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePluginWithChan(request *DeletePluginRequest) (<-chan *DeletePluginResponse, <-chan error) {
	responseChan := make(chan *DeletePluginResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePlugin(request)
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

// DeletePluginWithCallback invokes the cloudapi.DeletePlugin API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/deleteplugin.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePluginWithCallback(request *DeletePluginRequest, callback func(response *DeletePluginResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePluginResponse
		var err error
		defer close(result)
		response, err = client.DeletePlugin(request)
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

// DeletePluginRequest is the request struct for api DeletePlugin
type DeletePluginRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	PluginId      string `position:"Query" name:"PluginId"`
}

// DeletePluginResponse is the response struct for api DeletePlugin
type DeletePluginResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeletePluginRequest creates a request to invoke DeletePlugin API
func CreateDeletePluginRequest() (request *DeletePluginRequest) {
	request = &DeletePluginRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DeletePlugin", "apigateway", "openAPI")
	return
}

// CreateDeletePluginResponse creates a response to parse from DeletePlugin response
func CreateDeletePluginResponse() (response *DeletePluginResponse) {
	response = &DeletePluginResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
